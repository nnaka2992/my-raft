package main

import (
	"os"
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var s *StateMachine
var (
	sState = flag.String("state", "LEADER", "state of the node	(LEADER, FOLLOWER)")
	sAddr = flag.String("addr", "localhost", "address of the node")
	sPort = flag.String("port", "8080", "port of the node")
	bHelp = flag.Bool("help", false, "show this message")
	bHelpShort = flag.Bool("h", false, "show this message")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [<state> <addr> <port>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *bHelp || *bHelpShort {
		flag.Usage()
		os.Exit(0)
	}
	switch *sState {
	case "LEADER":
		s = NewNode(LEADER, *sAddr + ":" + *sPort)
	case "FOLLOWER":
		s = NewNode(FOLLOWER, *sAddr + ":" + *sPort)
	default:
		flag.Usage()
		os.Exit(0)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//	e.POST("/raft/heartbeat", PostReceiveHeartbeat)
	e.GET("/raft/health", GetHealth)
	e.POST("/raft/follower/new", PostNewFollower)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *sPort)))
}

func GetHealth(c echo.Context) error {
	s := struct { Status string }{ Status: "OK" }
	return c.JSON(http.StatusOK, s)
}

func PostNewFollower(c echo.Context) error {
	addr := c.QueryParam("client_addr")
	s := NewFollower(s, addr)
	return c.JSON(http.StatusOK, s)
}
