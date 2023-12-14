package models

//package

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type YAMLObject struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
	SSLM string `yaml:"sslmode"`
	DBNM string `yaml:"dbname"`
}

type CfgDBGetter interface {
	YAMLCfg(string)
	GetDB() *gorm.DB
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
	Number string `json:"target"`
	Object string `json:"object,omitempty"`
}

type PackUnpacker interface {
	Pack(*Contact)
	Unpack(*Contact)
}

func (j *JSONObject) Pack(c *Contact) {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	j.Object = string(data)
}

func (j *JSONObject) Unpack(c *Contact) {
	err := json.Unmarshal([]byte(j.Object), c)
	if err != nil {
		log.Fatal(err)
	}
}

func Packing(pu PackUnpacker, c *Contact) {
	pu.Pack(c)
}

func Unpacking(pu PackUnpacker, c *Contact) {
	pu.Unpack(c)
}
