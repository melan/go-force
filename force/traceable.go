package force

import "fmt"

type Traceable interface {
	TraceOn(prefix string, logger ForceApiLogger)
	TraceOff()
}

type traceable struct {
	logger    ForceApiLogger
	logPrefix string
}

// TraceOn turns on logging for this Api. After this is called, all
// requests, responses, and raw response bodies will be sent to the logger.
// If prefix is a non-empty string, it will be written to the front of all
// logged strings, which can aid in filtering log lines.
//
// Use TraceOn if you want to spy on the Api requests and responses.
//
// Note that the base log.Logger type satisfies ForceApiLogger, but adapters
// can easily be written for other logging packages (e.g., the
// golang-sanctioned glog framework).
func (t *traceable) TraceOn(prefix string, logger ForceApiLogger) {
	t.logger = logger
	if prefix == "" {
		t.logPrefix = prefix
	} else {
		t.logPrefix = fmt.Sprintf("%s ", prefix)
	}
}

// TraceOff turns off tracing. It is idempotent.
func (t *traceable) TraceOff() {
	t.logger = nil
	t.logPrefix = ""
}

func (t *traceable) trace(name string, value interface{}, format string) {
	if t.logger != nil {
		logMsg := "%s%s " + format + "\n"
		t.logger.Printf(logMsg, t.logPrefix, name, value)
	}
}
