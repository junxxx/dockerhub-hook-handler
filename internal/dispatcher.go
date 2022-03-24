package internal

type HandleFunc func()

type Handler struct {
	repoName string
	handler  HandleFunc
	script   string
}

func GetService(p Payload) {

}
