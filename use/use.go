package use

import (
	"fmt"

	"github.com/udonetsm/client/http"
)

func Create(number, name string, numlist []string) {
	http.CallCreateContact(number, name, numlist)
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
