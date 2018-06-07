package vod

import (
	"github.com/morlay/goaliyun"
)

type DescribeRefreshQuotaRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRefreshQuotaRequest) Invoke(client goaliyun.Client) (*DescribeRefreshQuotaResponse, error) {
	resp := &DescribeRefreshQuotaResponse{}
	err := client.Request("vod", "DescribeRefreshQuota", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRefreshQuotaResponse struct {
	RequestId         goaliyun.String
	RefreshCacheQuota DescribeRefreshQuotaRefreshCacheQuota
}

type DescribeRefreshQuotaRefreshCacheQuota struct {
	UrlQuota  goaliyun.String
	DirQuota  goaliyun.String
	UrlRemain goaliyun.String
	DirRemain goaliyun.String
}
