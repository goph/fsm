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
