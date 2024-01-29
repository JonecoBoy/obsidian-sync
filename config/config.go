package config

import(
	"io/ioutil"
	"encoding/json"
	"log"
)

type Config struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetConfig() (Config ,error){
	fileContent, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}
	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	return config,err

}