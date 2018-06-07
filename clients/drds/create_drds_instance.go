package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateDrdsInstanceRequest struct {
	Quantity       int64  `position:"Query" name:"Quantity"`
	Description    string `position:"Query" name:"Description"`
	Specification  string `position:"Query" name:"Specification"`
	Type           string `position:"Query" name:"Type"`
	VswitchId      string `position:"Query" name:"VswitchId"`
	IsHa           string `position:"Query" name:"IsHa"`
	InstanceSeries string `position:"Query" name:"InstanceSeries"`
	VpcId          string `position:"Query" name:"VpcId"`
	ZoneId         string `position:"Query" name:"ZoneId"`
	PayType        string `position:"Query" name:"PayType"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreateDrdsInstanceRequest) Invoke(client goaliyun.Client) (*CreateDrdsInstanceResponse, error) {
	resp := &CreateDrdsInstanceResponse{}
	err := client.Request("drds", "CreateDrdsInstance", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDrdsInstanceResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      CreateDrdsInstanceData
}

type CreateDrdsInstanceData struct {
	OrderId            goaliyun.Integer
	DrdsInstanceIdList CreateDrdsInstanceDrdsInstanceIdListList
}

type CreateDrdsInstanceDrdsInstanceIdListList []goaliyun.String

func (list *CreateDrdsInstanceDrdsInstanceIdListList) UnmarshalJSON(data []byte) error {
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
