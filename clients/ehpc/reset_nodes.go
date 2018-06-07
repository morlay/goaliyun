package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ResetNodesRequest struct {
	Instances *ResetNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId string                  `position:"Query" name:"ClusterId"`
	RegionId  string                  `position:"Query" name:"RegionId"`
}

func (req *ResetNodesRequest) Invoke(client goaliyun.Client) (*ResetNodesResponse, error) {
	resp := &ResetNodesResponse{}
	err := client.Request("ehpc", "ResetNodes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetNodesInstance struct {
	Id string `name:"Id"`
}

type ResetNodesResponse struct {
	RequestId goaliyun.String
}

type ResetNodesInstanceList []ResetNodesInstance

func (list *ResetNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResetNodesInstance)
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
