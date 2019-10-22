package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var mtUrl = "http://122.144.179.43:7862/sms?"
var mtUsername = "390064"
var mtPassword = "DZKkmX"

func main()  {
	sendSMS()
	select {

	}
}

func sendSMS() {
	client := &http.Client{}


	url := mtUrl + "action=send&"
	url += "account=" + mtUsername
	url += "&password=" + mtPassword
	url += "&mobile="+ "18611785986"
	url += "&content="+ "hello word 你好"
	url += "&extno=1069088964"
	url += "&rt=json"

	req, err := http.NewRequest("POST",
		url, nil)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "*/*")
	req.Header.Set("connection", "Keep-Alive")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(body)
	fmt.Println("jsonStr", jsonStr)
}