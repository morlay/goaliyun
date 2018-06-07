package aegis

import (
	"github.com/morlay/goaliyun"
)

type DescribeWebshellRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	GroupId         int64  `position:"Query" name:"GroupId"`
	Remark          string `position:"Query" name:"Remark"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeWebshellRequest) Invoke(client goaliyun.Client) (*DescribeWebshellResponse, error) {
	resp := &DescribeWebshellResponse{}
	err := client.Request("aegis", "DescribeWebshell", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeWebshellResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}
