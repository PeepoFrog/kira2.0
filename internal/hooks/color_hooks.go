package hooks

import "github.com/sirupsen/logrus"

// ColorHook is a Logrus hook to force color output
type ColorHook struct{}

// Fire is called when a log event is fired
func (h *ColorHook) Fire(entry *logrus.Entry) error {
	entry.Logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	return nil
}

// Levels returns the log levels that this hook should be triggered for
func (h *ColorHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
