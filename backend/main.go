package main

import (
	"github.com/zozoee27/cookbook/backend/application"
)

func main() {
	a := application.App{}
	a.Initialize("CookbookDB")

	a.Run(":8080")
}
