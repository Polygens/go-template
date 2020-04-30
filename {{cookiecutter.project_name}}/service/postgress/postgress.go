package postgress

import (
	"github.com/go-pg/pg/v9"
)

func Ping() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()
}
