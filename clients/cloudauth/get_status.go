package cloudauth

import (
	"github.com/morlay/goaliyun"
)

type GetStatusRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	TicketId        string `position:"Query" name:"TicketId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetStatusRequest) Invoke(client goaliyun.Client) (*GetStatusResponse, error) {
	resp := &GetStatusResponse{}
	err := client.Request("cloudauth", "GetStatus", "2018-05-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetStatusData
}

type GetStatusData struct {
	StatusCode       goaliyun.Integer
	TrustedScore     goaliyun.Float
	SimilarityScore  goaliyun.Float
	AuditConclusions goaliyun.String
}
