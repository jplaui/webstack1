package context

import "context"

type errorStateCtx string

var ErrorStateCtx errorStateCtx = "error"

func GetCtxErrorState(c context.Context) string {
	if errorString, ok := c.Value(ErrorStateCtx).(string); ok {
		return errorString
	}
	return ""
}

type successStateCtx string

var SuccessStateCtx successStateCtx = "success"

func GetCtxSuccessState(c context.Context) string {
	if successString, ok := c.Value(SuccessStateCtx).(string); ok {
		return successString
	}
	return ""
}

// returns notification message, error
func GetCtxNotification(c context.Context) (string, bool) {
	if successMessage, ok := c.Value(SuccessStateCtx).(string); ok {
		return successMessage, false
	} else {
		if errorMessage, ok := c.Value(ErrorStateCtx).(string); ok {
			return errorMessage, true
		} else {
			return "", false
		}
	}
}
