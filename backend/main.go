package main

import (
	"github.com/zozoee27/cookbook/backend/app"
)

func main() {
	a := app.App{}
	a.Initialize("CookbookDB")

	a.Run(":8080")
}
