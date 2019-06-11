/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 上午10:39
* */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func headErr(err error)  {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	html := get()
	reg(html)
}

func get() string {
	resp, err := http.Get("https://www.hao123.com/")
	headErr(err)
	defer func() {
		resp.Body.Close()
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	headErr(err)
	return fmt.Sprintf("%s",bytes)
}

func reg(html string)  {
	//fmt.Println(html)
	regUrl := `(http|https)://([\w]*).([\w]+).(com|net|com.cn)`
	regA := `<a[\s\S]*?href="(http[\s\S]+?)"`

	compile := regexp.MustCompile(regUrl)
	submatch := compile.FindAllStringSubmatch(html, -1)
	fmt.Printf("%s",submatch)

	mustCompile := regexp.MustCompile(regA)
	fmt.Println()

	stringSubmatch := mustCompile.FindAllStringSubmatch(html, -1)
	for _,i:=range stringSubmatch {
		fmt.Println(i[1])
	}
}