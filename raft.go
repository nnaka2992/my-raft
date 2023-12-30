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

func NewNode(s state) *StateMachine {
	sm := new(StateMachine)

	switch s {
	case LEADER:
		sm.State = s
		sm.CurrentTerm = 0
		sm.Log = make([]LogEntry, 10)
		sm.CommitIndex = 0
		sm.LastApplied = 0
		sm.NextIndex = make([]int, 10)
		sm.MatchIndex = make([]int, 10)
		sm.LeaderAddr = ""
		sm.FollowerAddr = make([]string, 10)
	case FOLLOWER:
	}
	return sm
}

func (s *StateMachine) NewFollower(addr string) *StateMachine {
	s.FollowerAddr = append(s.FollowerAddr, addr)

	sm := NewNode(FOLLOWER)
	return sm
}
