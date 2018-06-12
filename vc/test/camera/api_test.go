package camera

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
	resp, err := http.Get("https://uyhelatdx1.execute-api.ap-northeast-1.amazonaws.com/dev/camera/1234567")
	if err != nil {
	// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestPost(t *testing.T) {
	camera := repository.Camera{Mac:"1234567", Camera:"0524"}
	c, _ := json.Marshal(camera);

	resp, _ := http.Post("https://uyhelatdx1.execute-api.ap-northeast-1.amazonaws.com/dev/camera",
		"application/x-www-form-urlencoded", strings.NewReader(string(c)))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestGetAll(t *testing.T) {

	req, _ :=
		http.NewRequest(
			"GET",
			"https://uyhelatdx1.execute-api.ap-northeast-1.amazonaws.com/dev/camera",
			nil)

	c := http.DefaultClient;

	resp, err := c.Do(req)
	if err != nil {
		t.Log(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

