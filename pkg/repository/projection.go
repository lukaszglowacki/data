package repository

import (
	"database/sql"
	"time"
)

type Projection struct {
	db *sql.DB
}

func NewProjection(db *sql.DB) Projection {
	return Projection{db}
}

func (p Projection) Save(t time.Time, json string) error {
	stmt, err := p.db.Prepare(`INSERT INTO jobs (time, data) VALUES (?,?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t, json)
	if err != nil {
		return err
	}

	return nil
}

func (p Projection) Get() ([]string, error) {
	rows, err := p.db.Query("SELECT Data FROM JOBS")
	if err != nil {
		return nil, err
	}

	jobs := []string{}
	for rows.Next() {
		var json string
		err = rows.Scan(&json)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, json)
	}

	rows.Close()
	return jobs, nil
}
