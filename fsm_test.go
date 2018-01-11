package fsm_test

import (
	"testing"

	"github.com/goph/fsm"
	"github.com/goph/fsm/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"errors"
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

	delegate.On("Handle", "action", "current_state", "next_state", []interface{}{"argument"}).Return(nil)

	sm := fsm.NewStateMachine(delegate, transitions)

	err := sm.Trigger("current_state", "event", "argument")

	require.NoError(t, err)

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

	delegate.On("Handle", "action", "current_state", "next_state", []interface{}{"argument"}).Return(nil)

	sm := fsm.NewStateMachine(delegate, transitions)

	err := sm.Trigger("current_state", "event", "argument")

	require.NoError(t, err)

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

	err := sm.Trigger("current_state", "event", "argument")

	require.NoError(t, err)

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

	subject.On("GetState").Return("current_state")

	delegate.On("Handle", "action", "current_state", "next_state", []interface{}{subject, "argument"}).Return(nil)

	sm := fsm.NewStateMachine(delegate, transitions)

	err := sm.TriggerSubject(subject, "event", "argument")

	require.NoError(t, err)

	delegate.AssertExpectations(t)
	subject.AssertExpectations(t)
}

func TestStateMachine_InvalidTransition(t *testing.T) {
	delegate := new(mocks.Delegate)
	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
		},
	}

	sm := fsm.NewStateMachine(delegate, transitions)

	err := sm.Trigger("other_current_state", "event", "argument")

	require.Error(t, err)

	ierr := err.(*fsm.InvalidTransitionError)

	assert.EqualError(t, ierr, "cannot transition from \"other_current_state\" state triggered by \"event\" event")
	assert.Equal(t, "other_current_state", ierr.CurrentState())
	assert.Equal(t, "event", ierr.Event())
	assert.Equal(t, []interface{}{"argument"}, ierr.Arguments())

	delegate.AssertNotCalled(t, "Handle", "action", "current_state", "next_state", []interface{}{"argument"})
}

func TestStateMachine_StateMachineAwareDelegate(t *testing.T) {
	delegate := new(mocks.Delegate)
	smaDelegate := new(mocks.StateMachineAwareDelegate)

	smaDelegate.On("SetStateMachine", mock.Anything)

	type combinedDelegate struct {
		fsm.Delegate
		fsm.StateMachineAwareDelegate
	}

	delegate1 := &combinedDelegate{
		delegate,
		smaDelegate,
	}

	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
		},
	}

	sm := fsm.NewStateMachine(delegate1, transitions)

	smaDelegate.AssertCalled(t, "SetStateMachine", sm)
}

func TestStateMachine_DelegateError(t *testing.T) {
	delegate := new(mocks.Delegate)
	transitions := []fsm.Transition{
		{
			FromState: "current_state",
			Event:     "event",
			ToState:   "next_state",
			Action:    "action",
		},
	}

	delegateErr := errors.New("error happened")

	delegate.
		On("Handle", "action", "current_state", "next_state", []interface{}{"argument"}).
		Return(delegateErr)

	sm := fsm.NewStateMachine(delegate, transitions)

	err := sm.Trigger("current_state", "event", "argument")

	require.Error(t, err)

	derr := err.(*fsm.DelegateError)

	assert.EqualError(t, derr, "delegate reported an error during transition from \"current_state\" state triggered by \"event\" event: error happened")
	assert.Equal(t, delegateErr, derr.Cause())
	assert.Equal(t, "next_state", derr.NextState())
	assert.Equal(t, "action", derr.Action())
	assert.Equal(t, "current_state", derr.CurrentState())
	assert.Equal(t, "event", derr.Event())
	assert.Equal(t, []interface{}{"argument"}, derr.Arguments())

	delegate.AssertExpectations(t)
}
