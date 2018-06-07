package hsm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	Period          int64  `position:"Query" name:"Period"`
	PeriodUnit      string `position:"Query" name:"PeriodUnit"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	Quantity        int64  `position:"Query" name:"Quantity"`
	HsmDeviceType   string `position:"Query" name:"HsmDeviceType"`
	ClientToken     string `position:"Query" name:"ClientToken"`
	ZoneId          string `position:"Query" name:"ZoneId"`
	HsmOem          string `position:"Query" name:"HsmOem"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("hsm", "CreateInstance", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	RequestId   goaliyun.String
	InstanceIds CreateInstanceInstanceIdList
}

type CreateInstanceInstanceIdList []goaliyun.String

func (list *CreateInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
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
