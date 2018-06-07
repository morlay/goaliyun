package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportCensorJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportCensorJobResultRequest) Invoke(client goaliyun.Client) (*ReportCensorJobResultResponse, error) {
	resp := &ReportCensorJobResultResponse{}
	err := client.Request("mts", "ReportCensorJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportCensorJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
