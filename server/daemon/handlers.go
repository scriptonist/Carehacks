package daemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/scriptonist/Carehacks/server/azure"
	db "github.com/scriptonist/Carehacks/server/db"
	"github.com/scriptonist/Carehacks/server/structs"
)

// PongHandler Return pong
func PongHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello")
		},
	)
}

// SearchForMedicines --
// Returns medicine stores containing the given medicine
// In order of the distance between user and store.
// Endpoint - /search
// Expected POST Request
// {
// 		medicine:"",
// 		"lat":"",
// 		"lon"
// }
func SearchForMedicines() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			query := structs.SearchForMedicinesRequest{}
			err := json.NewDecoder(r.Body).Decode(&query)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, err.Error())
			}
			res := db.MedicineSearch(query)
			respondWithJSON(w, http.StatusOK, res)
		},
	)
}

// UserOrder --
// Initiates a user order for a set of medicines
// endpoint - /user/order
func UserOrder() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			respondWithJSON(w, http.StatusOK, "not implemented")
		},
	)
}

// UploadPrescription --
// Accepts image files with presciption details
func UploadPrescription() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Saving file")
		r.ParseMultipartForm(32 << 20)
		file, h, err := r.FormFile("file")
		mimeType := h.Header.Get("Content-Type")

		log.Println(file)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, structs.UploadPrescriptionResponse{
				Message: "Upload Failed",
				Done:    false,
			})
		}
		defer file.Close()
		fileContents := make([]byte, h.Size)

		l, err := file.Read(fileContents)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, structs.UploadPrescriptionResponse{
				Message: "Upload Failed",
				Done:    false,
			})
		}

		log.Println(l)
		url, err := azure.CreatePrescriptionBlob(h.Filename, &fileContents, mimeType)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, structs.UploadPrescriptionResponse{
				Message: "Upload Failed",
				Done:    false,
			})
		}

		respondWithJSON(w, http.StatusInternalServerError, structs.UploadPrescriptionResponse{
			Message: "Upload Done",
			Done:    true,
			URL:     url,
		})
	},
	)
}

// -------------- Util Funcs ----------- //
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, struct{ Error string }{
		Error: message,
	})
	return
}
