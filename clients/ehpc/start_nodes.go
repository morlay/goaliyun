package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StartNodesRequest struct {
	Role      string                  `position:"Query" name:"Role"`
	Instances *StartNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId string                  `position:"Query" name:"ClusterId"`
	RegionId  string                  `position:"Query" name:"RegionId"`
}

func (req *StartNodesRequest) Invoke(client goaliyun.Client) (*StartNodesResponse, error) {
	resp := &StartNodesResponse{}
	err := client.Request("ehpc", "StartNodes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartNodesInstance struct {
	Id string `name:"Id"`
}

type StartNodesResponse struct {
	RequestId goaliyun.String
}

type StartNodesInstanceList []StartNodesInstance

func (list *StartNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartNodesInstance)
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
