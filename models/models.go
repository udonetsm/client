package models

// This package can be imported as github.com/udonetsm/client/models.
// Server side uses it

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// this object for load database connection config
type YAMLObject struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
	SSLM string `yaml:"sslmode"`
	DBNM string `yaml:"dbname"`
}

// duck typing for load data base connection config
type CfgDBGetter interface {
	YAMLCfg(string)
	GetDB() *gorm.DB
}

// this method read config from target .yaml file and unpack it in object
func (y *YAMLObject) YAMLCfg(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, y)
}

// build database connection string
// using object built on YAMLCfg function
// and get database usin built config
func (y *YAMLObject) GetDB() (db *gorm.DB) {
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", y.User, y.Pass, y.Host, y.Port, y.DBNM, y.SSLM)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	return
}

// using duck typing for load database connection
func LoadCfgAndGetDB(yg CfgDBGetter, path string) (db *gorm.DB) {
	yg.YAMLCfg(path)
	db = yg.GetDB()
	return
}

// Contact object
type Contact struct {
	Number     string   `json:"num,omitempty"`
	Name       string   `json:"name,omitempty"`
	NumberList []string `json:"nlist,omitempty"`
}

// JSON object for making request to server side
// includes:
// target for fill entry_id in database
// Contact for build json string for use functions on the server side
type JSONObject struct {
	Number string `json:"target"`
	// Object can be empty if using the DeleteOrInfo function.
	// See package github.com/udonetsm/client/http
	Object *Contact `json:"object,omitempty"`
}

// Duck typing for json object
type PackUnpacker interface {
	Pack() []byte
	// Unpack for use it on the server side
	// This func unpacking json on the server side
	Unpack()
}

// Pack object to json string
func (j *JSONObject) Pack() (data []byte) {
	data, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Unpack object from json string to JSONObject struct
func (j *JSONObject) Unpack(data []byte) {
	err := json.Unmarshal(data, j)
	if err != nil {
		log.Fatal(err)
	}
}

// Use duck typing for pack
func Packing(pu PackUnpacker) (data []byte) {
	data = pu.Pack()
	return
}

// Use duck typing for unpack
func Unpacking(pu PackUnpacker) {
	pu.Unpack()
	return
}
