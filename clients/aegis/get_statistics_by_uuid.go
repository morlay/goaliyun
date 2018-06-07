package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetStatisticsByUuidRequest struct {
	Uuid     string `position:"Query" name:"Uuid"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetStatisticsByUuidRequest) Invoke(client goaliyun.Client) (*GetStatisticsByUuidResponse, error) {
	resp := &GetStatisticsByUuidResponse{}
	err := client.Request("aegis", "GetStatisticsByUuid", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetStatisticsByUuidResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      GetStatisticsByUuidEntityList
}

type GetStatisticsByUuidEntity struct {
	Uuid    goaliyun.String
	Account goaliyun.Integer
	Health  goaliyun.Integer
	Patch   goaliyun.Integer
	Trojan  goaliyun.Integer
	Online  bool
}

type GetStatisticsByUuidEntityList []GetStatisticsByUuidEntity

func (list *GetStatisticsByUuidEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetStatisticsByUuidEntity)
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
