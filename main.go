package main

import (
	"fmt"

	_ "k8s.io/client-go"
	_ "sigs.k8s.io/yaml"

	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/rest"
	_ "k8s.io/client-go/tools/clientcmd"

	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	fmt.Println("ok")
}
