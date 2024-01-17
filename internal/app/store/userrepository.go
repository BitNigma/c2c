package store

import (
	"website/internal/app/model"
)

//User repos
type UserRepository struct {
	store *Store
}

//Create user model
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT into users (email, encrypted_password)VALUES($1, $2)RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil

}

//Find users in Database
func (r *UserRepository) FindbyEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, ecrypted_password from users WHERE email=$1",
		email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {

		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	results := []model.User{}
	rows, err := r.store.db.Query("SELECT * from users")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()

	u := &model.User{}
	for rows.Next() {
		rows.Scan(&u.ID, &u.Email, &u.Password)
		results = append(results, *u)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return results, nil
}
