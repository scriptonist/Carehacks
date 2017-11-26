package db

import (
	"github.com/scriptonist/Carehacks/server/qr"
	"gopkg.in/mgo.v2/bson"

	"github.com/scriptonist/Carehacks/server/structs"
	mgo "gopkg.in/mgo.v2"
)

type CourseDetail struct {
	Name   string `bson:"name,omitempty"`
	Course string `bson:"course,omitempty"`
}

type User struct {
	Id        string         `bson:"_id,omitempty"`
	Name      string         `bson:"name,omitempty"`
	Place     string         `bson:"place,omitempty"`
	Mobile    string         `bson:"mobile,omitempty"`
	Order     string         `bson:"order,omitempty"`
	Medicines []CourseDetail `bson:"medicines,omitempty"`
}

type MedicineDetails struct {
	Name     string `bson:"name,omitempty"`
	Quantity string `bson:"quantity,omitempty"`
}

type OrderDetails struct {
	UserID       string            `bson:"userid,omitempty"`
	Medicines    []MedicineDetails `bson:"medicines,omitempty"`
	Status       string            `bson:"status,omitempty"`
	Prescription string            `bson:"presciption,omitempty"`
	QR           string            `bson:"QR,omitempty"`
}

type Store struct {
	id        string            `bson:"_id,omitempty"`
	Name      string            `bson:"Name,omitempty"`
	Location  []string          `bson:"Location,omitempty"`
	Orders    []OrderDetails    `bson:"orders,omitempty"`
	Inventory []MedicineDetails `bson:"inventory,omitempty"`
}

var (
	store *mgo.Collection
	users *mgo.Collection
)

func init() {
	session, err := mgo.Dial("mongodb://carehack:carehack123@ds159845.mlab.com:59845/carehack")

	if err != nil {
		panic(err)
	}

	store = session.DB("carehack").C("store")
	users = session.DB("carehack").C("users")
}

func getUserMedicineCourse(userID string) (*User, error) {
	var user User
	err := users.FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func insertUserMedicineCourse(user User) error {
	err := users.Update(bson.M{"_id": bson.ObjectIdHex(user.Id)}, bson.M{"$push": bson.M{"medicines": bson.M{"$each": user.Medicines}}})
	if err != nil {
		return err
	}

	return nil
}

func insertUserQRCode(user *User) (string, error) {
	// log.Println(user.Order)
	qrURL, err := qr.CreateQR("http://chack-server.azurewebsites.net/user/checkout_order/"+user.Id+"/"+user.Order, user.Id)

	err = users.Update(bson.M{"_id": bson.ObjectIdHex(user.Id)}, bson.M{"order": qrURL})
	if err != nil {
		return "", err
	}

	return qrURL, nil
}

func MedicineSearch(searchDetails structs.SearchForMedicinesRequest) *structs.SearchForMedicinesResponse {
	var searchResult []structs.StoreResult
	// "coordinates": bson.M{"$geoWithin": bson.M{"$centerSphere": []interface{}{[]interface{}{searchDetails.Lat, searchDetails.Lon}, 5 / 3963.2}}}}
	store.Find(bson.M{"inventory.name": bson.M{"$all": searchDetails.Medicine}}).All(&searchResult)

	return &structs.SearchForMedicinesResponse{Stores: searchResult}
}

func PlaceOrder(order structs.UserOrderRequest) (string, error) {
	order.Status = "False"
	err := store.Update(bson.M{"_id": bson.ObjectIdHex(order.StoreID)}, bson.M{"$push": bson.M{"orders": order}})
	if err != nil {
		return "", err
	}
	url, err := insertUserQRCode(&User{
		Id:    order.UserID,
		Order: order.StoreID,
	})
	if err != nil {
		return "", err
	}
	return url, err
}

func ListOrders(storeID string) Store {
	var orders Store
	err := store.Find(bson.M{"_id": bson.ObjectIdHex(storeID)}).Select(bson.M{"orders": 1}).One(&orders)
	if err != nil {
		panic(err)
	}

	return orders
}

func InsertQrURLInUsersCollection(userID string, QRurl string) error {
	err := users.Update(bson.M{"_id": bson.ObjectIdHex(userID)}, bson.M{"$set": bson.M{"QR": QRurl}})
	if err != nil {
		panic(err)
	}

	return nil
}

func DeleteFromOrders(userID, storeID string) error {
	err := store.Update(bson.M{"_id": bson.ObjectIdHex(storeID)}, bson.M{"$pull": bson.M{"orders": bson.M{"userid": userID}}})
	if err != nil {
		return err
	}

	return nil
}

// func main() {
// 	orders := ListOrders("5a1931df0197d51ab40f7a59")
// 	fmt.Println(orders)
// }
