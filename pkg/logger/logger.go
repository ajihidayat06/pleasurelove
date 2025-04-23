package logger

import (
	"context"
	"os"
	"pleasurelove/internal/constanta"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	if os.Getenv("ENV") == "production" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
			DisableQuote:  true,
		})
	}
}

func getTraceAndRequestID(ctx context.Context) (string, string) {
	traceID := ""
	requestID := ""

	if val := ctx.Value(constanta.TraceID); val != nil {
		traceID = val.(string)
	}
	if val := ctx.Value(constanta.RequestID); val != nil {
		requestID = val.(string)
	}

	return traceID, requestID
}

// Fungsi Info untuk logging dengan otomatis menambahkan trace_id dan request_id
func Info(ctx context.Context, message string, fields map[string]interface{}) {
	traceID, requestID := getTraceAndRequestID(ctx)

	// Menambahkan trace_id dan request_id ke dalam fields log
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["trace_id"] = traceID
	fields["request_id"] = requestID

	log.WithFields(fields).Info(message)
}

func Error(ctx context.Context, message string, err error) {
	LogWithCaller(ctx, message, err, 2)
}

// Fungsi Error untuk logging dengan otomatis menambahkan trace_id dan request_id
func LogWithCaller(ctx context.Context, message string, err error, callerSkip int) {
	traceID, requestID := getTraceAndRequestID(ctx)

	// Menambahkan trace_id dan request_id ke dalam fields log
	fields := logrus.Fields{
		"timestamp":  time.Now().Format(time.RFC3339),
		"error":      err,
		"trace_id":   traceID,
		"request_id": requestID,
	}

	// Menambahkan informasi tambahan terkait file dan fungsi jika diperlukan
	pc, file, line, _ := runtime.Caller(callerSkip)
	funcName := runtime.FuncForPC(pc).Name()
	fields["file"] = file
	fields["line"] = line
	fields["function"] = funcName

	log.WithFields(fields).Error(message)
}
