package http

import (
	"fmt"

	"github.com/udonetsm/client/models"
)

func Create(target, name string, nums []string) {
	contact := &models.Contact{target, name, nums}
	object := &models.JSONObject{Number: target}
	pu := models.Packing(object, contact)
	fmt.Println(string(pu))
	// call Create server function
}

func DeleteOrInfo(target string) {
	object := &models.JSONObject{Number: target}
	pu := models.Packing(object, &models.Contact{})
	fmt.Println(string(pu))
}

func Upgrade(target, num, name string, nums []string) {
	contact := &models.Contact{num, name, nums}
	object := &models.JSONObject{Number: target}
	pu := models.Packing(object, contact)
	fmt.Println(string(pu))
	// find contact in db and change its info using JSONObject
}
