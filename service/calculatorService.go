package service

import (
	"split-payment/entity"
)

type CalculatorService interface {
	CalcPaymentByEmail(list []entity.Product, emails entity.Emails) (map[string]float32, error)
	DivisionByEmail(totalPayment int, totalEmails int, emails entity.Emails) map[string]int
	TotalPayment(list []entity.Product) int
	TotalOfEmails(emails entity.Emails) int
}

type calculatorService struct {
	convertService ConvertService
}

func NewCalculatorService(convertService ConvertService) CalculatorService{
	return &calculatorService{convertService: convertService}			
}

func (s *calculatorService) CalcPaymentByEmail(list []entity.Product, emails entity.Emails) (map[string]float32, error) {
	totalPayment := s.TotalPayment(list)
	totalEmails := s.TotalOfEmails(emails)
	paymentByEmail := s.DivisionByEmail(totalPayment, totalEmails, emails)

	totalFloat := s.convertService.ConvertToHashMapFloat(paymentByEmail)
	return totalFloat, nil
}
func (s *calculatorService) DivisionByEmail(totalPayment int, totalEmails int, emails entity.Emails) (map[string]int) {
	totalPaymentByEmail := totalPayment / totalEmails
	rest := totalPayment % totalEmails

	totalByEmail := make(map[string]int)
	for _, email := range emails.Emails {
		totalByEmail[email] = totalPaymentByEmail
		if rest > 0{
			totalByEmail[email] = totalPaymentByEmail + 1
			rest -= 1
		}
	}
	return totalByEmail
}
func (s *calculatorService) TotalPayment(list []entity.Product) int {
	var total = 0
	for _, item := range list {
		quantity := item.Quantity
		price := item.Price
		value := s.convertService.ConvertToInt(quantity * price)
		
		total += value
	}
	return total
}
func (s *calculatorService) TotalOfEmails(emails entity.Emails) int {
	totalOfEmails := len(emails.Emails)
	return totalOfEmails
}