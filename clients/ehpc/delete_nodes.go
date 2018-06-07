package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteNodesRequest struct {
	ReleaseInstance string                   `position:"Query" name:"ReleaseInstance"`
	Instances       *DeleteNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId       string                   `position:"Query" name:"ClusterId"`
	RegionId        string                   `position:"Query" name:"RegionId"`
}

func (req *DeleteNodesRequest) Invoke(client goaliyun.Client) (*DeleteNodesResponse, error) {
	resp := &DeleteNodesResponse{}
	err := client.Request("ehpc", "DeleteNodes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNodesInstance struct {
	Id string `name:"Id"`
}

type DeleteNodesResponse struct {
	RequestId goaliyun.String
}

type DeleteNodesInstanceList []DeleteNodesInstance

func (list *DeleteNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteNodesInstance)
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
