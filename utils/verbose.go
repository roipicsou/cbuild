package utils

import "fmt"

func VerPrint(estVerb bool, dernierAppel bool, msg string) {
	if(estVerb) {
		fmt.Println(msg)
	} else if( dernierAppel) {
		fmt.Println(".")
	} else {
		fmt.Print(".")
	}
}