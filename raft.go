package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// state type is a enum that holds the type of the raft node
type state int
const (
	LEADER state = iota
	FOLLOWER
	CANDIDATE
)

type StateMachine struct {
	State state
	// Persistent state on all servers
	CurrentTerm int
	VotedFor int
	Log []LogEntry
	// Volatile state on all servers
	CommitIndex int
	LastApplied int
	// Volatile state on leaders
	NextIndex []int
	MatchIndex []int
	// Other nodes info
	LeaderAddr string
	FollowerAddr []string
}

type Command interface {
	Apply() error
}

type LogEntry struct {
	Term int
	Command Command
}

func (s *StateMachine) AppendEntries() {
	// TODO
}

func (s *StateMachine) RequestVote() {
	// TODO
}

func (s *StateMachine) SendHeartbeat() {
	// TODO
}

func (s *StateMachine) ReceiveHeartbeat() {
	// TODO

}

func (s *StateMachine) ApplyLog() {
	// TODO
}

// the following function call SendHeartbeat() periodically
func (s *StateMachine) HeartbeatTimer() {
	// TODO
}

func (s *StateMachine) ElectionTimer() {
	// TODO
}

func NewNode(st state, addr string, leader_addr string) (*StateMachine, error) {
	sm := new(StateMachine)
	c := http.Client{}
	switch st {
	case LEADER:
		sm.State = st
		sm.CurrentTerm = 0
		sm.Log = make([]LogEntry, 0)
		sm.CommitIndex = 0
		sm.LastApplied = 0
		sm.NextIndex = make([]int, 0)
		sm.MatchIndex = make([]int, 0)
		sm.LeaderAddr = addr
		sm.FollowerAddr = make([]string, 0)
	case FOLLOWER:
		if leader_addr == "" {
			return nil, fmt.Errorf("leader address is empty")
		}
		rp, err := http.Post("http://" + leader_addr + "/raft/follower/new?client_addr=" + addr, "", nil)
		if err != nil {
			rd, _ := http.NewRequest("DELETE", "http://" + leader_addr + "/raft/follower?client_addr=" + addr, nil)
			c.Do(rd)
			return nil, err
		}
		defer rp.Body.Close()

		rg, err := http.Get("http://" + leader_addr + "/raft/statemachine")
		if err != nil {
			return nil, err
		}
		defer rg.Body.Close()

		if rg.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("leader is not ready")
		}
		if err := json.NewDecoder(rg.Body).Decode(sm); err != nil {
			return nil, err
		}
		sm.State = st
	}
	return sm, nil
}

func NewFollower(s *StateMachine, addr string) {
	s.FollowerAddr = append(s.FollowerAddr, addr)
}

func (s *StateMachine) DeleteFollower(addr string) {
	for i, a := range s.FollowerAddr {
		if a == addr {
			s.FollowerAddr = append(s.FollowerAddr[:i], s.FollowerAddr[i+1:]...)
			break
		}
	}
}
