package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type SetQuotaRequest struct {
	TotalQuota int64  `position:"Query" name:"TotalQuota"`
	LibraryId  string `position:"Query" name:"LibraryId"`
	StoreName  string `position:"Query" name:"StoreName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *SetQuotaRequest) Invoke(client goaliyun.Client) (*SetQuotaResponse, error) {
	resp := &SetQuotaResponse{}
	err := client.Request("cloudphoto", "SetQuota", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetQuotaResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}
