package daemon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scriptonist/Carehacks/server/azure"
	"github.com/scriptonist/cleverhires/API/globals"
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

			respondWithJSON(w, http.StatusOK, "not implemented")
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
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		var fileContents []byte
		_, err = file.Read(fileContents)
		azure.CreatePrescriptionBlob(handler.Filename, &fileContents)

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
	respondWithJSON(w, code, globals.ErrorResponse{
		Error: message,
	})
	return
}
