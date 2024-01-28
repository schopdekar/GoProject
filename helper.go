package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func EncodeRequestBody(p PhonePeRequest) string {
	// Create an instance of the struct
	// data := ExampleStruct{
	// 	ID:      1,
	// 	Name:    "Example",
	// 	Details: "Some details about the example.",
	// }

	// Convert struct to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return "Error occured"
	}

	// Encode JSON as base64
	base64Encoded := base64.StdEncoding.EncodeToString(jsonData)

	// fmt.Println("Original Struct:", data)
	fmt.Println("JSON Representation:", string(jsonData))
	fmt.Println("Base64 Encoded:", base64Encoded)
	return base64Encoded
}

func GenerateSHA256(input string) string {
	// Create a new SHA-256 hash
	hasher := sha256.New()

	// Write the input string to the hash
	hasher.Write([]byte(input))

	// Get the final hash sum
	hashSum := hasher.Sum(nil)

	// Convert the hash sum to a hex-encoded string
	hashString := hex.EncodeToString(hashSum)

	return hashString
}

func CreateChecksum(encodedBody, apiEndPoint, salt string) string {
	sha256Value := GenerateSHA256(encodedBody + apiEndPoint + salt)
	return sha256Value + "###" + "1"

}
