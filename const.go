package virtualbox

import "time"

const (
	stringYes string = "Yes"
	osWindows string = "windows"

	// StopTimeout is the maximum time to wait for a VM to stop gracefully
	StopTimeout = 1 * time.Minute
)
