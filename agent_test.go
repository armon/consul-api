package consulapi

import (
	"testing"
)

func TestAgent_Self(t *testing.T) {
	c := makeClient(t)
	agent := c.Agent()

	info, err := agent.Self()
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	name := info["Config"]["NodeName"]
	if name == "" {
		t.Fatalf("bad: %v", info)
	}
}
