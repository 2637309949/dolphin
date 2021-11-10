package core

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

var mode = debugCode
var modeName = DebugMode

// SetMode sets gin mode according to input string.
func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	switch value {
	case DebugMode:
		mode = debugCode
	case ReleaseMode:
		mode = releaseCode
	case TestMode:
		mode = testCode
	default:
		panic("mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}
