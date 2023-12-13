package models

//package

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CfgDBGetter interface {
	YAMLCfg(string)
	GetDB() *gorm.DB
}

type YAMLObject struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
	SSLM string `yaml:"sslmode"`
	DBNM string `yaml:"dbname"`
}

func (y *YAMLObject) YAMLCfg(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, y)
}

func (y *YAMLObject) GetDB() (db *gorm.DB) {
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", y.User, y.Pass, y.Host, y.Port, y.DBNM, y.SSLM)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func LoadCfgAndGetDB(yg CfgDBGetter, path string) (db *gorm.DB) {
	yg.YAMLCfg(path)
	db = yg.GetDB()
	return
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

type PackUnpacker interface {
	UnpackRequest(*http.Request)
	Pack(interface{}) []byte
}

// general method
func (r *RequestJSON) UnpackRequest(req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, r)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *RequestJSON) Pack(data interface{}) (pdata []byte) {
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	pdata = d
	return
}

func Unpacking(pu PackUnpacker, r *http.Request) {
	pu.UnpackRequest(r)
}

func Packing(pu PackUnpacker, data interface{}) (pdata []byte) {
	pdata = pu.Pack(data)
	return
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
