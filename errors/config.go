package errors

// return this error to save configurations on exit
type ConfigChanged struct{}

func (ConfigChanged) Error() string {
	return ""
}
