package main

import (
	"fmt"

	serializer "github.com/mehdy/go-serializer"
)

type Example struct {
	ID       int    `serializer:"id,readOnly"`
	Username string `serializer:"username,omitempty"`
	Password string `serializer:"password,writeOnly"`
}

func main() {
	e := Example{1, "EXAMPLE USERNAME", "EXAMPLE PASSWORD"}
	jsoned, _ := serializer.Marshal(&e)
	fmt.Println(string(jsoned))
}
