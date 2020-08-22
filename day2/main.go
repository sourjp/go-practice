package main

import (
	_ "github.com/lib/pq"
	"github.com/sourjp/go-practice/day2/controllers"
)

func main() {
	controllers.Router()
}
