package main

import (
	DB "dbComparisonGo/database"
)

func main() {
	DB.ConnectDB("dev")
}
