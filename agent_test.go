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

func TestAgent_Services(t *testing.T) {
	c := makeClient(t)
	agent := c.Agent()

	reg := &AgentServiceRegistration{
		Name: "foo",
		Tags: []string{"bar", "baz"},
		Port: 8000,
		Check: &AgentServiceCheck{
			TTL: "15s",
		},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		t.Fatalf("err: %v", err)
	}

	services, err := agent.Services()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if _, ok := services["foo"]; !ok {
		t.Fatalf("missing service: %v", services)
	}

	checks, err := agent.Checks()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if _, ok := checks["service:foo"]; !ok {
		t.Fatalf("missing check: %v", checks)
	}

	if err := agent.ServiceDeregister("foo"); err != nil {
		t.Fatalf("err: %v", err)
	}
}
