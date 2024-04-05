package monkeylsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	//just specify the type of the params in all of the request types later.
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
	//result | error
	//each type define seperately
}

type Notification struct {
	RPD    string `json:"jsonrpc"`
	Method string `json:"method"`
}
