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
