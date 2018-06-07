package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyPrepayInstanceSpecRequest struct {
	InstanceTypeConfigs *ModifyPrepayInstanceSpecInstanceTypeConfigList `position:"Query" type:"Repeated" name:"InstanceTypeConfig"`
	ResourceOwnerId     int64                                           `position:"Query" name:"ResourceOwnerId"`
	ClusterId           string                                          `position:"Query" name:"ClusterId"`
	RegionId            string                                          `position:"Query" name:"RegionId"`
}

func (req *ModifyPrepayInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyPrepayInstanceSpecResponse, error) {
	resp := &ModifyPrepayInstanceSpecResponse{}
	err := client.Request("emr", "ModifyPrepayInstanceSpec", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyPrepayInstanceSpecInstanceTypeConfig struct {
	HostGroupId  string `name:"HostGroupId"`
	InstanceType string `name:"InstanceType"`
}

type ModifyPrepayInstanceSpecResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}

type ModifyPrepayInstanceSpecInstanceTypeConfigList []ModifyPrepayInstanceSpecInstanceTypeConfig

func (list *ModifyPrepayInstanceSpecInstanceTypeConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyPrepayInstanceSpecInstanceTypeConfig)
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
