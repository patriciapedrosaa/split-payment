package service_test

import (
	"split-payment/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var convertService = service.NewConvertService()

func Test_convertToInt(t *testing.T) {
	tests := []struct {
		name       string
		value      float32
		wantResult int
	}{
		{
			name:       "successfully converts to integer",
			value:      25.50,
			wantResult: 2550,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			valueInt := convertService.ConvertToInt(tt.value)
			assert.Equal(tt.wantResult, valueInt)
		})
	}
}
func Test_convertToHashMap(t *testing.T) {
	tests := []struct {
		name       string
		value      map[string]int
		wantResult map[string]float32
	}{
		{
			name:       "successfully convert to hash map",
			value:     map[string]int{"pessoa1@email.com": 4500, "pessoa2@email.com": 4500, "pessoa3@email.com": 4500},
			wantResult: map[string]float32{"pessoa1@email.com": 45.0, "pessoa2@email.com": 45.0, "pessoa3@email.com": 45.0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			hashMapFloat := convertService.ConvertToHashMapFloat(tt.value)
			assert.Equal(tt.wantResult, hashMapFloat)
		})
	}
}
