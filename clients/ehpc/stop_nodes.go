package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StopNodesRequest struct {
	Role      string                 `position:"Query" name:"Role"`
	Instances *StopNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId string                 `position:"Query" name:"ClusterId"`
	RegionId  string                 `position:"Query" name:"RegionId"`
}

func (req *StopNodesRequest) Invoke(client goaliyun.Client) (*StopNodesResponse, error) {
	resp := &StopNodesResponse{}
	err := client.Request("ehpc", "StopNodes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopNodesInstance struct {
	Id string `name:"Id"`
}

type StopNodesResponse struct {
	RequestId goaliyun.String
}

type StopNodesInstanceList []StopNodesInstance

func (list *StopNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StopNodesInstance)
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
