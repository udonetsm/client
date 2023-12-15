package http

import (
	"fmt"

	"github.com/udonetsm/client/models"
)

// Create get json with number, name and additional numbers
// and call create function on the server side using http
func Create(target, name string, nums []string) {
	contact := &models.Contact{target, name, nums}
	object := &models.JSONObject{Number: target}
	pu := models.Packing(object, contact)
	fmt.Println(string(pu))
	// call Create server function
}

// DeleteOrInfo get json with target number only
// and send it to the server side.
// if current command line command is delete
// this func call delete function using http on the server side
// if command line command is info
// this func call info on the server side using http and
// get full info about target contact
func DeleteOrInfo(target string) {
	object := &models.JSONObject{Number: target}
	// needs only target number. Contact should be empty
	pu := models.Packing(object, &models.Contact{})
	fmt.Println(string(pu))
}

// Upgrade get json with target contact, upgradable
// unit and new value of upgradable unit
// and call update func on the server side using http
// this func can update all of part some contact
func Upgrade(target, num, name string, nums []string) {
	// Contact includes only one field.
	// It set during type command line command
	contact := &models.Contact{num, name, nums}
	object := &models.JSONObject{Number: target}
	pu := models.Packing(object, contact)
	fmt.Println(string(pu))
	// find contact in db and change its info using JSONObject
}
