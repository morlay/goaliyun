package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type GetQuotaRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetQuotaRequest) Invoke(client goaliyun.Client) (*GetQuotaResponse, error) {
	resp := &GetQuotaResponse{}
	err := client.Request("cloudphoto", "GetQuota", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetQuotaResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Quota     GetQuotaQuota
}

type GetQuotaQuota struct {
	TotalQuota  goaliyun.Integer
	FacesCount  goaliyun.Integer
	PhotosCount goaliyun.Integer
	UsedQuota   goaliyun.Integer
	VideosCount goaliyun.Integer
}
