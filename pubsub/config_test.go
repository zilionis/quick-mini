package pubsub

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleConfig(t *testing.T) {

	config := HandleConfig()
	flag.CommandLine.Set("name", "Example name")

	assert.Equal(t, "Example name", *config.Name)
	assert.Equal(t, 6001, *config.PortSub)
}
