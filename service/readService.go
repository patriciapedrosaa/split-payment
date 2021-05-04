package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"split-payment/entity"
)

type ReadService interface {
	ReadShoppingListFile(document string) ([]entity.Product, error)
	ReadEmailsListFile(document string) (entity.Emails, error)
	ValidatesShoppingList(shoppingList []entity.Product) error
	ValidatesEmailsList(emails entity.Emails) error
}

type readService struct {
	convertService ConvertService
}

func NewReadService(convertService ConvertService) ReadService {
	return &readService{convertService: convertService}
}

func (s *readService) ReadShoppingListFile(document string) ([]entity.Product, error) {
	jsonFile, err := os.Open(document)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValueJSON, _ := ioutil.ReadAll(jsonFile)
	var objProduct []entity.Product
	err = json.Unmarshal(byteValueJSON, &objProduct)
	if err != nil {
		return nil, err
	}

	err = s.ValidatesShoppingList(objProduct)
	if err != nil {
		return nil, err
	}

	return objProduct, nil
}
func (s *readService) ReadEmailsListFile(document string) (entity.Emails, error) {
	jsonFile, err := os.Open(document)
	if err != nil {
		return entity.Emails{}, err
	}

	defer jsonFile.Close()
	byteValueJSON, _ := ioutil.ReadAll(jsonFile)
	objEmails := entity.Emails{}

	err = json.Unmarshal(byteValueJSON, &objEmails)
	if err != nil {
		return entity.Emails{}, err
	}

	err = s.ValidatesEmailsList(objEmails)
	if err != nil {
		return entity.Emails{}, err
	}
	return objEmails, nil
}
func (s *readService) ValidatesShoppingList(shoppingList []entity.Product) error {
	for i := 0; i < len(shoppingList); i++ {
		if shoppingList[i].DescriptionItem == "" {
			return errors.New("description item: Description Item invalid, please, check the list")
		}
		price := s.convertService.ConvertToInt(shoppingList[i].Price)
		if price <= 0 {
			return errors.New("price: Price is invalid, please, check the list")
		}

		quantity := shoppingList[i].Quantity
		if quantity <= 0 {
			return errors.New("quantity: Quantity is invalid, please check the list")
		}
	}
	return nil
}
func (s *readService) ValidatesEmailsList(emails entity.Emails) error {
	for i := 0; i < len(emails.Emails); i++ {
		if emails.Emails[i] == "" {
			return errors.New("email: email invalid, please, check the list")
		}
	}
	return nil
}
