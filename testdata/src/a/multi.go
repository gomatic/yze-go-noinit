package a

// Go permits multiple func init declarations in a single file; every one is
// the special package init function and each must be flagged independently.

func init() { _ = 1 } // want `func init is not permitted`

func init() { _ = 2 } // want `func init is not permitted`

func init() { _ = 3 } // want `func init is not permitted`
