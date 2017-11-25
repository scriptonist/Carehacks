package qr

import (
	"fmt"

	"github.com/scriptonist/Carehacks/server/azure"
	qrcode "github.com/skip2/go-qrcode"
)

// CreateQR creates QR Stores it in azure blob
// Returns a URL and eroor
func CreateQR(url, blobname string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", fmt.Errorf("Failed generate QR Code: %v", err)
	}

	// Create blob
	url, err = azure.CreateQRFileBlob(blobname, &png)
	if err != nil {
		return "", err
	}

	return url, nil
}

// func main() {
// 	blobURL, err := CreateQR("https://www.google.com", "google.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(blobURL)

// }
