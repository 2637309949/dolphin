package core

// IsDebugging returns true if the framework is running in debug mode.
func IsDebugging() bool {
	return mode == debugCode
}
