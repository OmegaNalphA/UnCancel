package main

import (
	"UnCancel/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func openJSON() structs.FBAccessToken {
	var access structs.FBAccessToken

	jsonFile, jsonErr := os.Open("access_token.json")
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	// Open our jsonFile
	byteval, byteErr := ioutil.ReadAll(jsonFile)
	if byteErr != nil {
		fmt.Println(byteErr)
	}

	json.Unmarshal(byteval, &access)

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()
	return access
	// defer the closing of our jsonFile so that we can parse it later on
}

func main() {
	fmt.Println("hello")
	access := openJSON()
	fmt.Println(access.FBID)
}
