package main

import (
	"github.com/qoxogus/GSM-Festival-Master-Back/database"
	"github.com/qoxogus/GSM-Festival-Master-Back/rest"
)

func main() {

	rest.RunAPI(":1324")
	database.Connect()
}
