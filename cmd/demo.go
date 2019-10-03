package main

import (
	"fmt"
	"github.com/serbi/secret"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")

	err := v.Set("demo_key_1", "some random value 1")
	if err != nil {
		panic(err)
	}

	err = v.Set("demo_key_2", "some random value 2")
	if err != nil {
		panic(err)
	}

	err = v.Set("demo_key_3", "some random value 3")
	if err != nil {
		panic(err)
	}

	plain, err := v.Get("demo_key_1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)

	plain, err = v.Get("demo_key_2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)

	plain, err = v.Get("demo_key_3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
