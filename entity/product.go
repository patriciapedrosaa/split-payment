package entity

type Product struct {
	DescriptionItem string  `json:"item"`
	Quantity        float32 `json:"quantity"`
	Price           float32 `json:"price"`
}