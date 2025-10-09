package models

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitCost  float64 `json:"unit_cost"` // Исправили опечатку
	MeasureID int     `json:"measure"`   // ID меры
}

type Measure struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
