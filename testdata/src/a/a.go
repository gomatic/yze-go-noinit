package a

// init is the special package init function and must be flagged.
func init() { _ = 0 } // want `func init is not permitted`

// setup is a normal function and must NOT be flagged.
func setup() {}

// T carries a method named init to prove the receiver guard.
type T struct{}

// init here is an ordinary method (it has a receiver), not the special package
// init function, and must NOT be flagged.
func (T) init() {}
