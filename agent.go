package consulapi

// Agent can be used to query the Agent endpoints
type Agent struct {
	c *Client
}

// Agent returns a handle to the agent endpoints
func (c *Client) Agent() *Agent {
	return &Agent{c}
}

// Self is used to query the agent we are speaking to for
// information about itself
func (a *Agent) Self() (map[string]map[string]interface{}, error) {
	r := a.c.newRequest("GET", "/v1/agent/self")
	_, resp, err := requireOK(a.c.doRequest(r))
	if err != nil {
		return nil, err
	}

	var out map[string]map[string]interface{}
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}
