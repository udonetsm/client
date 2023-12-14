package http

import (
	"fmt"

	"github.com/udonetsm/client/models"
)

func Create(target, name string, nums []string) {
	contact := &models.Contact{target, name, nums}
	object := &models.JSONObject{}
	models.Packing(object, contact)
	// call Create server function
}

func GetInfo(target string) {
	object := &models.JSONObject{}
	contact := &models.Contact{Number: target}
	models.Packing(object, contact)
	//call GetInfo server function
}

func Delete(target string) {
	object := &models.JSONObject{}
	contact := &models.Contact{Number: target}
	models.Packing(object, contact)
	// call Delete server function
}

func Upgrade(target, num, name string, nums []string) {
	contact := &models.Contact{num, name, nums}
	object := &models.JSONObject{Number: target}
	models.Packing(object, contact)
	fmt.Println(object)
	// find contact in db and change its general number
}
