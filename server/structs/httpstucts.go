package structs

import (
	"gopkg.in/mgo.v2/bson"
)

// ----------------- Handler - SearchForMedicines  ---------------//

//StoreResult Represents a store
// An array of this struct will be returned
// As result for SearchForMedicines
type StoreResult struct {
	StoreName string        `json:"storename" bson:"Name"`
	StoreID   bson.ObjectId `bson:"_id" json:"store_id"`
	Distance  string        `json:"distance"`
	Avialable bool          `json:"avialable"`
}

type Medicine struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// SearchForMedicinesRequest --
// Request format for
// endpoint - /search
type SearchForMedicinesRequest struct {
	Medicine []string `json:"medicine"`
	Lat      float32  `json:"lat"`
	Lon      float32  `json:"lon"`
}

// SearchForMedicinesResponse --
// Request format for
// endpoint - /search
type SearchForMedicinesResponse struct {
	Stores []StoreResult `json:"stores"`
}

// ----------------- Handler - UserOrder  ---------------//

// UserOrderRequest --
type UserOrderRequest struct {
	UserID       string     `json:"user_id"`
	Medicines    []Medicine `json:"medicines"`
	StoreID      string     `json:"store_id"`
	Prescription string     `json:"prescription"`
	QrURL        string     `json:"qr_url,omitempty"`
	status       string     `json:"status,omitempty"`
}

// UserOrderResponse a
type UserOrderResponse struct {
	QrURL string `json:"qr_url,omitempty"`
}

// --------------Handler - UploadPrescription -------//

// UploadPrescriptionResponse --
type UploadPrescriptionResponse struct {
	Message string `json:"message"`
	Done    bool   `json:"done"`
	URL     string `json:"url,omitempty"`
}
