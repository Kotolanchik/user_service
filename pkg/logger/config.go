package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Config struct {
	Formatter logrus.Formatter
	Out       io.Writer
}

func GetDefaultConfig() Config {
	return Config{
		FormatterJSON,
		os.Stderr,
	}
}
