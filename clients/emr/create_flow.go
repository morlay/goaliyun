package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateFlowRequest struct {
	CronExpr       string `position:"Query" name:"CronExpr"`
	StartSchedule  int64  `position:"Query" name:"StartSchedule"`
	Description    string `position:"Query" name:"Description"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	Type           string `position:"Query" name:"Type"`
	Graph          string `position:"Query" name:"Graph"`
	CreateCluster  string `position:"Query" name:"CreateCluster"`
	Name           string `position:"Query" name:"Name"`
	EndSchedule    int64  `position:"Query" name:"EndSchedule"`
	ProjectId      string `position:"Query" name:"ProjectId"`
	ParentCategory string `position:"Query" name:"ParentCategory"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreateFlowRequest) Invoke(client goaliyun.Client) (*CreateFlowResponse, error) {
	resp := &CreateFlowResponse{}
	err := client.Request("emr", "CreateFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
