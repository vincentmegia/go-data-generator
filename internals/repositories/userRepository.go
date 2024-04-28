package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/vincentmegia/go-data-generator/internals/common"
	"github.com/vincentmegia/go-data-generator/internals/models"
)

type UserRepository struct{}

func (ur *UserRepository) AddUser(user *models.User) error {
	database, error := openDatabase()
	if error != nil {
		log.Fatalln("FAILED: opening database connection")
		return error
	}
	statement, error := database.Prepare(`INSERT INTO public.users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`)
	if error != nil {
		log.Fatalln("FAILED: preparing add user statement:\t", error)
		return error
	}
	result, error := statement.Exec(user.Firstname, user.Lastname,
		user.Username, user.Password, user.Salt, user.Address_id,
		user.Created_date, user.Updated_date, user.Email, user.Birth_date)
	statement.Close()
	if error != nil {
		log.Fatalln("FAILED: statement execution:\t", error)
		return error
	}
	database.Close()
	lastId, _ := result.LastInsertId()
	totalRowsAffected, _ := result.RowsAffected()
	fmt.Println("SUCCESS: last insert id:\t", lastId)
	fmt.Println("SUCCESS: Rows affected:\t", totalRowsAffected)
	return nil
}

func (ur *UserRepository) BulkInsert(users *[]models.User) error {
	/* 	fmt.Println("-----------array batches: ", users) */
	database, error := openDatabase()
	if error != nil {
		log.Fatalln("FAILED: opening database connection")
		return error
	}
	transaction, error := database.Begin()
	statement, error := transaction.Prepare(pq.CopyIn("users", "firstname", "lastname", "username", "password", "salt", "address_id", "created_date", "updated_date", "email", "birth_date"))
	for _, user := range *users {
		_, error := statement.Exec(user.Firstname, user.Lastname, user.Username, user.Password, user.Salt, nil, user.Created_date, nil, user.Email, user.Birth_date)
		if error != nil {
			log.Fatalln("FAILED: failed executing bulk insert for user:\t", user)
		}
	}

	if _, error := statement.Exec(); error != nil {
		log.Fatalln("FAILED: error happened during bulk insert", error)
		return error
	}
	if error := statement.Close(); error != nil {
		log.Fatalln("FAILED: error has occured while closing statement", error)
		return error
	}
	if error := transaction.Commit(); error != nil {
		log.Fatalln("FAILED: commit error, rolling back transaction")
		transaction.Rollback()
		return error
	}
	return nil
}

func openDatabase() (*sql.DB, error) {
	//TODO: transfer out to env variables
	dbConfig := common.ConfigurationManager{}.GetDBConfig()
	fmt.Println("Opening database server: ", dbConfig.Hostname)
	connectionString := "postgresql://postgres:postgres@192.168.3.110/crm?sslmode=disable"
	database, error := sql.Open("postgres", connectionString)
	if error != nil {
		log.Fatalln("FAILED: connecting to db failed:\t", error)
		return nil, error
	}
	return database, nil
}
