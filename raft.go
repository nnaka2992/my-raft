package main

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

func NewNode(st state, addr string) *StateMachine {
	sm := new(StateMachine)

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
		sm.State = st
		sm.CurrentTerm = s.CurrentTerm
		sm.VotedFor = s.VotedFor
		sm.Log = s.Log
		sm.CommitIndex = s.CommitIndex
		sm.LastApplied = s.LastApplied
		sm.NextIndex = s.NextIndex
		sm.MatchIndex = s.MatchIndex
		sm.LeaderAddr = s.LeaderAddr
		sm.FollowerAddr = s.FollowerAddr
	}
	return sm
}

func (s *StateMachine) NewFollower(addr string) *StateMachine {
	s.FollowerAddr = append(s.FollowerAddr, addr)

	sm := NewNode(FOLLOWER, addr)
	return sm
}
