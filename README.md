consul-api
==========

This package provides the `consulapi` package which attempts to
provide programmatic access to the full Consul API. Eventually,
all commands will be supported.

Documentation
=============

The full documentation is available on [Godoc](http://godoc.org/github.com/armon/consul-api)

Status
======

The current status of the API:
* kv - Complete
* session - Complete
* status - Complete
* agent - Complete
* health - Complete
* catalog - All query endpoints

Usage
=====

Below is an example of using the Consul client:

```go
// Get a new client, with KV endpoints
client, _ := consulapi.NewClient(consulapi.DefaultConfig())
kv := client.KV()

// PUT a new KV pair
p := &consulapi.KVPair{Key: "foo", Value: []byte("test")}
_, err := kv.Put(p, nil)
if err != nil {
    panic(err)
}

// Lookup the pair
pair, _, err := kv.Get("foo", nil)
if err != nil {
    panic(err)
}
fmt.Printf("KV: %v", pair)

```

