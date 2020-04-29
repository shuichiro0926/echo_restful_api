package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var AppLog *logrus.Logger = logrus.New()

func Init() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	AppLog.Out = file
	AppLog.Formatter = &logrus.JSONFormatter{}
}
