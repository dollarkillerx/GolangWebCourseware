/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 上午11:04
* */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func handleErr(err error)  {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	html := get()
	reg(html)
}

func reg(html string)  {
	rexSFZ := `\d{18}`
	compile := regexp.MustCompile(rexSFZ)
	submatch := compile.FindAllStringSubmatch(html, -1)
	fmt.Printf("%s",submatch)

}

func get() string  {
	resp, err := http.Get("https://hb.qq.com/a/20180606/005098.htm")
	handleErr(err)
	defer func() {
		resp.Body.Close()
	}()
	bytes, err := ioutil.ReadAll(resp.Body)
	handleErr(err)
	return fmt.Sprintf("%s",bytes)
}
