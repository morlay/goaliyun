package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateFlowJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RunConf         string `position:"Query" name:"RunConf"`
	EnvConf         string `position:"Query" name:"EnvConf"`
	Description     string `position:"Query" name:"Description"`
	Type            string `position:"Query" name:"Type"`
	Params          string `position:"Query" name:"Params"`
	ParamConf       string `position:"Query" name:"ParamConf"`
	FailAct         string `position:"Query" name:"FailAct"`
	Mode            string `position:"Query" name:"Mode"`
	RetryInterval   int64  `position:"Query" name:"RetryInterval"`
	Name            string `position:"Query" name:"Name"`
	MaxRetry        int64  `position:"Query" name:"MaxRetry"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentCategory  string `position:"Query" name:"ParentCategory"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateFlowJobRequest) Invoke(client goaliyun.Client) (*CreateFlowJobResponse, error) {
	resp := &CreateFlowJobResponse{}
	err := client.Request("emr", "CreateFlowJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowJobResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
