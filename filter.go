package loggo

type Filter interface {
	ShouldLog(msg *Message) bool
}

type LogLevelFilter struct {
	MinLevel Level
}

func (f *LogLevelFilter) ShouldLog(msg *Message) bool {
	return msg.Level >= f.MinLevel
}
