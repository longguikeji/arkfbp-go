package node

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	APIModeDirect = "direct"
	APIModeProxy  = "proxy"
)

// APINode ...
type APINode struct {
	Node

	Mode    string
	URL     string
	Method  string
	Auth    interface{}
	Headers map[string]string
	Params  map[string]interface{}
}

// Name ...
func (n *APINode) Name() string {
	return "API"
}

// Kind ...
func (n *APINode) Kind() Kind {
	return KAPI
}

// Run ...
func (n *APINode) Run() interface{} {

	fmt.Printf("Mode: %s\n", n.Mode)

	switch n.Mode {
	case APIModeDirect:
		return n.requestDirectly()

	case APIModeProxy:
		return nil
	}

	return nil
}

func (n *APINode) requestDirectly() interface{} {
	client := &http.Client{}

	request, err := http.NewRequest(n.Method, n.URL, nil)
	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)

	n.Response().Status = response.StatusCode
	n.Response().StatusText = response.Status
	n.Response().Headers = response.Header
	n.Response().Data, _ = ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&n.Response().Data)
	_ = err

	return response.Body
}
