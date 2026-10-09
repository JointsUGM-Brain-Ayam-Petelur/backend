package auth

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Test(response string) string {
	return response
}
