package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/OuterInside/party/server/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3000, "server listen port")

	log.SetFlags(log.Llongfile)
}

func main() {
	e := echo.New()
	routes.New(e)

	log.Printf("PartyServer start (listen: %d)", port)
	e.Run(fasthttp.New(":" + strconv.Itoa(port)))
}
