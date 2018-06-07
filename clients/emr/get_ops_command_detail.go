package emr

import (
	"github.com/morlay/goaliyun"
)

type GetOpsCommandDetailRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OpsCommandName  string `position:"Query" name:"OpsCommandName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetOpsCommandDetailRequest) Invoke(client goaliyun.Client) (*GetOpsCommandDetailResponse, error) {
	resp := &GetOpsCommandDetailResponse{}
	err := client.Request("emr", "GetOpsCommandDetail", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOpsCommandDetailResponse struct {
	RequestId   goaliyun.String
	Id          goaliyun.Integer
	Name        goaliyun.String
	Description goaliyun.String
	TargetType  goaliyun.String
	ServiceName goaliyun.String
	Category    goaliyun.String
	Params      goaliyun.String
}
