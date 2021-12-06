package db

import (
	"fmt"
	"log"
	"time"

	"ycoem/lib/logs"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var Sqlx *sqlx.DB
var SqlBuilder = squirrel.StatementBuilder

type Eq = squirrel.Eq
type Or = squirrel.Or
type Like = squirrel.Like
type SelectBuilder = squirrel.SelectBuilder
type InsertBuilder = squirrel.InsertBuilder
type UpdateBuilder = squirrel.UpdateBuilder

// NewMySQL 创建一个连接MySQL的实体池
func NewMySQL(dbSource string, maxOpenConns, maxIdleConns int) (Sqlx *sqlx.DB) {
	Sqlx, err := sqlx.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	if maxOpenConns > 0 {
		Sqlx.SetMaxOpenConns(maxOpenConns)
	}

	if maxIdleConns > 0 {
		Sqlx.SetMaxIdleConns(maxIdleConns)
	}

	go func() {
		for {
			err := Sqlx.Ping()
			if err != nil {

				if logs.Logger == nil {
					log.Println("mysql db can't connect!")
				} else {
					logs.Logger.Error("mysql db can't connect!")
				}

			}
			time.Sleep(time.Minute)
		}
	}()

	return
}

func CreatePage(page, limit int) uint64 {
	tmp := (page - 1) * limit
	return uint64(tmp)
}

func CreateLike(column string) string {
	return fmt.Sprint("%", column, "%")
}
