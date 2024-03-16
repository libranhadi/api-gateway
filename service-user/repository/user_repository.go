package repository

import (
	"database/sql"
	"errors"
	"service-user/config"
	"service-user/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (userRepo *userRepositoryImpl) Create(user *model.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1,$2)"
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	_, errExec := userRepo.db.ExecContext(ctx, query, &user.Email, &user.Password)
	if errExec != nil {
		return errors.New("error, creating user")
		// return fmt.Errorf("error, create user %w", errExec)
	}
	return nil
}

func (userRepo *userRepositoryImpl) FindUserByEmail(email string) (*model.User, error) {
	query := "SELECT id, password ,email FROM users WHERE email = $1"
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	user := &model.User{}

	rows, err := userRepo.db.QueryContext(ctx, query, email)
	if err != nil {
		return user, errors.New("error checking email existence")
		// return user, fmt.Errorf("error checking email existence: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Password, &user.Email)
		if err != nil {
			return user, errors.New("error scanning email existence result")
			// return user, fmt.Errorf("error scanning email existence result: %w", err)
		}

		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}
