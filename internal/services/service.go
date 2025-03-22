package services

import "github.com/sirupsen/logrus"

type Service struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}
