package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_sql/config"
	"go_sql/models"
	"log"
	"time"
)

const (
	table          = "movie"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie
	db, err := config.MySql()

	// config db check
	if err != nil {
		log.Fatal("can't connect to sql", err)
	}

	// query
	queryText := fmt.Sprintf("SELECT * FROM %v Order by CREATED_AT DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var movie models.Movie
		var createdAt, updatedAt string
		if err := rowQuery.Scan(&movie.ID,
			&movie.Title,
			&movie.Year,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		movie.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		movie.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

func PostData(ctx context.Context, movie models.Movie) error {
	db, err := config.MySql()

	// config db check
	if err != nil {
		log.Fatal("can't connect to sql", err)
	}

	// query
	queryText := fmt.Sprintf("INSERT INTO %v (title, year, created_at, updated_at) value ('%v', '%v', NOW(), NOW())", table, movie.Title, movie.Year)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}
	return nil
}

func PutData(ctx context.Context, movie models.Movie, idMovie string) error {
	db, err := config.MySql()

	// config db check
	if err != nil {
		log.Fatal("can't connect to sql", err)
	}

	// query
	queryText := fmt.Sprintf("UPDATE %v SET title='%v', year='%v', updated_at= NOW() where id=%v", table, movie.Title, movie.Year, idMovie)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}
	return nil
}

func DelData(ctx context.Context, idMovie string) error {
	db, err := config.MySql()

	// config db check
	if err != nil {
		log.Fatal("can't connect to sql", err)
	}

	// query
	queryText := fmt.Sprintf("DELETE FROM %v where id=%v", table, idMovie)
	s, err := db.ExecContext(ctx, queryText)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	if check == 0 {
		return errors.New("id tidak ditemukan")
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
