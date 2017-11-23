package fsm_test

import (
	"testing"

	"github.com/goph/fsm"
	"github.com/goph/fsm/internal/mocks"
)

func TestActionMuxDelegate(t *testing.T) {
	delegate := new(mocks.Delegate)
	delegate.On("Handle", "action", "fromState", "toState", []interface{}{"argument"}).Return()

	delegates := map[string]fsm.Delegate{
		"action": delegate,
	}

	amd := fsm.NewActionMuxDelegate(delegates)

	amd.Handle("action", "fromState", "toState", []interface{}{"argument"})

	delegate.AssertExpectations(t)
}

func TestActionMuxDelegate_NoAction(t *testing.T) {
	delegate := new(mocks.Delegate)

	delegates := map[string]fsm.Delegate{
		"action": delegate,
	}

	amd := fsm.NewActionMuxDelegate(delegates)

	amd.Handle("other_action", "fromState", "toState", []interface{}{"argument"})

	delegate.AssertNotCalled(t, "Handle", "action", "fromState", "toState", []interface{}{"argument"})
}

func TestActionMuxDelegate_SetStateMachine(t *testing.T) {
	delegate := new(mocks.Delegate)
	smaDelegate := new(mocks.StateMachineAwareDelegate)
	sm := new(fsm.StateMachine)

	smaDelegate.On("SetStateMachine", sm)

	type combinedDelegate struct {
		fsm.Delegate
		fsm.StateMachineAwareDelegate
	}

	delegate1 := &combinedDelegate{
		delegate,
		smaDelegate,
	}

	delegate2 := new(mocks.Delegate)

	delegates := map[string]fsm.Delegate{
		"action1": delegate1,
		"action2": delegate2,
	}

	amd := fsm.NewActionMuxDelegate(delegates)

	amd.SetStateMachine(sm)

	smaDelegate.AssertExpectations(t)
}

func TestCompositeDelegate(t *testing.T) {
	delegate1 := new(mocks.Delegate)
	delegate1.On("Handle", "action", "fromState", "toState", []interface{}{"argument"}).Return()

	delegate2 := new(mocks.Delegate)
	delegate2.On("Handle", "action", "fromState", "toState", []interface{}{"argument"}).Return()

	delegates := []fsm.Delegate{
		delegate1,
		delegate2,
	}

	cd := fsm.NewCompositeDelegate(delegates)

	cd.Handle("action", "fromState", "toState", []interface{}{"argument"})

	delegate1.AssertExpectations(t)
	delegate2.AssertExpectations(t)
}

func TestCompositeDelegate_SetStateMachine(t *testing.T) {
	delegate := new(mocks.Delegate)
	smaDelegate := new(mocks.StateMachineAwareDelegate)
	sm := new(fsm.StateMachine)

	smaDelegate.On("SetStateMachine", sm)

	type combinedDelegate struct {
		fsm.Delegate
		fsm.StateMachineAwareDelegate
	}

	delegate1 := &combinedDelegate{
		delegate,
		smaDelegate,
	}

	delegate2 := new(mocks.Delegate)

	delegates := []fsm.Delegate{
		delegate1,
		delegate2,
	}

	cd := fsm.NewCompositeDelegate(delegates)

	cd.SetStateMachine(sm)

	smaDelegate.AssertExpectations(t)
}
