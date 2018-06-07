package emr

import (
	"github.com/morlay/goaliyun"
)

type KillFlowJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *KillFlowJobRequest) Invoke(client goaliyun.Client) (*KillFlowJobResponse, error) {
	resp := &KillFlowJobResponse{}
	err := client.Request("emr", "KillFlowJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type KillFlowJobResponse struct {
	RequestId goaliyun.String
	Data      bool
}
