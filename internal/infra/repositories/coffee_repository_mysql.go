package repositories

import (
	"database/sql"
	"localgems/internal/core/entity"
)

type MySQLCoffeeRepository struct {
	DB *sql.DB
}

func NewMySQLCoffeeRepository(db *sql.DB) CoffeeRepository {
	return &MySQLCoffeeRepository{DB: db}
}

// FindAll - Lấy toàn bộ cafe
func (r *MySQLCoffeeRepository) FindAll() ([]entity.Coffee, error) {
	rows, err := r.DB.Query("SELECT id, name, rating, reviews, price_range, type, address, review_text FROM coffees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []entity.Coffee
	for rows.Next() {
		var coffee entity.Coffee
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Rating,
			&coffee.Reviews,
			&coffee.PriceRange,
			&coffee.Type,
			&coffee.Address,
			&coffee.ReviewText,
		)
		if err != nil {
			return nil, err
		}
		coffees = append(coffees, coffee)
	}

	return coffees, nil
}

// FindByID - Tìm cafe theo ID
func (r *MySQLCoffeeRepository) FindByID(id int) (*entity.Coffee, error) {
	var coffee entity.Coffee
	err := r.DB.QueryRow("SELECT id, name, rating, reviews, price_range, type, address, review_text FROM coffees WHERE id = ?", id).
		Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Rating,
			&coffee.Reviews,
			&coffee.PriceRange,
			&coffee.Type,
			&coffee.Address,
			&coffee.ReviewText,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &coffee, nil
}

// Create - Tạo mới cafe
func (r *MySQLCoffeeRepository) Create(coffee *entity.Coffee) (int, error) {
	result, err := r.DB.Exec(
		"INSERT INTO coffees (name, rating, reviews, price_range, type, address, review_text) VALUES (?, ?, ?, ?, ?, ?, ?)",
		coffee.Name,
		coffee.Rating,
		coffee.Reviews,
		coffee.PriceRange,
		coffee.Type,
		coffee.Address,
		coffee.ReviewText,
	)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertedID), nil
}

// Update - Cập nhật cafe theo ID
func (r *MySQLCoffeeRepository) Update(id int, coffee *entity.Coffee) error {
	_, err := r.DB.Exec(
		"UPDATE coffees SET name = ?, rating = ?, reviews = ?, price_range = ?, type = ?, address = ?, review_text = ? WHERE id = ?",
		coffee.Name,
		coffee.Rating,
		coffee.Reviews,
		coffee.PriceRange,
		coffee.Type,
		coffee.Address,
		coffee.ReviewText,
		id,
	)
	return err
}

// Delete - Xóa cafe theo ID
func (r *MySQLCoffeeRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM coffees WHERE id = ?", id)
	return err
}

// Search - Tìm cafe theo từ khóa
func (r *MySQLCoffeeRepository) Search(query string) ([]entity.Coffee, error) {
	rows, err := r.DB.Query(
		"SELECT id, name, rating, reviews, price_range, type, address, review_text FROM coffee WHERE name LIKE ? OR address LIKE ?",
		"%"+query+"%",
		"%"+query+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []entity.Coffee
	for rows.Next() {
		var coffee entity.Coffee
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Rating,
			&coffee.Reviews,
			&coffee.PriceRange,
			&coffee.Type,
			&coffee.Address,
			&coffee.ReviewText,
		)
		if err != nil {
			return nil, err
		}
		coffees = append(coffees, coffee)
	}

	return coffees, nil
}
