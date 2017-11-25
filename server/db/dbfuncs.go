package db

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/scriptonist/Carehacks/server/daemon"
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

func MedicineSearch(searchDetails daemon.SearchForMedicinesRequest) *daemon.SearchForMedicinesResponse {
	var searchResult []daemon.StoreResult
	err := store.Find(bson.M{"inventory.name": bson.M{"$all": searchDetails.Medicine}}).All(&searchResult)
	if err != nil {
		panic(err)
	}
	return &daemon.SearchForMedicinesResponse{Stores: searchResult}
}

func PlaceOrder(storeID string, order OrderDetails) error {
	err := store.Update(bson.M{"_id": bson.ObjectIdHex(storeID)}, bson.M{"$push": bson.M{"orders": order}})
	if err != nil {
		return err
	}

	return nil
}

func ListOrders(storeID string) Store {
	var orders Store
	err := store.Find(bson.M{"_id": bson.ObjectIdHex(storeID)}).Select(bson.M{"orders": 1}).One(&orders)
	if err != nil {
		panic(err)
	}

	return orders
}

func main() {
	orders := ListOrders("5a1931df0197d51ab40f7a59")
	fmt.Println(orders)
}
