package service

const CENTS = 100

type ConvertService interface {
	ConvertToInt(value float32) int
	ConvertToHashMapFloat(total map[string]int) map[string]float32
}

type convertService struct{}

func NewConvertService() ConvertService {
	return &convertService{}
}

func (s *convertService) ConvertToInt(value float32) int {
	valueInt := value * CENTS

	return int(valueInt)
}

func (s *convertService) ConvertToHashMapFloat(total map[string]int) map[string]float32 {
	totalFloat := map[string]float32{}
	for email, value := range total {
		totalFloat[email] = float32(value) / CENTS
	}
	return totalFloat
}
