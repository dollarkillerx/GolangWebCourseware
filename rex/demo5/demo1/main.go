/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 上午11:13
* */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func handErr(err error)  {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	html := make(chan string, 10)
	img := make(chan string, 10)
	go get(html)
	go reg(html,img)
	go dow(img)
	time.Sleep(time.Second * 10)
}

func get(html chan string)  {
	resp, err := http.Get("https://www.hao123.com/")
	handErr(err)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	handErr(err)
	html <- fmt.Sprintf("%s",bytes)
}

func reg(html,img chan string)  {
	reg := `<img[\s\S]+?src="(http[\s\S]+?)"`
	compile := regexp.MustCompile(reg)
	for {
		select {
		case data := <-html :{
			submatch := compile.FindAllStringSubmatch(data, -1)
			for _,v:=range submatch  {
				img <- v[1]
			}
		}
		}
	}
}

func dow(img chan string) {
	for {
		select {
		case src := <-img :{
			download(src)
		}
		}
	}
}

func download(src string)  {
	resp, err := http.Get(src)
	if err != nil{
		return
	}
	split := strings.Split(src, ".")
	ec := split[len(split)-1]
	s, _ := NewUUIDSimplicity()
	newpath := "./images/" + s + "." + ec
	file, err := os.OpenFile(newpath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 00666)
	if err != nil {
		return
	}
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

func NewUUID() (string,error) {
	out, err := exec.Command("uuidgen").Output()

	oot := fmt.Sprintf("%s", out)
	return oot,err
}


func NewUUIDSimplicity() (string,error) {
	s, e := NewUUID()
	var u string
	for _,k :=range s {
		if k != '-' {
			u = fmt.Sprintf("%s%s",u,string(k))
		}
	}
	return u,e
}