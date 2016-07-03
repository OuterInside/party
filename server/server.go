package main

import (
	"flag"
	"io"
	"log/syslog"
	"os"
	"strconv"

	"github.com/OuterInside/party/server/routes"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

const tag = "PartyServer"

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3000, "server listen port")

	// set logger format
	writers := make([]io.Writer, 0, 0)
	syslogWriter, _ := syslog.New(syslog.LOG_NOTICE|syslog.LOG_USER, tag)
	writers = append(writers, syslogWriter)
	writers = append(writers, os.Stdout)
	writer := io.MultiWriter(writers...)
	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	routes.New(e)

	logrus.Infof("PartyServer start (listen: %d)", port)
	e.Run(fasthttp.New(":" + strconv.Itoa(port)))
}
