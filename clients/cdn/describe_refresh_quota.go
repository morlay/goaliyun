package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeRefreshQuotaRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeRefreshQuotaRequest) Invoke(client goaliyun.Client) (*DescribeRefreshQuotaResponse, error) {
	resp := &DescribeRefreshQuotaResponse{}
	err := client.Request("cdn", "DescribeRefreshQuota", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRefreshQuotaResponse struct {
	RequestId     goaliyun.String
	UrlQuota      goaliyun.String
	DirQuota      goaliyun.String
	UrlRemain     goaliyun.String
	DirRemain     goaliyun.String
	PreloadQuota  goaliyun.String
	BlockQuota    goaliyun.String
	PreloadRemain goaliyun.String
	BlockRemain   goaliyun.String
}
