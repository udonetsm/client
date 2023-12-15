package use

import (
	"errors"
	"log"
)

// LimitNumList checks len of additional number list
// and ask user to type correctly amount of additional
// number if list is too long
func LimitNumList(numlist []string) {
	if len(numlist) > 3 {
		log.Fatal(errors.New("Numlist too long. Maximum elements in the numlist is 3"))
	}
}
