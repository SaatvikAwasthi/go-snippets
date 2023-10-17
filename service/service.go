package service

import "fmt"

type service struct {
	messageFormat string
}

type Service interface {
	GetMessage(fName, lName string) string
}

func New() Service {
	return &service{
		messageFormat: "Hello, %s %s. Welcome to go snippets.",
	}
}

func (s *service) GetMessage(fName, lName string) string {
	return fmt.Sprintf(s.messageFormat, fName, lName)
}
