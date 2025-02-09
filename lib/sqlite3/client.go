package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3Client interface {
	Close() error
	Exec(stmt string)
	Read(stmt string, dest ...any) error
}

type sqlite3Client struct {
	c   *sql.DB
	ctx context.Context
}

func (sq sqlite3Client) Close() error {
	return sq.c.Close()
}

func (sq sqlite3Client) Exec(stmt string) {
	_, err := sq.c.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (sq sqlite3Client) Read(stmt string, dest ...any) error {
	rows, err := sq.c.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return err
		}
		fmt.Println(dest...)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func NewClient(ctx context.Context) Sqlite3Client {
	sc := &sqlite3Client{
		ctx: ctx,
	}
	db, err := sql.Open("sqlite3", "./db/sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
	sc.c = db
	return sc
}
