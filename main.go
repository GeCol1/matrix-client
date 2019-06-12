package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/valyala/fastjson"
	"gopkg.in/resty.v1"
)

type loginResponse struct {
	AccessToken string `json:"Access_token"`
	HomeServer  string `json:"Home_server"`
	UserID      string `json:"User_Id"`
	DeviceID    string `json:"Device_id"`
}

// BaseURL of the matrix server
const BaseURL = "matrix.fuz.re"

// RoomID of the room to connect. You can get in the advanced parameters in the riot client.
const RoomID = "!lCdApgaICssmlPaSnq:matrix.fuz.re"

func main() {
	// put your user here
	token := getToken("geco")

	readLatestMessages(token)
}

func readLatestMessages(token string) {
	respMsg, err := resty.R().Get("https://" + BaseURL + "/_matrix/client/api/v1/rooms/" + RoomID + "/messages?access_token=" + token + "&from=END&dir=b&limit=10")
	checkErr(err, "Could not get the messages from matrix API")
	// fmt.Printf("%s\n\n---------------\n", respMsg)

	var p fastjson.Parser
	m, err := p.Parse(string(respMsg.Body()))
	checkErr(err, "Could not decode json of messages")
	vals := m.GetArray("chunk")
	for _, val := range vals {
		fmt.Printf("  (%s) > %s\n", val.GetStringBytes("sender"), val.GetStringBytes("content", "body"))
	}
}

func postMessage(msg string, token string) {
	respMsg, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"msgtype":"m.text", "body":"hello from go code"}`)).
		Post("https://" + BaseURL + "/_matrix/client/r0/rooms/" + RoomID + "/send/m.room.message?access_token=" + token)

	checkErr(err, "Could not post the message")
	fmt.Println(respMsg)
}

func getToken(user string) string {
	pwd := getPasswordFromFile()

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
		"identifier": {
		  "type": "m.id.user",
		  "user": "` + user + `"
		},
		"initial_device_display_name": "Jungle Phone",
		"password": "` + pwd + `",
		"type": "m.login.password"
	  }`)).
		// SetResult(&AuthSuccess{}).
		Post("https://" + BaseURL + "/_matrix/client/r0/login")

	checkErr(err, "Could not authenticate")
	// fmt.Println(resp, err)

	var lr loginResponse
	err = json.Unmarshal(resp.Body(), &lr)
	checkErr(err, "Could not decode json of authentication")

	// fmt.Println(lr.Access_token)
	return lr.AccessToken
}

func simpleGet() {
	resp, err := resty.R().Get("http://httpbin.org/get")
	checkErr(err, "")

	fmt.Println(string(resp.Body()))
}

func getPasswordFromFile() (pwd string) {
	file, err := ioutil.ReadFile("riot.pass")
	checkErr(err, "Failed to open pass file")

	return string(file)
}

func checkErr(err error, errorMessage string) {
	if err != nil {
		log.Fatal(errorMessage)
		panic(err)
	}
}
