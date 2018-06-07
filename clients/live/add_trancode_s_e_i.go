package live

import (
	"github.com/morlay/goaliyun"
)

type AddTrancodeSEIRequest struct {
	Delay      int64  `position:"Query" name:"Delay"`
	AppName    string `position:"Query" name:"AppName"`
	Repeat     int64  `position:"Query" name:"Repeat"`
	DomainName string `position:"Query" name:"DomainName"`
	Pattern    string `position:"Query" name:"Pattern"`
	Text       string `position:"Query" name:"Text"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	StreamName string `position:"Query" name:"StreamName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AddTrancodeSEIRequest) Invoke(client goaliyun.Client) (*AddTrancodeSEIResponse, error) {
	resp := &AddTrancodeSEIResponse{}
	err := client.Request("live", "AddTrancodeSEI", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddTrancodeSEIResponse struct {
	RequestId goaliyun.String
}
