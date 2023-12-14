package http

import (
	"github.com/udonetsm/client/models"
)

func Create(target, name string, nums []string) {
	contact := &models.Contact{target, name, nums}
	object := &models.JSONObject{Number: target}
	models.Packing(object, contact)
	// call Create server function
}

func GetInfo(target string) {
	object := &models.JSONObject{Number: target}
	models.Packing(object, &models.Contact{})
	//call GetInfo server function
}

func Delete(target string) {
	object := &models.JSONObject{Number: target}
	models.Packing(object, &models.Contact{})
}

func Upgrade(target, num, name string, nums []string) {
	contact := &models.Contact{num, name, nums}
	object := &models.JSONObject{Number: target}
	models.Packing(object, contact)
	// find contact in db and change its info using JSONObject
}
