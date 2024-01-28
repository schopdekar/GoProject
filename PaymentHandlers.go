// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type PhonePePaymentInstrument struct {
	Type string `json:"type"`
}

// Define the request payload structure
type PhonePeRequest struct {
	MerchantId            string                   `json:"merchantId"`
	MerchantTransactionId string                   `json:"merchantTransactionId"`
	MerchantUserId        string                   `json:"merchantUserId"`
	Amount                string                   `json:"amount"`
	RedirectUrl           string                   `json:"redirectUrl"`
	RedirectMode          string                   `json:"redirectMode"`
	CallbackUrl           string                   `json:"callbackUrl"`
	MobileNumber          string                   `json:"mobileNumber"`
	PaymentInstrument     PhonePePaymentInstrument `json:"paymentInstrument"`
	// Add other required fields as per PhonePe API documentation
}

type PhonePeRequestObject struct {
	Request string `json:"request"`
}

// Define the response structure
type PhonePeResponse struct {
	// Define response fields based on PhonePe API documentation
}

type RequestBody struct {
	MerchantId            string `json:"merchantId"`
	MerchantTransactionId string `json:"merchantTransactionId"`
	MerchantUserId        string `json:"merchantUserId"`
	Amount                string `json:"amount"`
	MobileNumber          string `json:"mobileNumber"`
	// Add other required fields as per PhonePe API documentation
}

// PhonePe API endpoint
const phonePeEndpoint = "https://api-preprod.phonepe.com/apis/pg-sandbox/pg/v1/pay"

func InitiatePayment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(r.Body)
	var requestItem RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestItem)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Create a PhonePe request
	request := PhonePeRequest{
		MerchantId:            string(requestItem.MerchantId),
		MerchantTransactionId: string(requestItem.MerchantTransactionId),
		MerchantUserId:        string(requestItem.MerchantUserId),
		Amount:                string(requestItem.Amount),
		RedirectUrl:           "https://webhook.site/redirect-url",
		RedirectMode:          "REDIRECT",
		CallbackUrl:           "https://webhook.site/callback-url",
		MobileNumber:          string(requestItem.MobileNumber),
		PaymentInstrument: PhonePePaymentInstrument{
			Type: "PAY_PAGE",
		},
	}

	// fmt.Print(request)

	encodedRequest := EncodeRequestBody(request)
	// encodedRequest := "ewogICJtZXJjaGFudElkIjogIlBHVEVTVFBBWVVBVCIsCiAgIm1lcmNoYW50VHJhbnNhY3Rpb25JZCI6ICJNVDc4NTA1OTAwNjgxODgxMDQiLAogICJtZXJjaGFudFVzZXJJZCI6ICJNVUlEMTIzIiwKICAiYW1vdW50IjogMTAwMDAsCiAgInJlZGlyZWN0VXJsIjogImh0dHBzOi8vd2ViaG9vay5zaXRlL3JlZGlyZWN0LXVybCIsCiAgInJlZGlyZWN0TW9kZSI6ICJSRURJUkVDVCIsCiAgImNhbGxiYWNrVXJsIjogImh0dHBzOi8vd2ViaG9vay5zaXRlL2NhbGxiYWNrLXVybCIsCiAgIm1vYmlsZU51bWJlciI6ICI5OTk5OTk5OTk5IiwKICAicGF5bWVudEluc3RydW1lbnQiOiB7CiAgICAidHlwZSI6ICJQQVlfUEFHRSIKICB9Cn0="
	checksum := CreateChecksum(encodedRequest, "/pg/v1/pay", "099eb0cd-02cf-4e2a-8aca-3e6c6aff0399")
	// var requestObject = map[string]string{
	// 	"request": encodedRequest,
	// }

	requestObject := PhonePeRequestObject{
		Request: encodedRequest,
	}

	// re

	// jsonData, err := json.Marshal(requestObject)

	// Make a POST request to the PhonePe API
	client := resty.New()
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-VERIFY", checksum).
		SetBody(requestObject).
		Post(phonePeEndpoint)

	// fmt.Println(response, err)

	if err != nil {
		fmt.Println("Error making request to PhonePe API:", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
	}

	// Parse the response JSON
	// var phonePeResponse PhonePeResponse
	// if err := json.Unmarshal(response.Body(), &phonePeResponse); err != nil {
	// 	fmt.Println("Error decoding PhonePe API response:", err)
	// 	return
	// }

	var phonePeResponse PhonePeResponse
	json.Unmarshal(response.Body(), &phonePeResponse)

	// Process the PhonePe API response
	// Note: Handle the response according to PhonePe API documentation

	fmt.Println("PhonePe API Response:", phonePeResponse)
	//Sending Response
	json.NewEncoder(w).Encode(map[string]string{"data": string(response.Body())})
}

func redirectedUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirect function called")
	fmt.Println(r.Body)
	json.NewEncoder(w).Encode(map[string]string{"data": "User redirected successfully"})
}

func S2SCallbackForPayment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Call back function called")
	fmt.Println(r.Body)
	json.NewEncoder(w).Encode(map[string]string{"data": ""})

	// w.Header().Set("Content-Type", "application/json")
	// fmt.Println(r.Body)
	// var requestItem RequestBody
	// err := json.NewDecoder(r.Body).Decode(&requestItem)
	// if err != nil {
	// 	http.Error(w, "Error decoding JSON", http.StatusBadRequest)
	// 	return
	// }

	// // Create a PhonePe request
	// request := PhonePeRequest{
	// 	MerchantId:            string(requestItem.MerchantId),
	// 	MerchantTransactionId: string(requestItem.MerchantTransactionId),
	// 	MerchantUserId:        string(requestItem.MerchantUserId),
	// 	Amount:                string(requestItem.Amount),
	// 	RedirectUrl:           "https://webhook.site/redirect-url",
	// 	RedirectMode:          "REDIRECT",
	// 	CallbackUrl:           "https://webhook.site/callback-url",
	// 	MobileNumber:          string(requestItem.MobileNumber),
	// 	PaymentInstrument: PhonePePaymentInstrument{
	// 		Type: "PAY_PAGE",
	// 	},
	// }

	// // fmt.Print(request)

	// encodedRequest := EncodeRequestBody(request)
	// // encodedRequest := "ewogICJtZXJjaGFudElkIjogIlBHVEVTVFBBWVVBVCIsCiAgIm1lcmNoYW50VHJhbnNhY3Rpb25JZCI6ICJNVDc4NTA1OTAwNjgxODgxMDQiLAogICJtZXJjaGFudFVzZXJJZCI6ICJNVUlEMTIzIiwKICAiYW1vdW50IjogMTAwMDAsCiAgInJlZGlyZWN0VXJsIjogImh0dHBzOi8vd2ViaG9vay5zaXRlL3JlZGlyZWN0LXVybCIsCiAgInJlZGlyZWN0TW9kZSI6ICJSRURJUkVDVCIsCiAgImNhbGxiYWNrVXJsIjogImh0dHBzOi8vd2ViaG9vay5zaXRlL2NhbGxiYWNrLXVybCIsCiAgIm1vYmlsZU51bWJlciI6ICI5OTk5OTk5OTk5IiwKICAicGF5bWVudEluc3RydW1lbnQiOiB7CiAgICAidHlwZSI6ICJQQVlfUEFHRSIKICB9Cn0="
	// checksum := CreateChecksum(encodedRequest, "/pg/v1/pay", "099eb0cd-02cf-4e2a-8aca-3e6c6aff0399")
	// // var requestObject = map[string]string{
	// // 	"request": encodedRequest,
	// // }

	// requestObject := PhonePeRequestObject{
	// 	Request: encodedRequest,
	// }

	// // re

	// // jsonData, err := json.Marshal(requestObject)

	// // Make a POST request to the PhonePe API
	// client := resty.New()
	// response, err := client.R().
	// 	SetHeader("Content-Type", "application/json").
	// 	SetHeader("X-VERIFY", checksum).
	// 	SetBody(requestObject).
	// 	Post(phonePeEndpoint)

	// fmt.Println(response, err)

	// if err != nil {
	// 	fmt.Println("Error making request to PhonePe API:", err)
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
	// }

	// // Parse the response JSON
	// // var phonePeResponse PhonePeResponse
	// // if err := json.Unmarshal(response.Body(), &phonePeResponse); err != nil {
	// // 	fmt.Println("Error decoding PhonePe API response:", err)
	// // 	return
	// // }

	// var phonePeResponse PhonePeResponse
	// json.Unmarshal(response.Body(), &phonePeResponse)

	// // Process the PhonePe API response
	// // Note: Handle the response according to PhonePe API documentation

	// fmt.Println("PhonePe API Response:", phonePeResponse)
	//Sending Response
	// json.NewEncoder(w).Encode(map[string]string{"data": string(response.Body())})
}
