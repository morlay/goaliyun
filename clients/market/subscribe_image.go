package market

import (
	"github.com/morlay/goaliyun"
)

type SubscribeImageRequest struct {
	ProductCode string `position:"Query" name:"ProductCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SubscribeImageRequest) Invoke(client goaliyun.Client) (*SubscribeImageResponse, error) {
	resp := &SubscribeImageResponse{}
	err := client.Request("market", "SubscribeImage", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubscribeImageResponse struct {
	RequestId goaliyun.String
	Success   bool
}
