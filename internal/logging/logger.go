package logging

import (
	"io"
	"log"
)

func New(writer io.Writer) *log.Logger {
	return log.New(writer, "zero_control ", log.LstdFlags|log.Lmsgprefix)
}
