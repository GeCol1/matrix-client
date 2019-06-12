package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"

	"gopkg.in/resty.v1"
)

type AuthSuccess struct {
}

type loginResponse struct {
	Access_token string
	Home_server  string
	User_id      string
	Device_id    string
}

func main() {
	pwd := getPasswordFromFile()

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
		"identifier": {
		  "type": "m.id.user",
		  "user": "geco"
		},
		"initial_device_display_name": "Jungle Phone",
		"password": "` + pwd + `",
		"type": "m.login.password"
	  }`)).
		SetResult(&AuthSuccess{}).
		Post("https://matrix.fuz.re/_matrix/client/r0/login")

	if err != nil {
		fmt.Println(err)
		panic("Ouch")
	}

	// fmt.Println(resp, err)

	var lr loginResponse
	err = json.Unmarshal(resp.Body(), &lr)
	if err!=nil {
		panic("not able to decode json")
	}

	fmt.Println (lr.Access_token)

	// respMsg, err := resty.R().
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody([]byte(`{"msgtype":"m.text", "body":"hello from go code"}`)).
	// 	Post("https://matrix.fuz.re/_matrix/client/r0/rooms/!lCdApgaICssmlPaSnq:matrix.fuz.re/send/m.room.message?access_token="+lr.Access_token)

	// if err!=nil {
	// 	fmt.Println(err)
	// 	panic("Not able to send message")
	// }

	// fmt.Println(respMsg)

	respMsg, err := resty.R().Get("https://matrix.fuz.re/_matrix/client/api/v1/rooms/!lCdApgaICssmlPaSnq:matrix.fuz.re/messages?access_token=" + lr.Access_token + "&from=END&dir=b&limit=10")
	fmt.Println(respMsg, err)

}

func simpleGet() {
	resp, err := resty.R().Get("http://httpbin.org/get")
	if err != nil {
		panic("Aie aie aie")
	}

	fmt.Println(string(resp.Body()))
}

func getPasswordFromFile() (pwd string) {
	file, err := ioutil.ReadFile("riot.pass")
	if err != nil {
		panic("Failed to open pass file")
	}
	
	return string(file)
}
