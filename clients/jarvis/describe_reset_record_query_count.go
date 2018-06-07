package jarvis

import (
	"github.com/morlay/goaliyun"
)

type DescribeResetRecordQueryCountRequest struct {
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	SourceCode string `position:"Query" name:"SourceCode"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeResetRecordQueryCountRequest) Invoke(client goaliyun.Client) (*DescribeResetRecordQueryCountResponse, error) {
	resp := &DescribeResetRecordQueryCountResponse{}
	err := client.Request("jarvis", "DescribeResetRecordQueryCount", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeResetRecordQueryCountResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	QueryCount goaliyun.Integer
	Module     goaliyun.String
}
