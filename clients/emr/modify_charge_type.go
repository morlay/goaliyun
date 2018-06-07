package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyChargeTypeRequest struct {
	ResourceOwnerId   int64                                 `position:"Query" name:"ResourceOwnerId"`
	ChargeTypeConfigs *ModifyChargeTypeChargeTypeConfigList `position:"Query" type:"Repeated" name:"ChargeTypeConfig"`
	ClusterId         string                                `position:"Query" name:"ClusterId"`
	RegionId          string                                `position:"Query" name:"RegionId"`
}

func (req *ModifyChargeTypeRequest) Invoke(client goaliyun.Client) (*ModifyChargeTypeResponse, error) {
	resp := &ModifyChargeTypeResponse{}
	err := client.Request("emr", "ModifyChargeType", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyChargeTypeChargeTypeConfig struct {
	HostGroupId string `name:"HostGroupId"`
	ChargeType  string `name:"ChargeType"`
	Period      int64  `name:"Period"`
}

type ModifyChargeTypeResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}

type ModifyChargeTypeChargeTypeConfigList []ModifyChargeTypeChargeTypeConfig

func (list *ModifyChargeTypeChargeTypeConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyChargeTypeChargeTypeConfig)
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
