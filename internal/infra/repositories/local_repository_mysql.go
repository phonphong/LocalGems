package repositories

import (
	"database/sql"
	"local-gems-server/internal/core/entity"
)

type MySQLLocalRepository struct {
	DB *sql.DB
}

func NewMySQLLocalRepository(db *sql.DB) LocalRepository {
	return &MySQLLocalRepository{DB: db}
}

func (r *MySQLLocalRepository) FindAll() ([]entity.Local, error) {
	rows, err := r.DB.Query("SELECT id, name, rating, reviews, price_range, type, address, review_text FROM locals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locals []entity.Local
	for rows.Next() {
		var local entity.Local
		err := rows.Scan(
			&local.ID,
			&local.Name,
			&local.Rating,
			&local.Reviews,
			&local.PriceRange,
			&local.Type,
			&local.Address,
			&local.ReviewText,
		)
		if err != nil {
			return nil, err
		}
		locals = append(locals, local)
	}

	return locals, nil
}

// FindByID
func (r *MySQLLocalRepository) FindByID(id int) (*entity.Local, error) {
	var local entity.Local
	err := r.DB.QueryRow("SELECT id, name, rating, reviews, price_range, type, address, review_text FROM locals WHERE id = ?", id).
		Scan(
			&local.ID,
			&local.Name,
			&local.Rating,
			&local.Reviews,
			&local.PriceRange,
			&local.Type,
			&local.Address,
			&local.ReviewText,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &local, nil
}

// Create
func (r *MySQLLocalRepository) Create(local *entity.Local) (int, error) {
	result, err := r.DB.Exec(
		"INSERT INTO locals (name, rating, reviews, price_range, type, address, review_text) VALUES (?, ?, ?, ?, ?, ?, ?)",
		local.Name,
		local.Rating,
		local.Reviews,
		local.PriceRange,
		local.Type,
		local.Address,
		local.ReviewText,
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

// Update
func (r *MySQLLocalRepository) Update(id int, local *entity.Local) error {
	_, err := r.DB.Exec(
		"UPDATE locals SET name = ?, rating = ?, reviews = ?, price_range = ?, type = ?, address = ?, review_text = ? WHERE id = ?",
		local.Name,
		local.Rating,
		local.Reviews,
		local.PriceRange,
		local.Type,
		local.Address,
		local.ReviewText,
		id,
	)
	return err
}

// Delete
func (r *MySQLLocalRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM locals WHERE id = ?", id)
	return err
}

// Search
func (r *MySQLLocalRepository) Search(query string) ([]entity.Local, error) {
	rows, err := r.DB.Query(
		"SELECT id, name, rating, reviews, price_range, type, address, review_text FROM locals WHERE name LIKE ? OR address LIKE ?",
		"%"+query+"%",
		"%"+query+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locals []entity.Local
	for rows.Next() {
		var local entity.Local
		err := rows.Scan(
			&local.ID,
			&local.Name,
			&local.Rating,
			&local.Reviews,
			&local.PriceRange,
			&local.Type,
			&local.Address,
			&local.ReviewText,
		)
		if err != nil {
			return nil, err
		}
		locals = append(locals, local)
	}

	return locals, nil
}
