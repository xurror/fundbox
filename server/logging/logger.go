package logging

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
	"os"
	"strings"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)

	return logger
}

type AppLogger struct {
	Logger *logrus.Logger
}

var _ fxevent.Logger = (*AppLogger)(nil)

func (l *AppLogger) logEvent(msg string, fields logrus.Fields) {
	l.Logger.WithFields(fields).Log(logrus.InfoLevel, msg)
}

func (l *AppLogger) logError(msg string, fields logrus.Fields) {
	l.Logger.WithFields(fields).Log(logrus.ErrorLevel, msg)
}

func (l *AppLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logEvent("OnStart hook executing",
			logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			})
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.logError("OnStart hook failed",
				logrus.Fields{
					"callee": e.FunctionName,
					"caller": e.CallerName,
					"error":  e.Err.Error(),
				})
		} else {
			l.logEvent("OnStart hook executed",
				logrus.Fields{
					"callee":  e.FunctionName,
					"caller":  e.CallerName,
					"runtime": e.Runtime.String(),
				})
		}
	case *fxevent.OnStopExecuting:
		l.logEvent("OnStop hook executing",
			logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			})
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.logError("OnStop hook failed",
				logrus.Fields{
					"callee": e.FunctionName,
					"caller": e.CallerName,
					"error":  e.Err.Error(),
				})
		} else {
			l.logEvent("OnStop hook executed",
				logrus.Fields{
					"callee":  e.FunctionName,
					"caller":  e.CallerName,
					"runtime": e.Runtime.String(),
				})
		}
	case *fxevent.Supplied:
		mf, mv := moduleField(e.ModuleName)
		if e.Err != nil {
			l.logError("error encountered while applying options",
				logrus.Fields{
					"type":       e.TypeName,
					"stacktrace": e.StackTrace,
					mf:           mv,
					"error":      e.Err.Error(),
				})
		} else {
			l.logEvent("supplied",
				logrus.Fields{
					"type":       e.TypeName,
					"stacktrace": e.StackTrace,
					mf:           mv,
				})
		}
	case *fxevent.Provided:
		mf, mfv := moduleField(e.ModuleName)
		mb, mbv := maybeBool("private", e.Private)
		for _, rtype := range e.OutputTypeNames {
			l.logEvent("provided",
				logrus.Fields{
					"constructor": e.ConstructorName,
					//"stacktrace":  e.StackTrace,
					mf:     mfv,
					"type": rtype,
					mb:     mbv,
				})
		}
		if e.Err != nil {
			l.logError("error encountered while applying options",
				logrus.Fields{
					mf:           mfv,
					"stacktrace": e.StackTrace,
					"error":      e.Err.Error(),
				})
		}
	case *fxevent.Replaced:
		mf, mv := moduleField(e.ModuleName)
		for _, rtype := range e.OutputTypeNames {
			l.logEvent("replaced",
				logrus.Fields{
					"stacktrace": e.StackTrace,
					mf:           mv,
					"type":       rtype,
				})
		}
		if e.Err != nil {
			l.logError("error encountered while replacing",
				logrus.Fields{
					"stacktrace": e.StackTrace,
					mf:           mv,
					"error":      e.Err.Error(),
				})
		}
	case *fxevent.Decorated:
		mf, mv := moduleField(e.ModuleName)
		for _, rtype := range e.OutputTypeNames {
			l.logEvent("decorated",
				logrus.Fields{
					"decorator":  e.DecoratorName,
					"stacktrace": e.StackTrace,
					mf:           mv,
					"type":       rtype,
				})
		}
		if e.Err != nil {
			l.logError("error encountered while applying options",
				logrus.Fields{
					"stacktrace": e.StackTrace,
					mf:           mv,
					"error":      e.Err.Error(),
				})
		}
	case *fxevent.Run:
		mf, mv := moduleField(e.ModuleName)
		if e.Err != nil {
			l.logError("error returned",
				logrus.Fields{
					"name":  e.Name,
					"kind":  e.Kind,
					mf:      mv,
					"error": e.Err.Error(),
				})
		} else {
			l.logEvent("run",
				logrus.Fields{
					"name": e.Name,
					"kind": e.Kind,
					mf:     mv,
				})
		}
	case *fxevent.Invoking:
		mf, mv := moduleField(e.ModuleName)
		// Do not log stack as it will make logs hard to read.
		l.logEvent("invoking",
			logrus.Fields{
				"function": e.FunctionName,
				mf:         mv,
			})
	case *fxevent.Invoked:
		mf, mv := moduleField(e.ModuleName)
		if e.Err != nil {
			l.logError("invoke failed",
				logrus.Fields{
					"error":    e.Err.Error(),
					"stack":    e.Trace,
					"function": e.FunctionName,
					mf:         mv,
				})
		}
	case *fxevent.Stopping:
		l.logEvent("received signal",
			logrus.Fields{
				"signal": strings.ToUpper(e.Signal.String()),
			})
	case *fxevent.Stopped:
		if e.Err != nil {
			l.logError("stop failed", logrus.Fields{"error": e.Err.Error()})
		}
	case *fxevent.RollingBack:
		l.logError("start failed, rolling back", logrus.Fields{"error": e.StartErr.Error()})
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.logError("rollback failed", logrus.Fields{"error": e.Err.Error()})
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.logError("start failed", logrus.Fields{"error": e.Err.Error()})
		} else {
			l.logEvent("started", logrus.Fields{})
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.logError("custom logger initialization failed", logrus.Fields{"error": e.Err.Error()})
		} else {
			l.logEvent("initialized custom fxevent.Logger", logrus.Fields{"function": e.ConstructorName})
		}
	}
}

func (*AppLogger) String() string { return "AppLogger" }

func moduleField(name string) (string, interface{}) {
	if len(name) == 0 {
		return "", ""
	}
	return "module", name
}

func maybeBool(name string, b bool) (string, interface{}) {
	if b {
		return name, true
	}
	return "", ""
}
