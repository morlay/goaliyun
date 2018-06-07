package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamStreamStatusRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamStreamStatusRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamStreamStatusResponse, error) {
	resp := &DescribeLiveStreamStreamStatusResponse{}
	err := client.Request("cdn", "DescribeLiveStreamStreamStatus", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamStreamStatusResponse struct {
	RequestId    goaliyun.String
	StreamStatus goaliyun.String
}
