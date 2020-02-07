package response

import "encoding/json"

import "bytes"

// Response ...
type Response struct {
	Status     int
	StatusText string
	Headers    map[string][]string
	Data       []byte
}

// JSON ...
func (r *Response) JSON() (interface{}, error) {
	var ret interface{}
	reader := bytes.NewReader(r.Data)
	if err := json.NewDecoder(reader).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
