package models

//package

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v2"
)

type YAMLObject struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
	SSLM string `yaml:"sslmode"`
	DBNM string `yaml:"dbname"`
}

type Contact struct {
	Number     string   `json:"num"`
	Name       string   `json:"name,omitempty"`
	NumberList []string `json:"nlist,omitempty"`
}

type JSONObject struct {
	Number string
	Object string
}

type RequestJSON struct {
	Target  string `json:"target"`
	Upgrade string `json:"newdata"`
}

// Method for unpack yaml testing
func (y *YAMLObject) Unpack(data []byte) {
	yaml.Unmarshal(data, y)
}

func (o *JSONObject) Pack(c *Contact) {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	o.Object = string(data)
}

func NewContact(number, name string, numlist []string) *Contact {
	return &Contact{number, name, numlist}
}

func NewObjetToPackJSON(number string, contact *Contact) *JSONObject {
	return &JSONObject{Number: number}
}
