package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RunOpsCommandRequest struct {
	ResourceOwnerId int64                        `position:"Query" name:"ResourceOwnerId"`
	OpsCommandName  string                       `position:"Query" name:"OpsCommandName"`
	Comment         string                       `position:"Query" name:"Comment"`
	CustomParams    string                       `position:"Query" name:"CustomParams"`
	ClusterId       string                       `position:"Query" name:"ClusterId"`
	HostIdLists     *RunOpsCommandHostIdListList `position:"Query" type:"Repeated" name:"HostIdList"`
	Dimension       string                       `position:"Query" name:"Dimension"`
	RegionId        string                       `position:"Query" name:"RegionId"`
}

func (req *RunOpsCommandRequest) Invoke(client goaliyun.Client) (*RunOpsCommandResponse, error) {
	resp := &RunOpsCommandResponse{}
	err := client.Request("emr", "RunOpsCommand", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RunOpsCommandResponse struct {
	RequestId   goaliyun.String
	OperationId goaliyun.Integer
}

type RunOpsCommandHostIdListList []int64

func (list *RunOpsCommandHostIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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
