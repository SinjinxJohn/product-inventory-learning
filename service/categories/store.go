package categories

import (
	"database/sql"
	"fmt"

	"example.com/event-app/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllCategories() ([]*types.Categories, error) {
	rows, err := s.db.Query("SELECT ID,name,createdAt FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*types.Categories

	for rows.Next() {
		category, err := scanRowsInCategories(rows)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *Store) CreateCategory(cat *types.CreateCategoryPayload) error {
	query := `INSERT INTO category(name) VALUES(?)`
	_, err := s.db.Exec(query, cat.Name)

	if err != nil {
		return err
	}
	return nil

}

func (s *Store) GetCategoryByName(categoryName string) (*types.Categories, error) {
	rows, err := s.db.Query("SELECT (id,name,createdAt) FROM category WHERE name = ?")
	if err != nil {
		return nil, err
	}
	cat := new(types.Categories)
	for rows.Next() {
		cat, err = scanRowsInCategories(rows)
		if err != nil {
			return nil, err
		}
	}
	if cat.ID == 0 {
		return nil, fmt.Errorf("category not found")
	}
	return cat, nil

}

func scanRowsInCategories(rows *sql.Rows) (*types.Categories, error) {
	category := new(types.Categories)
	err := rows.Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return category, nil
}
