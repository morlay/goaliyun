package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateClusterForInternalRequest struct {
	ResourceOwnerId   int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterTemplateId int64  `position:"Query" name:"ClusterTemplateId"`
	ClusterCondition  string `position:"Query" name:"ClusterCondition"`
	UserId            string `position:"Query" name:"UserId"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *CreateClusterForInternalRequest) Invoke(client goaliyun.Client) (*CreateClusterForInternalResponse, error) {
	resp := &CreateClusterForInternalResponse{}
	err := client.Request("emr", "CreateClusterForInternal", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateClusterForInternalResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}
