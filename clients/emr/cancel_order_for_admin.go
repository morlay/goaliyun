package emr

import (
	"github.com/morlay/goaliyun"
)

type CancelOrderForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CancelOrderForAdminRequest) Invoke(client goaliyun.Client) (*CancelOrderForAdminResponse, error) {
	resp := &CancelOrderForAdminResponse{}
	err := client.Request("emr", "CancelOrderForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelOrderForAdminResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}
