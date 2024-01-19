package logging

const (
	CodeDebug = 10000
	CodeInfo  = 20000
	CodeWarn  = 30000
	CodeError = 40000

	CodeInfoAppStarted = 20100

	CodeWarmAttemptConnection = 30100

	CodeErrorAppStartFailed        = 40100
	CodeErrorReadConfigFailed      = 40200
	CodeErrorUnmarshalConfigFailed = 40300
	CodeErrorDBConnectionFailed    = 40400
	CodeErrorDBClosingFailed       = 40300

	CodeErrorGrpcConnectionCloseFailed = 40400
)
