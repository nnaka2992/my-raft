package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var s *StateMachine
func main() {
	//	s = StartRaftNode(FOLLOWER)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//	e.POST("/raft/heartbeat", PostReceiveHeartbeat)
	e.GET("/raft/health", GetHealth)
	e.Logger.Fatal(e.Start(":8080"))
}

// func PostReceiveHeartbeat(c echo.Context) {
// 	// TODO
// 	s.ReceiveHeartbeat()
// }

func GetHealth(c echo.Context) error {
	s := struct { Status string }{ Status: "OK" }
	return c.JSON(http.StatusOK, s)
}
