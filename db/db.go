package db

import (
	"database/sql"

	"githib/com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config)