package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyResourceQueueRequest struct {
	ResourceOwnerId int64                          `position:"Query" name:"ResourceOwnerId"`
	ParentQueueId   int64                          `position:"Query" name:"ParentQueueId"`
	Name            string                         `position:"Query" name:"Name"`
	QualifiedName   string                         `position:"Query" name:"QualifiedName"`
	ResourcePoolId  int64                          `position:"Query" name:"ResourcePoolId"`
	Id              string                         `position:"Query" name:"Id"`
	ClusterId       string                         `position:"Query" name:"ClusterId"`
	Leaf            string                         `position:"Query" name:"Leaf"`
	Configs         *ModifyResourceQueueConfigList `position:"Query" type:"Repeated" name:"Config"`
	RegionId        string                         `position:"Query" name:"RegionId"`
}

func (req *ModifyResourceQueueRequest) Invoke(client goaliyun.Client) (*ModifyResourceQueueResponse, error) {
	resp := &ModifyResourceQueueResponse{}
	err := client.Request("emr", "ModifyResourceQueue", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyResourceQueueConfig struct {
	Id          int64  `name:"Id"`
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type ModifyResourceQueueResponse struct {
	RequestId goaliyun.String
}

type ModifyResourceQueueConfigList []ModifyResourceQueueConfig

func (list *ModifyResourceQueueConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyResourceQueueConfig)
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
