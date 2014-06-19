package consulapi

import (
	"testing"
)

func TestCatalog_Datacenters(t *testing.T) {
	c := makeClient(t)
	catalog := c.Catalog()

	datacenters, err := catalog.Datacenters()
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(datacenters) == 0 {
		t.Fatalf("Bad: %v", datacenters)
	}
}

func TestCatalog_Nodes(t *testing.T) {
	c := makeClient(t)
	catalog := c.Catalog()

	nodes, meta, err := catalog.Nodes(nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if meta.LastIndex == 0 {
		t.Fatalf("Bad: %v", meta)
	}

	if len(nodes) == 0 {
		t.Fatalf("Bad: %v", nodes)
	}
}

func TestCatalog_Services(t *testing.T) {
	c := makeClient(t)
	catalog := c.Catalog()

	services, meta, err := catalog.Services(nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if meta.LastIndex == 0 {
		t.Fatalf("Bad: %v", meta)
	}

	if len(services) == 0 {
		t.Fatalf("Bad: %v", services)
	}
}

func TestCatalog_Service(t *testing.T) {
	c := makeClient(t)
	catalog := c.Catalog()

	services, meta, err := catalog.Service("consul", "", nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if meta.LastIndex == 0 {
		t.Fatalf("Bad: %v", meta)
	}

	if len(services) == 0 {
		t.Fatalf("Bad: %v", services)
	}
}

func TestCatalog_Node(t *testing.T) {
	c := makeClient(t)
	catalog := c.Catalog()

	name, _ := c.Agent().NodeName()
	info, meta, err := catalog.Node(name, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if meta.LastIndex == 0 {
		t.Fatalf("Bad: %v", meta)
	}
	if len(info.Services) == 0 {
		t.Fatalf("Bad: %v", info)
	}
}
