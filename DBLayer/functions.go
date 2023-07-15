package dblayer

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

/*
Schema:
CREATE TABLE urls (
  id INT AUTO_INCREMENT PRIMARY KEY,
  original_url VARCHAR(255) NOT NULL,
  short_code VARCHAR(10) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

*/

type URL struct {
	Original_url string
	Short_code   string
	Created_at   time.Time
}

//Notice that all members in urls struct are unexported.
//So they can only be accessed by either methods of urls struct, or in case when the urls object is passed as argument to the function

func InsertURL(db *sql.DB, u *URL) error {
	r, err := db.Exec("INSERT INTO urls(original_url,short_code,created_at) values(?,?,?)", u.Original_url, u.Short_code, u.Created_at)
	if err != nil {
		return err
	}
	count, err := r.RowsAffected()
	if count == 0 {
		return errors.New("no records inserted")
	}
	if err != nil {
		return err
	}
	return nil
}

func GetShortUrl(db *sql.DB, u *URL) (*URL, error) {
	query := "select short_code from urls where original_url=?"
	err := db.QueryRow(query, u.Original_url).Scan(&u.Short_code)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, nil
}

func GetLongUrl(db *sql.DB, u *URL) (*URL, error) {
	query := "select original_url from urls where short_code=?"
	err := db.QueryRow(query, u.Short_code).Scan(&u.Original_url)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, nil
}
func CheckUrlExists(db *sql.DB, u *URL) (*URL, error) {
	query := "select original_url from urls where original_url=?"
	err := db.QueryRow(query, u.Original_url).Scan(&u.Original_url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		log.Println("i am here", err)
		return u, err
	}
	return u, nil
}

func DeleteURL(db *sql.DB, u *URL) error {
	r, err := db.Exec("delete from urls where original_url=?", u.Original_url)
	if err != nil {
		return err
	}
	count, err := r.RowsAffected()
	if count == 0 {
		return errors.New("no records deleted")
	}
	if err != nil {
		return err
	}
	return nil
}
