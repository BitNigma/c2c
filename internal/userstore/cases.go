package userstore

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"p2p/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorByEmail  = errors.New("user doesn't exist")
	ErrorPassword = errors.New("wrong user password")
	UserCheckByID = errors.New("can't find user in Database by ID")
)

func (u *Store) CreateUser(c *fiber.Ctx, user model.User) error {

	ctx := context.Background()
	if err := u.PG.Ping(ctx); err != nil {
		log.Println("Connection is closed because store is empty", err)
	}

	// Set initialized default data for book:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UserStatus = 1 // 0 == draft, 1 == active

	if err := u.GetUserByEmail(user); err != nil {
		{
			switch {
			case errors.Is(err, ErrorByEmail):
				log.Println(ErrorByEmail)
				return c.Render("./static/signup.html", fiber.Map{
					"Islogin": false,
				})
			}
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}

	query := `INSERT INTO users(id, created_at, email, encrypted_password, user_status) VALUES(@Id, @Data, @Email, @Encrypted_password, @Status)`
	args := pgx.NamedArgs{
		"Id":                 user.ID,
		"Data":               user.CreatedAt,
		"Email":              user.Email,
		"Encrypted_password": string(hashedPassword),
		"Status":             user.UserStatus,
	}

	_, err = u.PG.db.Exec(ctx, query, args)
	if err != nil {
		log.Println("can't make a row in database", err)
		return err
	}

	return nil
}

func (u *Store) GetUserByEmail(user model.User) error {

	// Query to request
	query := `SELECT email FROM users WHERE email=$1`
	err := u.PG.db.QueryRow(context.Background(), query, user.Email).Scan(&user.Email)

	if err != nil {
		return err
	}

	if err == sql.ErrNoRows {
		return ErrorByEmail
	}
	return nil
}

func (u *Store) CheckAuthUser(user model.User) error {

	var passwd string
	query := `SELECT encrypted_password FROM users WHERE email=$1`

	err := u.PG.db.QueryRow(context.Background(), query, user.Email).Scan(&passwd)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return ErrorByEmail
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(user.Password))
	if err != nil {
		return ErrorPassword
	}
	return nil
}

func (u *Store) GetUserById(id string) error {

	user := model.NewUser()
	query := `SELECT id FROM users WHERE id=$1`

	err := u.PG.db.QueryRow(context.Background(), query, id).Scan(&user.ID)
	if err != nil {
		log.Println("can't made a request to find User by id", err)
		return err
	}

	return nil
}
