package consulapi

import (
	"os"
	"testing"
)

// ROOT is a management token for the tests
var CONSUL_ROOT string

func init() {
	CONSUL_ROOT = os.Getenv("CONSUL_ROOT")
}

func TestACL_List(t *testing.T) {
	if CONSUL_ROOT == "" {
		t.SkipNow()
	}
	c := makeClient(t)
	c.config.Token = CONSUL_ROOT
	acl := c.ACL()

	acls, qm, err := acl.List(nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(acls) < 2 {
		t.Fatalf("bad: %v", acls)
	}

	if qm.LastIndex == 0 {
		t.Fatalf("bad: %v", qm)
	}
	if !qm.KnownLeader {
		t.Fatalf("bad: %v", qm)
	}

}
