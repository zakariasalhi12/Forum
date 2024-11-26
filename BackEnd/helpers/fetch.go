package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Fetch(url string, data any) error {
	// Make the GET request
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error making GET request:", err)
		return err
	}
	defer res.Body.Close()
	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}
	// Unmarshal the JSON response into the provided data structure
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Println("Error unmarshaling JSON response:", err)
		return err
	}
	return nil
}
