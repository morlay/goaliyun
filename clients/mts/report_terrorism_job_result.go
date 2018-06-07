package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportTerrorismJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportTerrorismJobResultRequest) Invoke(client goaliyun.Client) (*ReportTerrorismJobResultResponse, error) {
	resp := &ReportTerrorismJobResultResponse{}
	err := client.Request("mts", "ReportTerrorismJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportTerrorismJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
