package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportCoverJobResultRequest struct {
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportCoverJobResultRequest) Invoke(client goaliyun.Client) (*ReportCoverJobResultResponse, error) {
	resp := &ReportCoverJobResultResponse{}
	err := client.Request("mts", "ReportCoverJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportCoverJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
