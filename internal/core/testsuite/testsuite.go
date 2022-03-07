package testsuite

import (
	"bytes"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pedrocmart/canvas/internal/core/errors"
)

type RetrySpecs struct {
	Failed        bool
	RetryRequired bool
	ISessionID    string
	NextRetryTime time.Time
}

type MockReader int

func (MockReader) Read(p []byte) (n int, err error) {
	return 0, errors.Wrap("test error")
}

type MockResponseWriter struct {
	Code        int
	Flushed     bool
	wroteHeader bool
	Body        *bytes.Buffer
	// result      *http.Response
	HeaderMap  http.Header
	snapHeader http.Header
}

func (mrw MockResponseWriter) Write(b []byte) (n int, err error) {
	return 0, errors.Wrap("mock writer error")
}

func (mrw *MockResponseWriter) Header() http.Header {
	m := mrw.HeaderMap
	if m == nil {
		m = make(http.Header)
		mrw.HeaderMap = m
	}

	return m
}

func (mrw *MockResponseWriter) WriteHeader(code int) {
	if mrw.wroteHeader {
		return
	}

	mrw.Code = code

	mrw.wroteHeader = true
	if mrw.HeaderMap == nil {
		mrw.HeaderMap = make(http.Header)
	}

	mrw.snapHeader = cloneHeader(mrw.HeaderMap)
}

func cloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))

	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}

	return h2
}

func GenUUID() uuid.UUID {
	return uuid.MustParse("a1c790af-48bd-4081-9086-604f6564303e")
}

func GenTime() time.Time {
	tt, _ := time.Parse("2006-01-02T15:04:05", "2020-01-02T03:04:05")

	return tt
}

func StringPointer(x string) *string {
	return &x
}

func StringToTime(input string) time.Time {
	tt, _ := time.Parse("2006-01-02T15:04:05", input)

	return tt
}

func BoolPointer(input bool) *bool {
	return &input
}

func TimePointer(input time.Time) *time.Time {
	return &input
}

func RetrySpec(sessionID string, nextRetry time.Time, retryRequired bool) *RetrySpecs {
	return &RetrySpecs{
		Failed:        true,
		RetryRequired: retryRequired,
		ISessionID:    sessionID,
		NextRetryTime: nextRetry,
	}
}
