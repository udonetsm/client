package use

import (
	"fmt"

	"github.com/udonetsm/client/http"
	"github.com/udonetsm/client/models"
)

func Create(number, name string, numlist []string) {
	contact := models.NewContact(number, name, numlist)
	obj := models.NewObjetToPackJSON(number, contact)
	obj.Pack(contact)
	//call remote function by http using restAPI
	//pass obj.Object to it
}

func UpdateNumber(target, newNumber string) {
	//call function from remote server by restAPI
	http.CallUpdateNumber(target, newNumber)
	fmt.Printf("[%s] changed to [%s]\n", target, newNumber)
}

func UpdateName(target, newName string) {
	//call function from remote server by restAPI
	fmt.Printf("name [%s] set for [%s]\n", newName, target)
}

func UpdateNumList(target, object string) {
	//get contact
}

func Delete(target string) {
	//call function from remote server by restAPI
	fmt.Printf("Contact [%s] has been deleted\n", target)
}
