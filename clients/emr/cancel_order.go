package emr

import (
	"github.com/morlay/goaliyun"
)

type CancelOrderRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CancelOrderRequest) Invoke(client goaliyun.Client) (*CancelOrderResponse, error) {
	resp := &CancelOrderResponse{}
	err := client.Request("emr", "CancelOrder", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelOrderResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}
