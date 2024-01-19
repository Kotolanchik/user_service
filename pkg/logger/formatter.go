package logging

import "github.com/sirupsen/logrus"

var (
	FormatterJSON = &logrus.JSONFormatter{}
	FormatterText = &logrus.TextFormatter{}
)
