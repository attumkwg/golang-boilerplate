package main

import (
	"github.com/attumkwg/mplatmm/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	defer db.Connection.Close()
	r := infrastructure.NewRouting(db)
	r.Run()
}
