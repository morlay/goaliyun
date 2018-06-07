package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteResourcePoolRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ResourcePoolId  string `position:"Query" name:"ResourcePoolId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteResourcePoolRequest) Invoke(client goaliyun.Client) (*DeleteResourcePoolResponse, error) {
	resp := &DeleteResourcePoolResponse{}
	err := client.Request("emr", "DeleteResourcePool", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteResourcePoolResponse struct {
	RequestId goaliyun.String
}
