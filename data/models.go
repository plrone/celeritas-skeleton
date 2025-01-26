package data

import (
	"database/sql"
	"fmt"
	"os"

	up "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper up.Session

type Models struct {
	// any models inserted here are easily
	// accessible throught the entire applicationsa

	// Users  User
	// Tokens Token
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		upper, _ = postgresql.New(databasePool)
	}

	return Models{
		// Users:  User{},
		// Tokens: Token{},
	}
}

func getInsertID(i up.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
