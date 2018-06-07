package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type QueryTaskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryTaskRequest) Invoke(client goaliyun.Client) (*QueryTaskResponse, error) {
	resp := &QueryTaskResponse{}
	err := client.Request("r-kvstore", "QueryTask", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTaskResponse struct {
	RequestId goaliyun.String
	Action    goaliyun.String
	Progress  goaliyun.Integer
}
