package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeCustomLogConfigRequest struct {
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCustomLogConfigRequest) Invoke(client goaliyun.Client) (*DescribeCustomLogConfigResponse, error) {
	resp := &DescribeCustomLogConfigResponse{}
	err := client.Request("cdn", "DescribeCustomLogConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCustomLogConfigResponse struct {
	RequestId goaliyun.String
	Remark    goaliyun.String
	Sample    goaliyun.String
	Tag       goaliyun.String
}
