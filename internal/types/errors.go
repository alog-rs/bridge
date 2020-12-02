package types

// Error represents an internal non GRPC error that the application may encounter
type Error int

const (
	// ErrorRequestFailed is returned when a request failed
	ErrorRequestFailed Error = iota
	// ErrorResourceNotFound is returned when the requested resource was not found
	ErrorResourceNotFound
	// ErrorResourceNotPublic is returned when the requested resource is private
	ErrorResourceNotPublic
	// ErrorInternal is returned when there was an internal application error
	ErrorInternal
	// ErrorUnknown is returned when an unknown error occured
	ErrorUnknown
	// ErrorNone is returned when there was no error returned
	ErrorNone
)

// IsPresent is a helper which checks if the error is not ErrorNone
func (e Error) IsPresent() bool {
	return e != ErrorNone
}
