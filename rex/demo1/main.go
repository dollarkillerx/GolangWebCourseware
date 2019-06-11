/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 上午9:47
* */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func handleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	url := "http://www.jihaoba.com/escrow/"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	handleError(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	html := fmt.Sprintf("%s",bytes)

	//rexPhone := `1[3456789]\d{9}`
	rexPhone := `(1[3456789]{2})(\d{4})(\d{4})`
	rex := regexp.MustCompile(rexPhone)
	//allString := rex.FindAllString(html, -1) // -1 匹配全部
	//fmt.Printf("%s",allString)
	submatch := rex.FindAllStringSubmatch(html, -1) // -1 匹配全部 这个数字是匹配多少个
	fmt.Printf("%s",submatch) // FindAllStringSubmatch 可以获得分组的子匹配
}

func test()  {
	url := "http://www.jihaoba.com/escrow/"
	resp, err := http.Get(url)
	handleError(err)

	file, err := os.OpenFile("./he.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 00666)
	handleError(err)
	reader := bufio.NewReader(resp.Body)
	writer := bufio.NewWriter(file)

	bytes := make([]byte, 1024)
	for {
		_, err := reader.Read(bytes)
		if err == io.EOF {
			break
		}else{
			writer.Write(bytes)
		}
	}
	defer func() {
		writer.Flush()
		resp.Body.Close()
		file.Close()
	}()
}