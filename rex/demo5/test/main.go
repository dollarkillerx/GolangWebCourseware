/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 上午11:38
* */
package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "hepps.png"
	split := strings.Split(path, ".")
	fmt.Println(split[len(split)-1])
}
