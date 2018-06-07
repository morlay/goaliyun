package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyFlowJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RunConf         string `position:"Query" name:"RunConf"`
	EnvConf         string `position:"Query" name:"EnvConf"`
	Description     string `position:"Query" name:"Description"`
	Params          string `position:"Query" name:"Params"`
	ParamConf       string `position:"Query" name:"ParamConf"`
	FailAct         string `position:"Query" name:"FailAct"`
	Mode            string `position:"Query" name:"Mode"`
	RetryInterval   int64  `position:"Query" name:"RetryInterval"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	MaxRetry        int64  `position:"Query" name:"MaxRetry"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyFlowJobRequest) Invoke(client goaliyun.Client) (*ModifyFlowJobResponse, error) {
	resp := &ModifyFlowJobResponse{}
	err := client.Request("emr", "ModifyFlowJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFlowJobResponse struct {
	RequestId goaliyun.String
	Data      bool
}
