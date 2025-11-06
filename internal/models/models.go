package models

type Manager struct {
	ID       int    `json:"-"`
	Login    string `json:"login"`
	FullName string `json:"full_name"`
}

type ManagerResponse struct {
	Login    string `json:"login"`
	FullName string `json:"full_name"`
}

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitCost  float64 `json:"unit_cost"`
	MeasureID int     `json:"measure"`
	ManagerID int     `json:"-"`
}

type Measure struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
