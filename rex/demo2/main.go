/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 上午10:28
* */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	html := get()
	rex(html)
}

func get() string {
	resp, err := http.Get("https://zhidao.baidu.com/question/1772307510687057620.html")
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		resp.Body.Close()
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s",bytes)
}

func rex(html string) {
	rexEmail := `(\w+)@(\w+).(com|net)`
	compile := regexp.MustCompile(rexEmail)
	submatch := compile.FindAllStringSubmatch(html, -1)
	fmt.Print("%s\n",submatch)
}