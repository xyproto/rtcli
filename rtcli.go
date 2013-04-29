package main

import (
	"fmt"
	"github.com/xyproto/ais/wizard"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://rt.uio.no/"
	restURL = "/REST/1.0/"
)

type UserPass struct {
	user string
	pass string
}

func (up *UserPass) GetLoginString() string {
	return "?user=" + up.user + "&pass=" + up.pass
}

func (up *UserPass) GetItem(category, id string) string {
	var client http.Client
	url := baseURL + restURL + category + "/" + id + up.GetLoginString()
	resp, err := client.Get(url)
	if err != nil {
		panic("Could not get url:" + url)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Could not dump body")
	}
	return string(b)
}

func (up *UserPass) GetTicket(id string) string {
	return up.GetItem("ticket", id)
}

func NewUserPass(username, password string) *UserPass {
	return &UserPass{username, password}
}

func main() {
	var username, password string
	filename := "userpass.txt"
	if Exists(filename) {
		username, password = ReadTwoLines(filename)
	} else {
		username = wizard.Ask("Username: ")
		password = wizard.AskPassword("Password: ")
	}
	up := NewUserPass(username, password)
	ticketid := wizard.Ask("Ticked ID (for instance 905968): ")
	fmt.Println(up.GetTicket(ticketid))
}
