package helpers

// Version holds the version string for the alog-rs application
var Version string

// BuildVersion returns the version for the alog-rs application
func BuildVersion() string {
	if len(Version) == 0 {
		return "dev"
	}

	return Version
}
