package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Jsonfilecreation(apiUri,filePath string) error {
	// API call
	client := &http.Client{}
	req, err := http.NewRequest("GET",apiUri , nil)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Get API Response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	defer resp.Body.Close()

	// Read Body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	// Unmarshal data into interface
	var responseObject interface{}
	json.Unmarshal(bodyBytes, &responseObject)

	// Create response json file
	file, _ := json.MarshalIndent(responseObject, "", " ")
	
	err = os.WriteFile(filePath+"/DogDetails_"+time.Now().Format("20060102_150405")+".json", file, 0644)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func main() {
	Jsonfilecreation("https://http.dog/200.json","./JsonOutput")
}
