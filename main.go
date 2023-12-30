package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var s *StateMachine
func main() {
	s = NewNode(LEADER)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//	e.POST("/raft/heartbeat", PostReceiveHeartbeat)
	e.GET("/raft/health", GetHealth)
	e.POST("/raft/follower/new", PostNewFollower)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetHealth(c echo.Context) error {
	s := struct { Status string }{ Status: "OK" }
	return c.JSON(http.StatusOK, s)
}

func PostNewFollower(c echo.Context) error {
	addr := c.QueryParam("client_addr")
	s := s.NewFollower(addr)
	return c.JSON(http.StatusOK, s)
}
