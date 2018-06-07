package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportTagJobResultRequest struct {
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportTagJobResultRequest) Invoke(client goaliyun.Client) (*ReportTagJobResultResponse, error) {
	resp := &ReportTagJobResultResponse{}
	err := client.Request("mts", "ReportTagJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportTagJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
