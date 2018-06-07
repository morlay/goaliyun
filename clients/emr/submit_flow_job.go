package emr

import (
	"github.com/morlay/goaliyun"
)

type SubmitFlowJobRequest struct {
	JobId           string `position:"Query" name:"JobId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	Conf            string `position:"Query" name:"Conf"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SubmitFlowJobRequest) Invoke(client goaliyun.Client) (*SubmitFlowJobResponse, error) {
	resp := &SubmitFlowJobResponse{}
	err := client.Request("emr", "SubmitFlowJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitFlowJobResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
