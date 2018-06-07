package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateResourcePoolRequest struct {
	Note            string                        `position:"Query" name:"Note"`
	ResourceOwnerId int64                         `position:"Query" name:"ResourceOwnerId"`
	Name            string                        `position:"Query" name:"Name"`
	Active          string                        `position:"Query" name:"Active"`
	ClusterId       string                        `position:"Query" name:"ClusterId"`
	YarnSiteConfig  string                        `position:"Query" name:"YarnSiteConfig"`
	Configs         *CreateResourcePoolConfigList `position:"Query" type:"Repeated" name:"Config"`
	PoolType        string                        `position:"Query" name:"PoolType"`
	RegionId        string                        `position:"Query" name:"RegionId"`
}

func (req *CreateResourcePoolRequest) Invoke(client goaliyun.Client) (*CreateResourcePoolResponse, error) {
	resp := &CreateResourcePoolResponse{}
	err := client.Request("emr", "CreateResourcePool", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateResourcePoolConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	ConfigType  string `name:"ConfigType"`
	TargetId    string `name:"TargetId"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type CreateResourcePoolResponse struct {
	RequestId goaliyun.String
}

type CreateResourcePoolConfigList []CreateResourcePoolConfig

func (list *CreateResourcePoolConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateResourcePoolConfig)
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
