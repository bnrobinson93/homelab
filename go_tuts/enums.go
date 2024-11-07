package main

import "fmt"

type ServerState int

const (
  StateIdle ServerState = iota // This generates sequentially, 0
  StateConnected // 1
  StateError     // 2
  StateRetrying  // 3
)

var stateName = map[ServerState]string{
  StateIdle:      "Idle",
  StateConnected: "Connected",
  StateError:     "Error",
  StateRetrying:  "Retrying",
}

func (ss ServerState) String() string {
  return stateName[ss]
}

func main() {
  ns := transition(StateIdle)
  fmt.Println(ns)

  ns2 := transition(ns)
  fmt.Println(ns2)

  ns3 := transition(3)
  fmt.Println(ns3)
}

func transition(s ServerState) ServerState {
  switch s {
  case StateIdle:
    return StateConnected
  case StateConnected, StateRetrying:
    return StateIdle
  case StateError:
    return StateError
  default:
    panic(fmt.Errorf("unknown state: %s\n", s))
  }
}
