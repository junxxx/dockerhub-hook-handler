package handler

type ReadNewsService struct {
	image     string // docker image
	repoUrl   string // docker hub repo_url
	container string // container name
}

func GetInstance() *ReadNewsService {
	return &ReadNewsService{}
}

func (s *ReadNewsService) Run() {
}

func (s *ReadNewsService) CallBack() {

}
