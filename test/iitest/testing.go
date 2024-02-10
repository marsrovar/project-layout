package iitest

import (
	"iserver/server"
	"iserver/server/protocol"
	"iserver/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func init() {
	gin.SetMode(gin.TestMode)

	// connect mockdb redis
	// connect mockdb db

	go server.StartServer(":9999")
}

func TestingDoF(tF func(), preDo ...func() error) {
	TruncateAllTable()
	FlushRedis()
	defer TruncateAllTable()
	defer FlushRedis()

	if len(preDo) != 0 && preDo != nil {
		preDo[0]()
	}

	tF()
}

func TestingNewAccountDo(tF func()) {
	TestingDoF(tF, func() error {
		if err := insertAccount(); err != nil {
			panic(err)
		}

		return nil
	})
}

func TruncateAllTable() {

}

func FlushRedis() {

}

type Client struct {
	Host string
}

func NewClient(c *Client) *resty.Client {
	client := resty.New()
	client.SetBaseURL(c.Host)
	client.JSONMarshal = utils.Json.Marshal
	client.JSONUnmarshal = utils.Json.Unmarshal

	return client
}

type IClient struct {
	Body        interface{}
	Header      map[string]string
	QueryParams map[string]string
}

func NewIClient() *IClient {
	return &IClient{}
}

func (ic *IClient) SetHeader(header map[string]string) *IClient {
	ic.Header = header
	return ic
}

func (ic *IClient) SetBody(body interface{}) *IClient {
	ic.Body = body
	return ic
}

func (ic *IClient) SetQueryParams(params map[string]string) *IClient {
	ic.QueryParams = params
	return ic
}

func (ic *IClient) Do(result interface{}, method, url string) (int, int, error) {
	nReq := NewClient(&Client{
		Host: "http://localhost:9999",
	}).NewRequest()

	r := &protocol.RespData{
		Data: result,
	}

	if len(ic.Header) != 0 {
		nReq.SetHeaders(ic.Header)
	}

	nReq.SetResult(r).SetError(r)

	var resp *resty.Response
	var err error
	switch method {
	case "GET":
		if len(ic.QueryParams) != 0 {
			nReq.SetQueryParams(ic.QueryParams)
		}

		resp, err = nReq.Get(url)
	case "POST", "PUT":
		if ic.Body != nil {
			nReq.SetBody(ic.Body)
		}

		if method == "POST" {
			resp, err = nReq.Post(url)
		} else {
			resp, err = nReq.Put(url)
		}
	default:
		panic("method not found: " + method)
	}

	if err != nil {
		return 0, 0, err
	}

	return r.Code, resp.StatusCode(), nil
}
