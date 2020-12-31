package models

import (
	"os"
	"github.com/Gin58/gin-api/db"
)

var server = os.Getenv("DATABASE")
var databaseName = os.Getenv("DATABASE_NAME")
var dbConnect = db.NewConnection(server, databaseName)
