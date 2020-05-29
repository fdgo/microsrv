package logger

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNsqSentryLogger(t *testing.T) {
	l, err := NewNsqSentryLogger(
		"localhost:4150",
		"write_test",
		1,
		"test",
		"https://xyz.20",
	)
	assert.Nil(t, err)

	m := map[string]interface{}{
		"request_id": "123-4232-1323-244",
		"appname":    "reply_provider",
	}
	d, _ := json.Marshal(m)
	l.Info(string(d))
	time.Sleep(1 * time.Second)
}
