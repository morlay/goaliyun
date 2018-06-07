package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeIpInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	IP            string `position:"Query" name:"IP"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeIpInfoRequest) Invoke(client goaliyun.Client) (*DescribeIpInfoResponse, error) {
	resp := &DescribeIpInfoResponse{}
	err := client.Request("cdn", "DescribeIpInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeIpInfoResponse struct {
	RequestId   goaliyun.String
	CdnIp       goaliyun.String
	ISP         goaliyun.String
	IspEname    goaliyun.String
	Region      goaliyun.String
	RegionEname goaliyun.String
}
