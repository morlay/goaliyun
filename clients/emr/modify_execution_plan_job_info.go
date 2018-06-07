package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyExecutionPlanJobInfoRequest struct {
	ResourceOwnerId int64                                    `position:"Query" name:"ResourceOwnerId"`
	Id              string                                   `position:"Query" name:"Id"`
	JobIdLists      *ModifyExecutionPlanJobInfoJobIdListList `position:"Query" type:"Repeated" name:"JobIdList"`
	RegionId        string                                   `position:"Query" name:"RegionId"`
}

func (req *ModifyExecutionPlanJobInfoRequest) Invoke(client goaliyun.Client) (*ModifyExecutionPlanJobInfoResponse, error) {
	resp := &ModifyExecutionPlanJobInfoResponse{}
	err := client.Request("emr", "ModifyExecutionPlanJobInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyExecutionPlanJobInfoResponse struct {
	RequestId goaliyun.String
}

type ModifyExecutionPlanJobInfoJobIdListList []string

func (list *ModifyExecutionPlanJobInfoJobIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
