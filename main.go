package main

import (
	"fmt"

	"github.com/hamedhaghi/card-scheme/utils"
)

func main() {
	cardNumber := "5105105105105100"
	isValid := utils.IsCardNumberValid(cardNumber)
	scheme := utils.GetCardScheme(cardNumber)
	fmt.Printf("%t, %s", isValid, scheme)

}
