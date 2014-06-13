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
* agent - Partial
* catalog - Not started
* health - Not started

Usage
=====

Below is an example of using the Consul client:

```go
client, _ := consulapi.NewClient(consulapi.DefaultConfig())

...
```

