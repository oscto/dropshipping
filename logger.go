package dropshipping

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func (c *Client) SetLogger(logger io.Writer) {
	c.Logger = logger
}

func (c *Client) log(req *http.Request, rep *http.Response) {
	if c.Logger != nil {
		var (
			reqData string
			repData []byte
		)

		if req != nil {
			reqData = fmt.Sprintf("%s %s. Data: %s", req.Method, req.URL.String(), req.Form.Encode())
		}
		if rep != nil {
			repData, _ = httputil.DumpResponse(rep, true)
		}
		c.Logger.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqData, string(repData))))
	}
}
