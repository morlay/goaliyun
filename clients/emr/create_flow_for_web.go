package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateFlowForWebRequest struct {
	CronExpr       string `position:"Query" name:"CronExpr"`
	StartSchedule  int64  `position:"Query" name:"StartSchedule"`
	Name           string `position:"Query" name:"Name"`
	Description    string `position:"Query" name:"Description"`
	EndSchedule    int64  `position:"Query" name:"EndSchedule"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	ProjectId      string `position:"Query" name:"ProjectId"`
	Graph          string `position:"Query" name:"Graph"`
	ParentCategory string `position:"Query" name:"ParentCategory"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreateFlowForWebRequest) Invoke(client goaliyun.Client) (*CreateFlowForWebResponse, error) {
	resp := &CreateFlowForWebResponse{}
	err := client.Request("emr", "CreateFlowForWeb", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowForWebResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
