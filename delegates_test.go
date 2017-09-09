package fsm_test

import (
	"testing"

	. "github.com/goph/fsm"
	"github.com/goph/fsm/internal/mocks"
)

func TestActionMuxDelegate(t *testing.T) {
	delegate := new(mocks.Delegate)
	delegate.On("Handle", "action", "fromState", "toState", []interface{}{"argument"}).Return()

	delegates := map[string]Delegate{
		"action": delegate,
	}

	amd := NewActionMuxDelegate(delegates)

	amd.Handle("action", "fromState", "toState", []interface{}{"argument"})

	delegate.AssertExpectations(t)
}

func TestActionMuxDelegate_NoAction(t *testing.T) {
	delegate := new(mocks.Delegate)

	delegates := map[string]Delegate{
		"action": delegate,
	}

	amd := NewActionMuxDelegate(delegates)

	amd.Handle("other_action", "fromState", "toState", []interface{}{"argument"})

	delegate.AssertNotCalled(t, "Handle", "action", "fromState", "toState", []interface{}{"argument"})
}

func TestCompositeDelegate(t *testing.T) {
	delegate1 := new(mocks.Delegate)
	delegate1.On("Handle", "action", "fromState", "toState", []interface{}{"argument"}).Return()

	delegate2 := new(mocks.Delegate)
	delegate2.On("Handle", "action", "fromState", "toState", []interface{}{"argument"}).Return()

	delegates := []Delegate{
		delegate1,
		delegate2,
	}

	cd := NewCompositeDelegate(delegates)

	cd.Handle("action", "fromState", "toState", []interface{}{"argument"})

	delegate1.AssertExpectations(t)
	delegate2.AssertExpectations(t)
}
