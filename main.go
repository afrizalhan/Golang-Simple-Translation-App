package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	fmt.Println("=========ENGLISH TO INDONESIAN TRANSLATOR==========")
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter English Sentence : ")
	sentence, _ := reader.ReadString('\n')


	origin, trans := translate(sentence)

	fmt.Println("=========RESULT==========")
	fmt.Println("Origin Word :", origin)
	fmt.Println("Translated Word :", trans)
}

func translate(sentence string)(interface{}, interface{}){
	endpoint := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&ie=UTF-8&oe=UTF-8&inputm=2&otf=2&iid=1dd3b944-fa62-4b55-b330-74909a99969e"

	data := url.Values{}
    data.Set("sl", "en")
    data.Set("tl", "id")
    data.Set("q", sentence)

    client := &http.Client{}
    r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
    if err != nil {
        log.Fatal(err)
    }
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    res, err := client.Do(r)
    if err != nil {
        log.Fatal(err)
    }
    // fmt.Println(res.Status)
	if res.Status != "200 OK"{
		log.Fatal("failed to fetch result")
	}
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
	var data1 map[string][]interface{}
	json.Unmarshal([]byte(body), &data1)
	data2 := data1["sentences"][0].(map[string]interface{})
	origin := data2["orig"]
	trans := data2["trans"]
	return origin, trans
}