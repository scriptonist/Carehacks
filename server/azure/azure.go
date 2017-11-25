package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/storage"
)

var (
	blobClient            storage.BlobStorageClient
	carehackQr            *storage.Container
	carehackPrescriptions *storage.Container
)

func init() {

	accountName := "carehack"
	accountKey := "/UTpXRdb3vv/hWh1kR5R/WUqwIbk6ASfSeYCvsrrGc0lowl7KxdMZ3X8QkZk96shE+zndWk9zAm32l3QZmdavw=="
	client, err := storage.NewBasicClient(accountName, accountKey)
	if err != nil {
		panic("Connecting to Blob Service Failed!")
	}

	blobClient = client.GetBlobService()

	// Reference QR Container
	carehackQr = blobClient.GetContainerReference("carehack-qr")
	carehackPrescriptions = blobClient.GetContainerReference("carehack-prescriptions")

}

// CreateQRFileBlob a blob for QR Code
func CreateQRFileBlob(filename string, content *[]byte) (string, error) {
	// Get reference to blob
	blob := carehackQr.GetBlobReference(filename)
	blob.Properties.ContentType = "image/png"
	// Create block blob
	err := blob.PutAppendBlob(nil)

	if err != nil {
		return "", fmt.Errorf("Creating blob failed. %v", err)
	}
	// write contents to blob
	err = blob.AppendBlock(*content, nil)
	if err != nil {
		return "", fmt.Errorf("Cannot Add contents to blob. %v", err)
	}
	return blob.GetURL(), nil
}

// CreatePrescriptionBlob a blob for QR Code
func CreatePrescriptionBlob(filename string, content *[]byte, mimeType string) (string, error) {
	// Get reference to blob
	blob := carehackPrescriptions.GetBlobReference(filename)
	blob.Properties.ContentType = mimeType
	// Create block blob
	err := blob.PutAppendBlob(nil)

	if err != nil {
		return "", fmt.Errorf("Creating blob failed. %v", err)
	}
	// write contents to blob
	err = blob.AppendBlock(*content, nil)
	if err != nil {
		return "", fmt.Errorf("Cannot Add contents to blob. %v", err)
	}
	return blob.GetURL(), nil
}
