package helpers

import (
	"log"
)

// HandleErrors function handling the errors by using log.Fetal() function
func HandleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
