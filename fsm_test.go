package fsm_test

import (
	"testing"

	"github.com/goph/fsm"
	"github.com/goph/fsm/internal/mocks"
)

func TestStateMachine_DelegateInvoked(t *testing.T) {
	delegate := new(mocks.Delegate)
	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
			Action:    "action",
		},
	}

	delegate.On("Handle", "action", "current_state", "next_state", []interface{}{"argument"}).Return()

	sm := fsm.NewStateMachine(delegate, transitions)

	sm.Trigger("current_state", "event", "argument")

	delegate.AssertExpectations(t)
}

func TestStateMachine_FirstTransition(t *testing.T) {
	delegate := new(mocks.Delegate)
	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
			Action:    "action",
		},
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "other_next_state",
			Action:    "other_action",
		},
	}

	delegate.On("Handle", "action", "current_state", "next_state", []interface{}{"argument"}).Return()

	sm := fsm.NewStateMachine(delegate, transitions)

	sm.Trigger("current_state", "event", "argument")

	delegate.AssertExpectations(t)
}

func TestStateMachine_NoAction(t *testing.T) {
	delegate := new(mocks.Delegate)
	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
		},
	}

	sm := fsm.NewStateMachine(delegate, transitions)

	sm.Trigger("current_state", "event", "argument")

	delegate.AssertNotCalled(t, "Handle", "action", "current_state", "next_state", []interface{}{"argument"})
}

func TestStateMachine_Subject(t *testing.T) {
	delegate := new(mocks.Delegate)
	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
			Action:    "action",
		},
	}

	subject := new(mocks.Subject)

	subject.On("State").Return("current_state")

	delegate.On("Handle", "action", "current_state", "next_state", []interface{}{subject, "argument"}).Return()

	sm := fsm.NewStateMachine(delegate, transitions)

	sm.TriggerSubject(subject, "event", "argument")

	delegate.AssertExpectations(t)
	subject.AssertExpectations(t)
}
