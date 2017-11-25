package main

import (
	"gopkg.in/mgo.v2/bson"

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

func MedicineSearch(searchTerm string) []bson.M {

	var searchResult []bson.M
	regex := bson.M{"$regex": bson.RegEx{Pattern: searchTerm, Options: "i"}}
	pipe := store.Pipe([]bson.M{
		{"$unwind": "$inventory"},
		{"$match": bson.M{"inventory.name": regex}},
		{"$project": bson.M{"Name": 1, "Location": 1, "inventory.name": 1}}})

	err := pipe.All(&searchResult)
	if err != nil {
		panic(err)
	}
	return searchResult
}

func main() {
	MedicineSearch("vic")
}
