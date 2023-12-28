package pubsub

import (
	"bou.ke/monkey"
	"bytes"
	"github.com/TwiN/go-color"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetAddressFromPort(t *testing.T) {
	testCases := []struct {
		portNumber int
		expected   string
	}{
		{6001, "0.0.0.0:6001"},
		{80, "0.0.0.0:80"},
	}

	for _, tc := range testCases {
		result := GetAddressFromPort(&tc.portNumber)
		assert.Equal(t, tc.expected, result)
	}
}

func TestLogger(t *testing.T) {
	var buf bytes.Buffer

	logger := Logger("Demo", color.BlackBackground)
	logger.SetOutput(&buf)

	timeMock := time.Date(2023, time.December, 1, 10, 11, 12, 4, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return timeMock })
	defer patch.Unpatch()

	logger.Print("foo")

	defer func() {
		log.SetOutput(os.Stdout)
	}()

	assert.Equal(t, "\x1b[1m\x1b[40m\x1b[97m[ Demo ]\x1b[0m 10:11:12 foo\n", buf.String(), "Output is not the same")
}
