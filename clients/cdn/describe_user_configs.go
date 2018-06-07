package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeUserConfigsRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Config        string `position:"Query" name:"Config"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserConfigsRequest) Invoke(client goaliyun.Client) (*DescribeUserConfigsResponse, error) {
	resp := &DescribeUserConfigsResponse{}
	err := client.Request("cdn", "DescribeUserConfigs", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserConfigsResponse struct {
	RequestId goaliyun.String
	Configs   DescribeUserConfigsConfigs
}

type DescribeUserConfigsConfigs struct {
	OssLogConfig       DescribeUserConfigsOssLogConfig
	GreenManagerConfig DescribeUserConfigsGreenManagerConfig
}

type DescribeUserConfigsOssLogConfig struct {
	Enable goaliyun.String
	Bucket goaliyun.String
	Prefix goaliyun.String
}

type DescribeUserConfigsGreenManagerConfig struct {
	Quota goaliyun.String
	Ratio goaliyun.String
}
