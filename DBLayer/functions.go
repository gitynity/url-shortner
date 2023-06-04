package dblayer

import (
	"database/sql"
	"errors"
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

type urls struct{
  original_url string 
  short_code string
  created_at *time.Time
  last_updated_at *time.Time
}

//Notice that all members in urls struct are unexported.
//So they can only be accessed by either methods of urls struct, or in case when the urls object is passed as argument to the function

func InsertURL(db *sql.DB, u *urls) error{
  r,err:=db.Exec("INSERT INTO urls(original_url,short_code,created_at) values(?,?,?)", u.original_url,u.short_code,u.created_at)
  if err != nil {
    return err
  }
  count,err:=r.RowsAffected()
  if count==0{
    return errors.New("no records inserted")
  }
  if err!=nil{
    return err
  }
  return nil
}

func getShortUrl(db *sql.DB, u *urls) (*urls,error){
    query:="select short_code from urls where original_url=?"
    err:=db.QueryRow(query, u.original_url).Scan(u.short_code)
    if err != nil {
      return u, err
    }
    return u,nil
}  
