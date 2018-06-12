package cauth_test

import (
	"fmt"
	"strings"
	"encoding/json"
	"testing"
	"net/http"
	"io/ioutil"
	"logitech.com/vc/lib/repository"
)


func TestGet(t *testing.T) {
	req, _ :=
		http.NewRequest(
			"GET",
			"https://uyhelatdx1.execute-api.ap-northeast-1.amazonaws.com/prod/camera/1237",
			nil)
	req.Header.Set("x-api-key", "elq4U1ZtEkazTEZTYqw0U6A0xGxUA5XoaqhdhDbi")

	c := http.DefaultClient;

	resp, err := c.Do(req)
	if err != nil {
		t.Log(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestPost(t *testing.T) {
	camera := repository.Camera{Mac:"1237", Camera:"0524"}
	c, _ := json.Marshal(camera);


	req, _ :=
		http.NewRequest(
			"POST",
			"https://uyhelatdx1.execute-api.ap-northeast-1.amazonaws.com/prod/camera",
			strings.NewReader(string(c)))
	req.Header.Set("x-api-key", "elq4U1ZtEkazTEZTYqw0U6A0xGxUA5XoaqhdhDbi")

	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		t.Log(err.Error())
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestGetAll(t *testing.T) {

	req, _ :=
		http.NewRequest(
			"GET",
			"https://uyhelatdx1.execute-api.ap-northeast-1.amazonaws.com/prod/camera",
			nil)
	req.Header.Set("x-api-key", "elq4U1ZtEkazTEZTYqw0U6A0xGxUA5XoaqhdhDbi")

	c := http.DefaultClient

	resp, err := c.Do(req)
	if err != nil {
		t.Log(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

