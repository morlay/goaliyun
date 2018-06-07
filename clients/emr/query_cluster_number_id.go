package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryClusterNumberIdRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryClusterNumberIdRequest) Invoke(client goaliyun.Client) (*QueryClusterNumberIdResponse, error) {
	resp := &QueryClusterNumberIdResponse{}
	err := client.Request("emr", "QueryClusterNumberId", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryClusterNumberIdResponse struct {
	RequestId       goaliyun.String
	ClusterNumberId goaliyun.String
}
