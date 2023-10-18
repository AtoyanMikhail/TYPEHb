package main

import (
	"typenb/router"
	"typenb/textgen"
)

func main() {
	tg := textgen.NewLocal("./textgen/words.txt")
	r := router.New(tg)

	r.Start(":8000")
}
