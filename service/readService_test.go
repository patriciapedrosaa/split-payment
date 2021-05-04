package service_test

import (
	"errors"
	"split-payment/entity"
	"split-payment/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	readService = service.NewReadService(service.NewConvertService())
)

func Test_readShoppingListFile(t *testing.T) {
	tests := []struct {
		name         string
		shoppingList string
		wantErr      error
		wantResult   []entity.Product
	}{
		{
			name:         "Product valid",
			shoppingList: "../resources/mock/shoppingList_valid.json",
			wantErr:      nil,
			wantResult: []entity.Product{
				{DescriptionItem: "biscoito", Quantity: 1, Price: 2.50},
				{DescriptionItem: "banana", Quantity: 1, Price: 2.50},
			},
		},
		{
			name:         "Description invalid",
			shoppingList: "../resources/mock/shoppingList_description_error.json",
			wantErr:      errors.New("description item: Description Item invalid, please, check the list"),
			wantResult:   nil,
		},
		{
			name:         "Quantity is zero",
			shoppingList: "../resources/mock/shoppingList_quantity_zero.json",
			wantErr:      errors.New("quantity: Quantity is invalid, please check the list"),
			wantResult:   nil,
		},
		{
			name:         "Quantity is negative",
			shoppingList: "../resources/mock/shoppingList_quantity_negative.json",
			wantErr:      errors.New("quantity: Quantity is invalid, please check the list"),
			wantResult:   nil,
		},
		{
			name:         "Price is zero",
			shoppingList: "../resources/mock/shoppingList_price_zero.json",
			wantErr:      errors.New("price: Price is invalid, please, check the list"),
			wantResult:   nil,
		},
		{
			name:         "Price is negative",
			shoppingList: "../resources/mock/shoppingList_price_negative.json",
			wantErr:      errors.New("price: Price is invalid, please, check the list"),
			wantResult:   nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			totalByEmail, err := readService.ReadShoppingListFile(tt.shoppingList)
			assert.Equal(tt.wantErr, err)
			assert.Equal(tt.wantResult, totalByEmail)
		})
	}
}
func Test_readEmailsListFile(t *testing.T) {
	tests := []struct {
		name       string
		emails     string
		wantErr    error
		wantResult entity.Emails
	}{
		{
			name:    "Email list valid",
			emails:  "../resources/mock/emails_valid.json",
			wantErr: nil,
			wantResult: entity.Emails{
				Emails: []string{
					"pessoa1@email.com",
					"pessoa2@email.com",
					"pessoa3@email.com",
					"pessoa4@emails.com",
					"pessoa5@email.com",
					"pessoa6@email.com",
					"pessoa7@emails.com",
					"pessoa8@email.com",
					"pessoa9@email.com",
				},
			},
		},
		{
			name:       "Email list is invalid",
			emails:     "../resources/mock/emails_empty.json",
			wantErr:    errors.New("email: email invalid, please, check the list"),
			wantResult: entity.Emails{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			totalByEmail, err := readService.ReadEmailsListFile(tt.emails)
			assert.Equal(tt.wantErr, err)
			assert.Equal(tt.wantResult, totalByEmail)
		})
	}
}