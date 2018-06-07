package dcdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeDcdnRefreshQuotaRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnRefreshQuotaRequest) Invoke(client goaliyun.Client) (*DescribeDcdnRefreshQuotaResponse, error) {
	resp := &DescribeDcdnRefreshQuotaResponse{}
	err := client.Request("dcdn", "DescribeDcdnRefreshQuota", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnRefreshQuotaResponse struct {
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
