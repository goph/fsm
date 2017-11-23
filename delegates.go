package fsm

// ActionMuxDelegate allows to register a set of delegates per action.
type ActionMuxDelegate struct {
	delegates map[string]Delegate
}

// NewActionMuxDelegate returns a new ActionMuxDelegate.
func NewActionMuxDelegate(delegates map[string]Delegate) *ActionMuxDelegate {
	return &ActionMuxDelegate{delegates}
}

// Handle calls the underlying delegate for an action if any.
func (d *ActionMuxDelegate) Handle(action string, fromState string, toState string, args []interface{}) {
	if delegate, ok := d.delegates[action]; ok {
		delegate.Handle(action, fromState, toState, args)
	}
}

func (d *ActionMuxDelegate) SetStateMachine(sm *StateMachine) {
	for _, delegate := range d.delegates {
		if smaDelegate, ok := delegate.(StateMachineAwareDelegate); ok {
			smaDelegate.SetStateMachine(sm)
		}
	}
}

// CompositeDelegate allows to multiplex the single delegate in the state machine.
type CompositeDelegate struct {
	delegates []Delegate
}

// NewCompositeDelegate returns a new CompositeDelegate.
func NewCompositeDelegate(delegates []Delegate) *CompositeDelegate {
	return &CompositeDelegate{delegates}
}

// Handle calls the underlying delegates.
func (d *CompositeDelegate) Handle(action string, fromState string, toState string, args []interface{}) {
	for _, delegate := range d.delegates {
		delegate.Handle(action, fromState, toState, args)
	}
}

func (d *CompositeDelegate) SetStateMachine(sm *StateMachine) {
	for _, delegate := range d.delegates {
		if smaDelegate, ok := delegate.(StateMachineAwareDelegate); ok {
			smaDelegate.SetStateMachine(sm)
		}
	}
}
