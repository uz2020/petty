package main

import (
	"fmt"
	"github.com/uz2020/petty/xq/cmd"
)

func main() {
	fmt.Println("good")
	print(1)
	print("haha")
	cmd.Execute()
	print("bug fix")
}
