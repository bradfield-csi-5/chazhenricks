package main

import (
	"fmt"
	"main/sm"
)

const (
	Off sm.StateType = "Off"
	On  sm.StateType = "On"

	SwitchOff sm.EventType = "SwitchOff"
	SwitchOn  sm.EventType = "SwitchOn"
)

type OffAction struct{}

func (a *OffAction) Execute(ctx sm.EventContext) sm.EventType {
	fmt.Println("The light has been switched off")

	return sm.NoOp
}

type OnAction struct{}

func (a *OnAction) Execute(ctx sm.EventContext) sm.EventType {
	fmt.Println("The light has been switched on")

	return sm.NoOp
}

func newLightSwitchFSM() *sm.StateMachine {
	return &sm.StateMachine{
		States: sm.States{
			sm.Default: sm.State{
				Events: sm.Events{
					SwitchOff: Off,
				},
			},
			Off: sm.State{
				Action: &OffAction{},
				Events: sm.Events{
					SwitchOn: On,
				},
			},
			On: sm.State{
				Action: &OnAction{},
				Events: sm.Events{
					SwitchOff: Off,
				},
			},
		},
	}
}

func main() {
// Create a new instance of the light switch state machine.
	lightSwitchFsm := newLightSwitchFSM()

	// Set the initial "off" state in the state machine.
	err := lightSwitchFsm.SendEvent(SwitchOff, nil)
	if err != nil {
		fmt.Printf("Couldn't set the initial state of the state machine, err: %v", err)
	}

	// Send the switch-off event again and expect the state machine to return an error.
	err = lightSwitchFsm.SendEvent(SwitchOff, nil)
	if err != sm.ErrEventRejected {
		fmt.Printf("Expected the event rejected error, got nil")
	}

	// Send the switch-on event and expect the state machine to transition to the
	// "on" state.
	err = lightSwitchFsm.SendEvent(SwitchOn, nil)
	if err != nil {
		fmt.Printf("Couldn't switch the light on, err: %v", err)
	}

	// Send the switch-on event again and expect the state machine to return an error.
	err = lightSwitchFsm.SendEvent(SwitchOn, nil)
	if err != sm.ErrEventRejected {
		fmt.Printf("Expected the event rejected error, got nil")
	}

	// Send the switch-off event and expect the state machine to transition back
	// to the "off" state.
	err = lightSwitchFsm.SendEvent(SwitchOff, nil)
	if err != nil {
		fmt.Printf("Couldn't switch the light off, err: %v", err)
	}
}
