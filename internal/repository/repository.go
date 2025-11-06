package repository

import (
	"database/sql"
	"errors"

	"github.com/Romchick1/testGO/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Products
func (r *Repository) GetAllProducts() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT id, name, quantity, unit_cost, measure_id FROM public.products")
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

func (r *Repository) GetProductByID(id int) (models.Product, error) {
	var p models.Product
	err := r.db.QueryRow("SELECT id, name, quantity, unit_cost, measure_id FROM public.products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID)
	if err == sql.ErrNoRows {
		return p, errors.New("product not found")
	}
	return p, err
}

func (r *Repository) CreateProduct(p models.Product) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO public.products (name, quantity, unit_cost, measure_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Quantity, p.UnitCost, p.MeasureID).Scan(&id)
	return id, err
}

func (r *Repository) UpdateProduct(id int, p models.Product) error {
	_, err := r.db.Exec("UPDATE public.products SET name=$1, quantity=$2, unit_cost=$3, measure_id=$4 WHERE id=$5",
		p.Name, p.Quantity, p.UnitCost, p.MeasureID, id)
	return err
}

func (r *Repository) DeleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM public.products WHERE id=$1", id)
	return err
}

func (r *Repository) GetAllMeasures() ([]models.Measure, error) {
	rows, err := r.db.Query("SELECT id, name FROM public.measures")
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

func (r *Repository) GetMeasureByID(id int) (models.Measure, error) {
	var m models.Measure
	err := r.db.QueryRow("SELECT id, name FROM public.measures WHERE id = $1", id).
		Scan(&m.ID, &m.Name)
	if err == sql.ErrNoRows {
		return m, errors.New("measure not found")
	}
	return m, err
}

func (r *Repository) CreateMeasure(m models.Measure) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO public.measures (name) VALUES ($1) RETURNING id", m.Name).Scan(&id)
	return id, err
}

func (r *Repository) UpdateMeasure(id int, m models.Measure) error {
	_, err := r.db.Exec("UPDATE public.measures SET name=$1 WHERE id=$2", m.Name, id)
	return err
}

func (r *Repository) DeleteMeasure(id int) error {
	_, err := r.db.Exec("DELETE FROM public.measures WHERE id=$1", id)
	return err
}

func (r *Repository) GetManagerByLogin(login string) (models.Manager, error) {
	var m models.Manager
	err := r.db.QueryRow("SELECT id, login, full_name FROM public.managers WHERE login = $1", login).
		Scan(&m.ID, &m.Login, &m.FullName)
	if err == sql.ErrNoRows {
		return m, errors.New("manager not found")
	}
	return m, err
}

func (r *Repository) GetAllManagers() ([]models.Manager, error) {
	rows, err := r.db.Query("SELECT id, login, full_name FROM public.managers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var managers []models.Manager
	for rows.Next() {
		var m models.Manager
		if err := rows.Scan(&m.ID, &m.Login, &m.FullName); err != nil {
			return nil, err
		}
		managers = append(managers, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return managers, nil
}

func (r *Repository) CreateManager(m models.Manager) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO public.managers (login, full_name) VALUES ($1, $2) RETURNING id",
		m.Login, m.FullName).Scan(&id)
	return id, err
}

func (r *Repository) UpdateManager(login string, m models.Manager) error {
	_, err := r.db.Exec("UPDATE public.managers SET full_name = $1 WHERE login = $2", m.FullName, login)
	return err
}

func (r *Repository) DeleteManager(login string) error {
	_, err := r.db.Exec("DELETE FROM public.managers WHERE login = $1", login)
	return err
}

func (r *Repository) GetProductsByManagerID(managerID int) ([]models.Product, error) {
	rows, err := r.db.Query(`
        SELECT p.id, p.name, p.quantity, p.unit_cost, p.measure_id 
        FROM public.products p 
        WHERE p.manager_id = $1`, managerID)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Repository) UpdateProductWithManagerCheck(id int, p models.Product, managerID int) error {
	_, err := r.db.Exec(`
        UPDATE public.products 
        SET name=$1, quantity=$2, unit_cost=$3, measure_id=$4 
        WHERE id=$5 AND manager_id=$6`,
		p.Name, p.Quantity, p.UnitCost, p.MeasureID, id, managerID)
	return err
}

func (r *Repository) DeleteProductWithManagerCheck(id int, managerID int) error {
	_, err := r.db.Exec("DELETE FROM public.products WHERE id=$1 AND manager_id=$2", id, managerID)
	return err
}
