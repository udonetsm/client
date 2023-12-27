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
	Number string `json:"target"`
	// Object can be empty if using the DeleteOrInfo function.
	// See package github.com/udonetsm/client/http
	Object string `json:"object,omitempty"`
	Error  error  `gorm:"-" json:"omitempty"`
}

// Pack object to json string
func (j *Entries) Pack(contact *Contact) (data []byte) {
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
func (j *Entries) Unpack(data []byte) {
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
	Error      error    `json:"error,omitempty"`
}

func (c *Contact) Unpack(data []byte) {
	err := json.Unmarshal(data, c)
	if err != nil {
		log.Println(err)
		c.Error = err
		return
	}
}

type PackUnpackerContact interface {
	Unpack([]byte)
}

func UnpackingContact(p PackUnpackerContact, data []byte) {
	p.Unpack(data)
}

// Duck typing for json object
type PackUnpacker interface {
	Pack(*Contact) []byte
	// Unpack for use it on the server side
	// This func unpacking json on the server side
	Unpack([]byte)
}

// Use duck typing for pack
func Packing(pu PackUnpacker, c *Contact) (data []byte) {
	data = pu.Pack(c)
	return
}

// Use duck typing for unpack
func Unpacking(pu PackUnpacker, data []byte) {
	pu.Unpack(data)
}
