package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/junxxx/dockerhub-webhook/internal/runner"
)

type state int

const (
	Success state = iota
	Failure
	Cuserror
)

// response callback state
var StateText = map[state]string{
	Success:  "success",
	Failure:  "failure",
	Cuserror: "error",
}

// callbask payload
type CallBackPayload struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Context     string `json:"context"`
	TargetUrl   string `json:"target_url"`
}

// ignore other col
type Payload struct {
	CallbackUrl string `json:"callback_url"`
	Repository  struct {
		RepoUrl  string `json:"repo_url"`
		Name     string `json:"name"`
		RepoName string `json:"repo_name"`
	} `json:"repository"`
}

func HookHandler(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return
	}
	go runTask(payload)
	log.Println("payload", payload)
}

func (p *Payload) getScript() string {
	config := GetConf()
	for _, c := range config.Services {
		if c.RepoName == p.Repository.RepoName {
			return c.Script
		}
	}
	return ""
}

// run the correspond task
func runTask(p Payload) {
	script := p.getScript()
	log.Println("run local script:", script)
	if script != "" {
		runner.RunAndWait(script)
	}
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
