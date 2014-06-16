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

func TestAgent_Members(t *testing.T) {
	c := makeClient(t)
	agent := c.Agent()

	members, err := agent.Members(false)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(members) != 1 {
		t.Fatalf("bad: %v", members)
	}
}
