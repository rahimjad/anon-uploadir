package entities

import (
	"fmt"
	"log"
	"time"

	"../db"
)

// S3Metadata contains data from the s3_metadata table
type S3Metadata struct {
	ID            string `db:"id"`
	FileName      string `db:"file_name"`
	FileExtension string `db:"file_ext"`
	Location      string `db:"location"`
	FileSize      int    `db:"file_size"`
	CreatedAt     int    `db:"created_at"`
	UpdatedAt     int    `db:"updated_at"`
}

func (record *S3Metadata) QueryRow(id string) {
	sql := fmt.Sprintf(`
		SELECT *
		FROM s3_metadata
		WHERE id = '%s'
		LIMIT 1
	`, id)

	row := db.QueryRow(sql)

	err := row.Scan(
		&record.ID,
		&record.FileName,
		&record.FileExtension,
		&record.Location,
		&record.FileSize,
		&record.CreatedAt,
		&record.UpdatedAt,
	)

	if err != nil {
		log.Panic(err)
	}

	if record.ID == "" {
		panic("Record does not exist")
	}
}

// Insert will create an s3_metadata record in the DB
func (record *S3Metadata) Insert() (*S3Metadata, error) {
	currentTime := int(time.Now().Unix())
	record.UpdatedAt = currentTime

	if record.ID == "" {
		record.ID = db.GenerateUUID()
		record.CreatedAt = currentTime
	}

	sql := fmt.Sprintf(`
		INSERT INTO s3_metadata (id, file_name, file_ext, file_size, location, created_at, updated_at)
		VALUES ('%s', '%s', '%s', %d, '%s', %d, %d);`,
		record.ID,
		record.FileName,
		record.FileExtension,
		record.FileSize,
		record.Location,
		record.CreatedAt,
		record.UpdatedAt,
	)

	_, err := db.Exec(sql)

	return record, err
}

func (record *S3Metadata) Delete() bool {
	sql := fmt.Sprintf(`DELETE FROM s3_metadata WHERE id = %d`, &record.ID)

	_, err := db.Exec(sql)

	return err == nil
}

func (record *S3Metadata) CountRecords() int {
	var count int

	row := db.QueryRow("SELECT count(id) as count FROM s3_metadata")
	row.Scan(&count)

	return count
}
