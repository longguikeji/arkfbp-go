package request

// Request ...
type Request struct {
	Schema   string
	Hostname string
	Method   string
	Path     string
	Headers  map[string][]string
	Cookies  map[string]string
	Params   map[string]interface{}
	Body     []byte
}
