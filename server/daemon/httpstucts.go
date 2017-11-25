package daemon

// ----------------- Handler - SearchForMedicines  ---------------//

//StoreResult Represents a store
// An array of this struct will be returned
// As result for SearchForMedicines
type StoreResult struct {
	StoreName string `json:"storename" bson:"Name"`
	StoreID   string `bson:"_id"`
	Distance  string `json:"distance"`
	Avialable bool   `json:"avialable"`
}

// SearchForMedicinesRequest --
// Request format for
// endpoint - /search
type SearchForMedicinesRequest struct {
	Medicine []string `json:"medicine"`
	Lat      int      `json:"lat"`
	Lon      int      `json:"lon"`
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
	UserID       string   `json:"user_id"`
	Medicines    []string `json:"medicines"`
	StoreID      string   `json:"store_id"`
	Prescription string   `json:"prescription"`
}

// UserOrderResponse a
type UserOrderResponse struct {
}

// --------------Handler - UploadPrescription -------//

// UploadPrescriptionResponse --
type UploadPrescriptionResponse struct {
	Message string `json:"message"`
	Done    bool   `json:"done"`
	URL     string `json:"url,omitempty"`
}
