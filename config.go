package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	dbDriver             string `json:"dbDriver"`
	dbUser               string `json:"dbUser"`
	dbPass               string `json:"dbPass"`
	dbName               string `json:"dbName"`
	fileSize             string `json:"fileSize"`
	filesFromRequest     string `json:"filesFromRequest"`
	accountIdFromRequest string `json:"accountIdFromRequest"`
	recordIdFromRequest  string `json:"recordIdFromRequest"`
}

func initConfiguration() Configuration {
	fmt.Println("Initialise configuration")

	jsonFile, err := os.Open("conf.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	configuration1 = Configuration{}
	configuration1.dbDriver = result["configuration"].(map[string]interface{})["dbDriver"].(string)
	configuration1.dbUser = result["configuration"].(map[string]interface{})["dbUser"].(string)
	configuration1.dbPass = result["configuration"].(map[string]interface{})["dbPass"].(string)
	configuration1.dbName = result["configuration"].(map[string]interface{})["dbName"].(string)
	configuration1.fileSize = result["configuration"].(map[string]interface{})["fileSize"].(string)
	configuration1.filesFromRequest = result["configuration"].(map[string]interface{})["filesFromRequest"].(string)
	configuration1.accountIdFromRequest = result["configuration"].(map[string]interface{})["accountIdFromRequest"].(string)
	configuration1.recordIdFromRequest = result["configuration"].(map[string]interface{})["recordIdFromRequest"].(string)
	return configuration1
}
