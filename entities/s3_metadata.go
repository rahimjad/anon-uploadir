package entities

import (
	"fmt"
	"time"

	"../db"
)

// S3Metadata contains data from the s3_metadata table
type S3Metadata struct {
	ID        string `db:"id"`
	FileName  string `db:"file_name"`
	FileSize  int64  `db:"file_size"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
}

func (record *S3Metadata) QueryRow(id string) error {
	sql := fmt.Sprintf(`
		SELECT id, file_name, file_size, created_at, updated_at
		FROM s3_metadata
		WHERE id = '%s'
		LIMIT 1
	`, id)

	row := db.QueryRow(sql)

	err := row.Scan(
		&record.ID,
		&record.FileName,
		&record.FileSize,
		&record.CreatedAt,
		&record.UpdatedAt,
	)

	return err
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
		INSERT INTO s3_metadata (id, file_name, file_size, created_at, updated_at)
		VALUES ('%s', '%s', %d, %d, %d);`,
		record.ID,
		record.FileName,
		record.FileSize,
		record.CreatedAt,
		record.UpdatedAt,
	)

	fmt.Println(sql)

	_, err := db.Exec(sql)

	return record, err
}

func (record *S3Metadata) Delete() error {
	sql := fmt.Sprintf(`DELETE FROM s3_metadata WHERE id = %s`, record.ID)

	_, err := db.Exec(sql)

	return err
}

func (record *S3Metadata) CountRecords() int {
	var count int

	row := db.QueryRow("SELECT count(id) as count FROM s3_metadata")
	row.Scan(&count)

	return count
}
