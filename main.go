package main

import (
	"fmt"
	"split-payment/service"
)

const (
	PATH_SHOPPING_LIST_FILE = "resources/shoppingList.json"
	PATH_EMAILS_LIST_FILE   = "resources/emails.json"
)

var (
	readService       = service.NewReadService(service.NewConvertService())
	calculatorService = service.NewCalculatorService(service.NewConvertService())
)

func main() {
	shoppingList, err := readService.ReadShoppingListFile(PATH_SHOPPING_LIST_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	emailsList, err := readService.ReadEmailsListFile(PATH_EMAILS_LIST_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	splitPayment, err := calculatorService.CalcPaymentByEmail(shoppingList, emailsList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(splitPayment)
}
