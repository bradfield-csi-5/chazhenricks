package sm

import (
	"errors"
	"sync"
)

// error returned when cannot process event for the state were currently in
var ErrEventRejected = errors.New("event rejected")

//oopsie error. Probably a config issue? who knows
var Oopsie = errors.New("Oopsie - probs a config error, idk")

const (
	//initial default state of the system
	Default StateType = ""

	//noop event
	NoOp EventType = "NoOp"
)

//Names of the States 
type StateType string

//names of the Events
type EventType string

//catch-all interface for user-defined context. 
//used to pass some additional data in with an event 
type EventContext interface{}

//Action to be executed in a given event. 
type Action interface {
	Execute(eventCtx EventContext) EventType
}

type Events map[EventType]StateType

//Actions and Events tied to an State
type State struct {
	Action Action
	Events Events
}

//map that maps the state Type with the Actions/Events for that state
type States map[StateType]State

//Data Structure that holds all the state 
type StateMachine struct {
	Previous StateType
	Current  StateType
	States   States
	//ensures that only 1 event is processed at a time
	mutex sync.Mutex
}

//returns the next state for the event given the machines current state, or an
//error if the event cant be handled
func (s *StateMachine) getNextState(event EventType) (StateType, error){
  //get current states Actions and Events from map
  if state, ok := s.States[s.Current]; ok {
    //if the current state has events is can respond to 
    if state.Events != nil {
      //if the event exists in the list of events we can respond to
      if next, ok := state.Events[event]; ok {
        return next, nil
      }
    }
  }

  //if any of those checks fail, retrn rejected
  return Default, ErrEventRejected
}


//SendEvent sends the event to the machine 
func (s *StateMachine) SendEvent (event EventType, eventCtx EventContext) error {
  //lock up front and unlock when were done
  s.mutex.Lock()
  defer s.mutex.Unlock()

  for {

    //whats the next state for the machines current state
    nextState, err := s.getNextState(event)
    if err != nil {
      return ErrEventRejected
    }

    //Identify what the next state should be
    state, ok := s.States[nextState]
    if !ok || state.Action == nil {
     return Oopsie 
    }

    //transition state to the next one
    s.Previous = s.Current 
    s.Current = nextState 

    //execute the next state action and do one mo 'gin if 
    //the event returned an action. This will probs nooop but just in case
    nextEvent := state.Action.Execute(eventCtx)
    if nextEvent == NoOp{
      return nil
    }
    event = nextEvent
  }
}


