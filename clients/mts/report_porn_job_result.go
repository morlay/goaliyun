package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportPornJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportPornJobResultRequest) Invoke(client goaliyun.Client) (*ReportPornJobResultResponse, error) {
	resp := &ReportPornJobResultResponse{}
	err := client.Request("mts", "ReportPornJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportPornJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
