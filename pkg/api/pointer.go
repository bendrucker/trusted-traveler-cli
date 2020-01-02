package ttapi

// Bool returns a pointer for a boolean value
func Bool(v bool) *bool {
	return &v
}

// String returns a pointer for a string value
func String(v string) *string {
	return &v
}
