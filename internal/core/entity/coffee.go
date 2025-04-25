package entity

type Coffee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Rating     float64 `json:"rating"`
	Reviews    int     `json:"reviews"`
	PriceRange string  `json:"price_range"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
	ReviewText string  `json:"review_text"`
}
