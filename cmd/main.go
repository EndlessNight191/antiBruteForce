package main

import (
	app "test/internal"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	app.Run()
}