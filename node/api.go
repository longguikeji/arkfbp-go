package node

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	APIModeDirect = "direct"
	APIModeProxy  = "proxy"
)

type APIResponse struct {
}

// APINode ...
type APINode struct {
	Node

	Mode    string
	URL     string
	Method  string
	Auth    interface{}
	Headers map[string]string
	Params  map[string]interface{}

	resp       APIResponse
	status     int
	statusText string
	header     map[string][]string
	data       interface{}
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

// Status ...
func (n *APINode) Status() int {
	return n.status
}

// StatusText ...
func (n *APINode) StatusText() string {
	return n.statusText
}

func (n *APINode) requestDirectly() interface{} {
	client := &http.Client{}

	fmt.Println(n.Method, n.URL)

	request, err := http.NewRequest(n.Method, n.URL, nil)
	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)

	n.status = response.StatusCode
	n.statusText = response.Status
	n.header = response.Header

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&n.data)
	_ = err

	fmt.Println(response)

	return response.Body
}
