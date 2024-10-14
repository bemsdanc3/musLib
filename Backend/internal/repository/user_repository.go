package repository

import (
	"Backend/internal/entities"
	"database/sql"
	"errors"
)

type UserRepository interface {
	GetAllUsers() ([]*entities.User, error)
	CreateUSer(user *entities.User) error
	GetUserByID(id int) (user *entities.User, err error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]*entities.User, error) {
	query := `SELECT id, login, email from users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(&user.ID, &user.Login, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) CreateUSer(user *entities.User) error {
	query := `INSERT INTO users (login, email, pass) VALUES (?,?,?)`
	_, err := r.db.Exec(query, user.Login, user.Email, user.Pass)
	return err
}

func (r *userRepository) GetUserByID(id int) (*entities.User, error) {
	query := `SELECT login, email FROM users WHERE id = ?`
	user := &entities.User{}

	err := r.db.QueryRow(query, id).Scan(&user.Login, &user.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
