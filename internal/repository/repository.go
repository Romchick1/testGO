package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Romchick1/testGO/internal/models"
)

func GetAllProducts() ([]models.Product, error) {
	log.Println("Executing query: SELECT id, name, quantity, unit_cost, measure_id FROM public.products")
	rows, err := DB.Query("SELECT id, name, quantity, unit_cost, measure_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	var p models.Product
	log.Println("Executing query: SELECT id, name, quantity, unit_cost, measure_id FROM products WHERE id = $1")
	err := DB.QueryRow("SELECT id, name, quantity, unit_cost, measure_id FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID)
	if err == sql.ErrNoRows {
		return p, errors.New("product not found")
	}
	return p, err
}

func CreateProduct(p models.Product) (int, error) {
	var id int
	err := DB.QueryRow("INSERT INTO products (name, quantity, unit_cost, measure_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Quantity, p.UnitCost, p.MeasureID).Scan(&id)
	return id, err
}

func UpdateProduct(id int, p models.Product) error {
	_, err := DB.Exec("UPDATE products SET name=$1, quantity=$2, unit_cost=$3, measure_id=$4 WHERE id=$5",
		p.Name, p.Quantity, p.UnitCost, p.MeasureID, id)
	return err
}

func DeleteProduct(id int) error {
	_, err := DB.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

func GetAllMeasures() ([]models.Measure, error) {
	rows, err := DB.Query("SELECT id, name FROM measures")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measures []models.Measure
	for rows.Next() {
		var m models.Measure
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			return nil, err
		}
		measures = append(measures, m)
	}
	return measures, nil
}

func GetMeasureByID(id int) (models.Measure, error) {
	var m models.Measure
	err := DB.QueryRow("SELECT id, name FROM measures WHERE id = $1", id).
		Scan(&m.ID, &m.Name)
	if err == sql.ErrNoRows {
		return m, errors.New("measure not found")
	}
	return m, err
}

func CreateMeasure(m models.Measure) (int, error) {
	var id int
	err := DB.QueryRow("INSERT INTO measures (name) VALUES ($1) RETURNING id", m.Name).Scan(&id)
	return id, err
}

func UpdateMeasure(id int, m models.Measure) error {
	_, err := DB.Exec("UPDATE measures SET name=$1 WHERE id=$2", m.Name, id)
	return err
}

func DeleteMeasure(id int) error {
	_, err := DB.Exec("DELETE FROM measures WHERE id=$1", id)
	return err
}
