package infraestructure

import (
	"database/sql"
	"fmt"

	"github.com/Marcxz/academy-go-q32021/conf"

	_ "github.com/lib/pq"
)

type GeoDB interface {
}

type geoDB struct {
	db  *sql.DB
	con *conf.Config
}

func NewGeoDB(con *conf.Config) GeoDB {
	return geoDB{
		db:  nil,
		con: con,
	}
}

func (gdb *geoDB) InitDB() error {
	var err error
	con := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", gdb.con.Postgres_user, gdb.con.Postgres_password, gdb.con.Postgres_host, gdb.con.Postgres_port, gdb.con.Postgres_db)
	gdb.db, err = sql.Open("postgres", con)

	if err != nil {
		return err
	}

	err = gdb.db.Ping()

	if err != nil {
		return err
	}
	return nil
}
