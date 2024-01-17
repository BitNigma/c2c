package store

import (
	"database/sql"

	_ "github.com/lib/pq" //..
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New Database conn
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open connect to Database
func (s *Store) Open() error {

	// test local var dbaddr := "host=localhost port=5432 user=pquser dbname=mytest sslmode=disable"
	db, err := sql.Open("postgres", s.config.DataBaseURL) //s.config.DataBaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil

}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
