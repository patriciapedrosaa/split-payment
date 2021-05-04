package service_test

import (
	"split-payment/entity"
	"split-payment/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	calculatorService = service.NewCalculatorService(service.NewConvertService())
)

func Test_calcPaymentByEmail(t *testing.T) {
	tests := []struct {
		name         string
		shoppingList []entity.Product
		emails       entity.Emails
		wantErr      error
		wantResult   map[string]float32
	}{
		{
			name: "calculates payment by email successfully - Integer division",
			shoppingList: []entity.Product{
				{
					DescriptionItem: "Carne Bovina",
					Quantity:        1.5,
					Price:           30.0,
				},
				{
					DescriptionItem: "Carne Suína",
					Quantity:        1.5,
					Price:           30.0,
				},
				{
					DescriptionItem: "Bacon",
					Quantity:        1.5,
					Price:           30.0,
				},
			},
			emails: entity.Emails{
				Emails: []string{"pessoa1@email.com", "pessoa2@email.com", "pessoa3@email.com"},
			},
			wantErr:    nil,
			wantResult: map[string]float32{"pessoa1@email.com": 45.0, "pessoa2@email.com": 45.0, "pessoa3@email.com": 45.0},
		},
		{
			name: "calculates payment by email successfully - Decimal division",
			shoppingList: []entity.Product{
				{
					DescriptionItem: "Biscoito",
					Quantity:        1,
					Price:           0.50,
				},
				{
					DescriptionItem: "Bolacha",
					Quantity:        1,
					Price:           0.50,
				},
			},
			emails: entity.Emails{
				Emails: []string{"pessoa1@email.com", "pessoa2@email.com", "pessoa3@email.com"},
			},
			wantErr:    nil,
			wantResult: map[string]float32{"pessoa1@email.com": 0.34, "pessoa2@email.com": 0.33, "pessoa3@email.com": 0.33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			totalByEmail, err := calculatorService.CalcPaymentByEmail(tt.shoppingList, tt.emails)
			assert.Equal(tt.wantErr, err)
			assert.Equal(tt.wantResult, totalByEmail)
		})
	}
}
func Test_divisionByEmail(t *testing.T) {
	tests := []struct {
		name         string
		totalPayment int
		totalEmails  int
		emails       entity.Emails
		wantResult   map[string]int
	}{
		{
			name:         "integer Division",
			totalPayment: 3000,
			totalEmails:  3,
			emails: entity.Emails{
				Emails: []string{"pessoa1@email.com", "pessoa2@email.com", "pessoa3@email.com"},
			},
			wantResult: map[string]int{"pessoa1@email.com": 1000, "pessoa2@email.com": 1000, "pessoa3@email.com": 1000},
		},
		{
			name:         "decimal division",
			totalPayment: 100,
			totalEmails:  3,
			emails: entity.Emails{
				Emails: []string{"pessoa1@email.com", "pessoa2@email.com", "pessoa3@email.com"},
			},
			wantResult: map[string]int{"pessoa1@email.com": 34, "pessoa2@email.com": 33, "pessoa3@email.com": 33},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			totalValue := calculatorService.DivisionByEmail(tt.totalPayment, tt.totalEmails, tt.emails)
			assert.Equal(tt.wantResult, totalValue)
		})
	}
}
