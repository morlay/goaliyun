package sas_api

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpgradeInstanceRequest struct {
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	BuyNumber   int64  `position:"Query" name:"BuyNumber"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	VersionCode string `position:"Query" name:"VersionCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *UpgradeInstanceRequest) Invoke(client goaliyun.Client) (*UpgradeInstanceResponse, error) {
	resp := &UpgradeInstanceResponse{}
	err := client.Request("sas-api", "UpgradeInstance", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeInstanceResponse struct {
	Code        goaliyun.String
	Message     goaliyun.String
	OrderId     goaliyun.String
	InstanceId  goaliyun.String
	RequestId   goaliyun.String
	InstanceIds UpgradeInstanceInstanceIdList
}

type UpgradeInstanceInstanceIdList []goaliyun.String

func (list *UpgradeInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
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
