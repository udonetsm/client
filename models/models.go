package models

// This package can be imported as github.com/udonetsm/client/models.
// Server side uses it

import (
	"encoding/json"
	"log"
)

// JSON object for making request to server side
// includes:
// target for fill entry_id in database
// Contact for build json string for use functions on the server side
type Entries struct {
	Number string `json:"number"`
	// Object can be empty if using the DeleteOrInfo function.
	// See package github.com/udonetsm/client/http
	Object string `json:"object,omitempty"`
	Error  error  `gorm:"-" json:"omitempty"`
}

// Pack object to json string
func (j *Entries) PackEntries(contact *Contact) (data []byte) {
	data, err := json.Marshal(contact)
	if err != nil {
		log.Println(err)
		j.Error = err
		return
	}
	j.Object = string(data)
	data, err = json.Marshal(j)
	if err != nil {
		log.Println(err)
		j.Error = err
		return
	}
	return
}

// Unpack object from json string to JSONObject struct
func (j *Entries) UnpackEntries(data []byte) {
	err := json.Unmarshal(data, j)
	if err != nil {
		log.Println(err)
		j.Error = err
		return
	}
}

// Contact object
type Contact struct {
	Number     string   `json:"num,omitempty"`
	Name       string   `json:"name,omitempty"`
	NumberList []string `json:"nlist,omitempty"`
}

func (c *Contact) UnpackContact(e *Entries) {
	err := json.Unmarshal([]byte(e.Object), c)
	if err != nil {
		log.Println(err)
		e.Error = err
		return
	}
}

type PackUnpackerContact interface {
	Unpack(*Entries)
}

func UnpackingContact(p PackUnpackerContact, e *Entries) {
	p.Unpack(e)
}
