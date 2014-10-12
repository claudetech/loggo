package loggo

type Filter interface {
	ShouldLog(msg *Message) bool
}

type MinLogLevelFilter struct {
	MinLevel Level
}

func (f *MinLogLevelFilter) ShouldLog(msg *Message) bool {
	return msg.Level >= f.MinLevel
}

type MaxLogLevelFilter struct {
	MaxLevel Level
}

func (f *MaxLogLevelFilter) ShouldLog(msg *Message) bool {
	return msg.Level <= f.MaxLevel
}
