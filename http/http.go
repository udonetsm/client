package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/udonetsm/server/models"
)

// Create get json with number, name and additional numbers
// and call create function on the server side using http
func Create(target, name string, nums []string) {
	contact := &models.Contact{target, name, nums}
	object := &models.Entries{Number: target}
	pu, err := models.PackingEntries(object, contact)
	if err != nil {
		log.Fatal(err)
		return
	}
	DoReq("http://localhost:8080", "/create", http.MethodPost, pu)
	// call Create server function
}

// Delete get json with target number only
// and send it to the server side.
// this func call delete function using http on the server side
func Delete(target string) {
	object := &models.Entries{Number: target}
	// needs only target number. Contact should be empty
	pu, err := models.PackingEntries(object, &models.Contact{})
	if err != nil {
		log.Fatal(err)
		return
	}
	DoReq("http://localhost:8080", "/delete", http.MethodPost, pu)
}

// this func call info function using http on the server side
func Info(target string) {
	// needs only target number. Contact should be empty
	pu, err := models.PackingEntries(&models.Entries{Number: target}, &models.Contact{})
	if err != nil {
		log.Fatal(err)
		return
	}
	DoReq("http://localhost:8080", "/info", http.MethodPost, pu)
}

// Upgrade get json with target contact, upgradable
// unit and new value of upgradable unit
// and call update func on the server side using http
// this func can update all of part some contact
func Upgrade(target, upgradable, num, name string, nums []string) {
	// Contact includes only one field.
	// It set during type command line command
	contact := &models.Contact{num, name, nums}
	object := &models.Entries{Number: target}
	pu, err := models.PackingEntries(object, contact)
	if err != nil {
		log.Fatal(err)
		return
	}
	DoReq("http://localhost:8080", fmt.Sprintf("/update/%s", upgradable), http.MethodPost, pu)
	// find contact in db and change its info using JSONObject
}

func Search(name string) {
	contact := &models.Contact{Name: name}
	object := &models.Entries{}
	pu, err := models.PackingEntries(object, contact)
	if err != nil {
		log.Fatal(err)
		return
	}
	DoReq("http://localhost:8080", "/search", http.MethodPost, pu)
}

// making request and get result from server side
// url example <http://localhost:8080>
// uri example </targetfunction>
func DoReq(url, uri, method string, body []byte) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", url, uri), bytes.NewBuffer(body))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	// do request on the server side
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// read answer from server
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
