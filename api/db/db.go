package db

import (
	"database/sql"
	"log"
	"os/exec"
	"strings"

	"../config"

	_ "github.com/lib/pq"
)

// connect establishes a connection to your DB.
func connect() *sql.DB {
	dbConfig := config.New().DB
	connectionString := dbConfig.BuildConnectionString()

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic(err)
	}

	return db
}

// Exec will execute a given SQL string on the db and return sql.Result
func Exec(sql string) (sql.Result, error) {
	db := connect()
	defer db.Close()

	return db.Exec(sql)
}

// QueryRow will query for and return a row of data
func QueryRow(sql string) *sql.Row {
	db := connect()
	defer db.Close()

	return db.QueryRow(sql)
}

// Query will query for and return multiple rows of data
func Query(sql string) (*sql.Rows, error) {
	db := connect()
	defer db.Close()

	return db.Query(sql)
}

// Migrate is a hard coded migration function to create the s3 table
// were this app to require more migrations & control, I would use something more robust
// including version tracking for example golang-migrate. Additionally, I'd also add some indexes
// to improve query performance. I'm also using unix time for now, ideally I would use timestamps
// however for this itteration I'm going as simple as possible.
func Migrate() (sql.Result, error) {
	return Exec(`
		CREATE TABLE IF NOT EXISTS s3_metadata (
			id         uuid,
			file_name  varchar(255),
			file_ext   varchar(255),
			location   text,
			file_size  bigint,
			created_at bigint,
			updated_at bigint
		)
	`)
}

// Rollback is for quick testing, drops the s3_metadata table
func Rollback() (sql.Result, error) {
	return Exec(`DROP TABLE IF EXISTS s3_metadata`)
}

// GenerateUUID is used to generate the ID value
func GenerateUUID() string {
	uuid, _ := exec.Command("uuidgen").Output()
	lowerUUID := strings.ToLower(string(uuid))
	return strings.TrimSpace(lowerUUID)
}
