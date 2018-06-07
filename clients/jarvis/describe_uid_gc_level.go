package jarvis

import (
	"github.com/morlay/goaliyun"
)

type DescribeUidGcLevelRequest struct {
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	SourceCode string `position:"Query" name:"SourceCode"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeUidGcLevelRequest) Invoke(client goaliyun.Client) (*DescribeUidGcLevelResponse, error) {
	resp := &DescribeUidGcLevelResponse{}
	err := client.Request("jarvis", "DescribeUidGcLevel", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUidGcLevelResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	Gclevel   goaliyun.String
}
