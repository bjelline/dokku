package registry

var (
	// DefaultProperties is a map of all valid registry properties with corresponding default property values
	DefaultProperties = map[string]string{
		"image-repo":      "",
		"push-on-release": "false",
		"server":          "",
	}

	// GlobalProperties is a map of all valid global registry properties
	GlobalProperties = map[string]bool{
		"image-repo-template": true,
		"push-on-release":     true,
		"server":              true,
	}
)
