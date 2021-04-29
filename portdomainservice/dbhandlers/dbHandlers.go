package dbhandlers

import (
	"log"
	srv "portdomainservice/service"
)

const (
	DBHandler_MEM = 1
)

type dbType int

func NewDBHandler(db dbType, dbConfig *DBConfig) srv.DBHandlerInterf {
	switch db {
	case DBHandler_MEM:
		return NewMemDBHandler(dbConfig)
	default:
		log.Printf("DB type undefined")
		return nil
	}
}
