package main

import (
	"fmt"

	_ "k8s.io/client-go"
	_ "sigs.k8s.io/yaml"
)

func main() {
	fmt.Println("ok")
}
