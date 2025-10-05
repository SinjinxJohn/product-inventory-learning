package user

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

func (s *Store) GetUserByID(ID int) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, firstName, lastName, email, password,role, createdAt FROM users WHERE ID=?", ID)

	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsInUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil

}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, firstName, lastName, email, password,role, createdAt FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanRowsInUser(rows)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("Scanned User: %+v\n", user)
		return user, nil
	}
	return nil, fmt.Errorf("user not found")

}
func (s *Store) CreateUser(user *types.RegisterUser) error {

	//create a new types.RegisterUser
	query := `INSERT INTO users(firstName,lastName,email,password,role)
	VALUES(?,?,?,?,?)
	`
	_, err := s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role)

	if err != nil {
		return err
	}
	return nil

	//return user
}

func scanRowsInUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
