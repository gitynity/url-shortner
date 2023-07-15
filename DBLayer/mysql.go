package dblayer

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBconfig() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/urlShortnerDB?parseTime=true")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*
why would you import go-sql-driver/mysql if you are not going to use it?

The import statement `_ "go-sql-driver/mysql"` is used to import the MySQL driver package in Go, which is required to connect and interact with a MySQL database using the `database/sql` package.

Even though it may seem like the imported package is not directly used in the code snippet you provided, it is still necessary for the proper functioning of the `database/sql` package with MySQL. By importing the package, the MySQL driver is registered with the `database/sql` package, allowing you to use the `sql.Open` function with the `"mysql"` driver name.
go get mvdan.cc/gofumpt

The `_` before the import statement is used to indicate that you only want to import the package for its side effects, such as registering the driver, but not directly reference any of its functions or variables in your code.

Therefore, even if you don't explicitly use the imported package in your code snippet, it is essential for establishing a connection to a MySQL database using the `database/sql` package.

Lets see it in more depth..

The `database/sql` package in Go provides a generic interface for working with databases. It defines a set of common methods and types that can be used to interact with different database systems. However, in order to connect and communicate with a specific database system, such as MySQL, you need a driver that implements the interface defined by `database/sql`.

The `_ "go-sql-driver/mysql"` import statement is used to import the MySQL driver package, which is an implementation of the `database/sql` interface for MySQL databases. This driver package is responsible for handling the low-level communication with the MySQL database server, such as establishing connections, executing queries, and retrieving results.

By importing the `go-sql-driver/mysql` package and using the `sql.Open` function with the `"mysql"` driver name, you are instructing the `database/sql` package to use the MySQL driver for establishing a connection to a MySQL database.

Without importing and registering the MySQL driver package, the `database/sql` package wouldn't be aware of the specific MySQL driver implementation, and it wouldn't know how to interact with MySQL databases. Importing the driver package ensures that the necessary code is executed to set up the communication between `database/sql` and MySQL.

In summary, importing the `go-sql-driver/mysql` package is necessary for the proper functioning of the `database/sql` package with MySQL because it provides the implementation of the `database/sql` interface for MySQL databases, enabling you to connect, query, and interact with a MySQL database using the `database/sql` package's methods and types.

*/
