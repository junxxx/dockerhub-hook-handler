package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type state int

const (
	Success state = iota
	Failure
	Cuserror
)

// response callback state
var stateText = map[state]string{
	Success:  "success",
	Failure:  "failure",
	Cuserror: "error",
}

// callbask payload
type callBackPayload struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Context     string `json:"context"`
	TargetUrl   string `json:"target_url"`
}

// ignore other col
type Payload struct {
	CallbackUrl string `json:"callback_url"`
	// PushData    struct {
	// } `json:"push_data"`
	Repository struct {
		RepoUrl string `json:"repo_url"`
	} `json:"repository"`
}

func HookHanlder(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return
	}
	runTask(payload)
	log.Println("payload", payload)
}

// run the correspond task
func runTask(p Payload) {
	fmt.Println("pull the latest image and run it")
}

func callBack() {
	// change later
	url := "http://localhost:8081"
	payload := strings.NewReader("....")

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	res.Body.Close()
}
