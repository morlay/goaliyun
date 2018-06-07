package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryClusterByUserRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryClusterByUserRequest) Invoke(client goaliyun.Client) (*QueryClusterByUserResponse, error) {
	resp := &QueryClusterByUserResponse{}
	err := client.Request("emr", "QueryClusterByUser", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryClusterByUserResponse struct {
	RequestId goaliyun.String
	Exist     bool
}
