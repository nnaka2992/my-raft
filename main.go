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

func main() {
	sState := flag.String("state", "", "state of the node	(LEADER, FOLLOWER)")
	sAddr := flag.String("addr", "", "address of the node")
	sPort := flag.String("port", "", "port of the node")
	sLeaderAddr := flag.String("leader-addr", "", "address of the leader node. Only used when the state is FOLLOWER")
	sLeaderPort := flag.String("leader-port", "", "port of the leader node. Only used when the state is FOLLOWER")
	bHelp := flag.Bool("help", false, "show this message")
	bHelpShort := flag.Bool("h", false, "show this message")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -state=<state> -addr=<addr> -port=<port> [-leader-addr=<leader-addr> - leader-port=<leader-port>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *bHelp || *bHelpShort {
		flag.Usage()
		os.Exit(0)
	}
	switch *sState {
	case "LEADER":
		var err error
		if *sAddr == "" || *sPort == "" {
			fmt.Fprintf(os.Stderr, "Usage: %s -state=LEADER -addr=<addr> -port=<port>\n", os.Args[0])
			os.Exit(1)
		}
		s, err = NewNode(LEADER, *sAddr + ":" + *sPort, "")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	case "FOLLOWER":
		var err error
		if *sAddr == "" || *sPort == "" || *sLeaderAddr == "" || *sLeaderPort == "" {
			fmt.Fprintf(os.Stderr, "Usage: %s -state=FOLLOWER -addr=<addr> -port=<port> -leader-addr=<leader-addr> -leader-port=<leader-port>\n", os.Args[0])
			os.Exit(1)
		}
		s, err = NewNode(FOLLOWER, *sAddr + ":" + *sPort, *sLeaderAddr + ":" + *sLeaderPort)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(2)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//	e.POST("/raft/heartbeat", PostReceiveHeartbeat)
	e.GET("/raft/health", GetHealth)
	e.GET("/raft/statemachine", GetStateMachine)
	e.POST("/raft/follower/new", PostNewFollower)
	e.DELETE("/raft/follower", DeleteFollower)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *sPort)))
}

func GetHealth(c echo.Context) error {
	r := struct { Status string }{ Status: "OK" }
	return c.JSON(http.StatusOK, r)
}

func GetStateMachine(c echo.Context) error {
	return c.JSON(http.StatusOK, s)
}

func PostNewFollower(c echo.Context) error {
	addr := c.QueryParam("client_addr")
	NewFollower(s, addr)
	r := struct { Status string }{ Status: "OK" }
	return c.JSON(http.StatusOK, r)
}

func DeleteFollower(c echo.Context) error {
	addr := c.QueryParam("client_addr")
	s.DeleteFollower(addr)
	r := struct { Status string }{ Status: "OK" }
	return c.JSON(http.StatusOK, r)
}
