package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyResourcePoolRequest struct {
	ResourceOwnerId int64                         `position:"Query" name:"ResourceOwnerId"`
	Name            string                        `position:"Query" name:"Name"`
	Active          string                        `position:"Query" name:"Active"`
	Id              string                        `position:"Query" name:"Id"`
	ClusterId       string                        `position:"Query" name:"ClusterId"`
	Yarnsiteconfig  string                        `position:"Query" name:"Yarnsiteconfig"`
	Configs         *ModifyResourcePoolConfigList `position:"Query" type:"Repeated" name:"Config"`
	RegionId        string                        `position:"Query" name:"RegionId"`
}

func (req *ModifyResourcePoolRequest) Invoke(client goaliyun.Client) (*ModifyResourcePoolResponse, error) {
	resp := &ModifyResourcePoolResponse{}
	err := client.Request("emr", "ModifyResourcePool", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyResourcePoolConfig struct {
	Id          string `name:"Id"`
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type ModifyResourcePoolResponse struct {
	RequestId goaliyun.String
}

type ModifyResourcePoolConfigList []ModifyResourcePoolConfig

func (list *ModifyResourcePoolConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyResourcePoolConfig)
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
