package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	// Other string `json:"-"` the - makes the parser ignore that field
}

func Marshal() {
	fmt.Println("---------- MARSHAL ----------")
	// json.Marshal() - converts from json to struct/map
	// json.Unmarshal() - converts from map/struct to json
	c := dog{"paçoca", "pitbull"}
	fmt.Println(c)

	dogJSON, err := json.Marshal(c) //Returns the json encoded
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dogJSON) //Print encoded json

	dogJSONHR := bytes.NewBuffer(dogJSON) //transform the encoded json to json format

	fmt.Println(dogJSONHR) //print json
	fmt.Println("-------------------------------")
	fmt.Println(" ")
}

func Unmarshal() {
	fmt.Println("---------- UNMARSHAL ----------")
	dogJSON := `{"name":"paçoca", "race":"pitbull"}` //json

	var dogOutput dog //variable that will receive transformed json

	dogJSONSliceByte := []byte(dogJSON) //parse json to []byte

	if err := json.Unmarshal(dogJSONSliceByte, &dogOutput); err != nil { //operation to parse the []byte that represents the initial json to a struct
		log.Fatal(err)
	}
	fmt.Println(dogOutput)

	dogJSON2 := `{"name":"tobby", "race":"rotweiller"}`

	dogOutput2 := make(map[string]string)
	if err := json.Unmarshal([]byte(dogJSON2), &dogOutput2); err != nil { //parsing to a map
		log.Fatal(err)
	}
	fmt.Println(dogOutput2)
	fmt.Println("-------------------------------")
	fmt.Println(" ")
}

func main() {
	Marshal()
	Unmarshal()
}
