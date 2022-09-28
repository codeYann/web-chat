// Package repositories defines a repository to all data layers.
package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/codeYann/web-chat/models"
)

// PostgresRepository defines a connection to postgres database
type PostgresRepository struct {
	db *sql.DB
}

// CreatePostgresRepository returns a new instance of PostgresRepository
func CreatePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// FindAll returns a list of users gets by 'SELECT * FROM users' query.
func (r *PostgresRepository) FindAll() (models.Users, error) {
	var userList models.Users
	usersResponseChannel := make(chan models.User)

	go func() {
		response, err := r.db.Query(`SELECT * FROM users`)
		if err != nil {
			log.Fatal("Unable to run SELECT query", err.Error())
		}
		defer response.Close()
		for response.Next() {
			var user models.User

			err := response.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Password,
				&user.Nickname,
			)
			if err != nil {
				log.Fatal(
					"Error on iterate over rows returned by 'SELECT * FROM users' query.",
					err.Error(),
				)
			}

			usersResponseChannel <- user
		}

		close(usersResponseChannel)
		r.db.Close()
	}()

	for user := range usersResponseChannel {
		userList = append(userList, user)
	}
	return userList, nil
}

// FindOne returns a single user get by "SELECT * FROM users WHERE id = id"
func (r *PostgresRepository) FindOne(ID uint64) (models.User, error) {
	var user models.User
	userReponseChannel := make(chan models.User)

	go func() {
		query := fmt.Sprintf(`SELECT * FROM users WHERE id = %d`, ID)
		stmt, err := r.db.Query(query)
		if err != nil {
			log.Fatal("Unable to prepare SELECT query", err.Error())
		}

		for stmt.Next() {
			err = stmt.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Password,
				&user.Nickname,
			)
			if err != nil {
				log.Fatal("Error on scan row", err.Error())
			}
			userReponseChannel <- user
		}
		close(userReponseChannel)
		r.db.Close()
	}()

	return <-userReponseChannel, nil
}

// SaveOne function creates a new user in the database.
func (r *PostgresRepository) SaveOne(user models.User) (models.User, error) {
	UserSavedChan := make(chan models.User)

	go func() {
		query := fmt.Sprintf(
			`INSERT INTO users 
        (name, email, password, nickname) 
       VALUES
        ('%s', '%s', '%s', '%s')
      `, user.Name, user.Email, user.Password, user.Nickname)

		stmt, err := r.db.Query(query)
		if err != nil {
			log.Fatal("Unable to prepare INSERT INTO query", err.Error())
		}
		for stmt.Next() {
		}
		err = stmt.Close()
		if err != nil {
			log.Fatal("Unable to close statement query", err.Error())
		}
		UserSavedChan <- user
		close(UserSavedChan)
		r.db.Close()
	}()

	return <-UserSavedChan, nil
}

// UpdateOne updates a user nickname and return this user.
func (r *PostgresRepository) UpdateOne(ID uint64, nickname string) (models.User, error) {
	var user models.User

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func(wg *sync.WaitGroup) {
		query := fmt.Sprintf(`UPDATE users SET nickname = '%s' WHERE id = %d`, nickname, ID)

		stmt, err := r.db.Query(query)
		if err != nil {
			log.Fatal("Unable to create UPDATE query", err.Error())
		}

		err = stmt.Close()
		if err != nil {
			log.Fatal("Unable to close UPDATE query", err.Error())
		}

		wg.Done()
		r.db.Close()
	}(&waitGroup)

	waitGroup.Wait()
	return user, nil
}

// DeleteOne deletes a user by id.
func (r *PostgresRepository) DeleteOne(ID uint64) (models.User, error) {
	var user models.User

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func(wg *sync.WaitGroup) {
		query := fmt.Sprintf(`DELETE FROM users WHERE id = %d`, ID)

		stmt, err := r.db.Query(query)
		if err != nil {
			log.Fatal("Unable to prepare DELETE query", err.Error())
		}

		err = stmt.Close()
		if err != nil {
			log.Fatal("Unable to close DELETE query", err.Error())
		}

		wg.Done()
		r.db.Close()
	}(&waitGroup)

	waitGroup.Wait()
	return user, nil
}
