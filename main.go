package main

import (
	"P1/db"
	"P1/router"
)

func main() {
	db.InitDB()
	router.SetRouter()
}
