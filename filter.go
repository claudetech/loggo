package loggo

// Filter interface is used to check if the log
// should be written by then appender
type Filter interface {
	ShouldLog(msg *Message) bool
}

// MinLogLevelFilter filters all messages
// with a log level strictly lower than MinLevel
type MinLogLevelFilter struct {
	MinLevel Level
}

// ShouldLog returns true if msg.Level is greater or equal to
// the filter MinLevel
func (f *MinLogLevelFilter) ShouldLog(msg *Message) bool {
	return msg.Level >= f.MinLevel
}

// MaxLogLevelFilter filters all messages
// with a log level strictly greater than MaxLevel
type MaxLogLevelFilter struct {
	MaxLevel Level
}

// ShouldLog returns true if msg.Level is lower or equal to
// the filter MaxLevel
func (f *MaxLogLevelFilter) ShouldLog(msg *Message) bool {
	return msg.Level <= f.MaxLevel
}
