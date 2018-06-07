package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeFlowJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowJobRequest) Invoke(client goaliyun.Client) (*DescribeFlowJobResponse, error) {
	resp := &DescribeFlowJobResponse{}
	err := client.Request("emr", "DescribeFlowJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowJobResponse struct {
	RequestId     goaliyun.String
	Id            goaliyun.String
	GmtCreate     goaliyun.Integer
	GmtModified   goaliyun.Integer
	Name          goaliyun.String
	Description   goaliyun.String
	FailAct       goaliyun.String
	MaxRetry      goaliyun.Integer
	RetryInterval goaliyun.Integer
	Params        goaliyun.String
	ParamConf     goaliyun.String
	EnvConf       goaliyun.String
	RunConf       goaliyun.String
	CategoryId    goaliyun.String
	Mode          goaliyun.String
	Resource      DescribeFlowJobResourceList
}

type DescribeFlowJobResourceList []goaliyun.String

func (list *DescribeFlowJobResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
