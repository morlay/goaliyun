package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateResourceQueueRequest struct {
	ResourceOwnerId int64                          `position:"Query" name:"ResourceOwnerId"`
	ParentQueueId   int64                          `position:"Query" name:"ParentQueueId"`
	Name            string                         `position:"Query" name:"Name"`
	QualifiedName   string                         `position:"Query" name:"QualifiedName"`
	ResourcePoolId  int64                          `position:"Query" name:"ResourcePoolId"`
	ClusterId       string                         `position:"Query" name:"ClusterId"`
	Leaf            string                         `position:"Query" name:"Leaf"`
	Configs         *CreateResourceQueueConfigList `position:"Query" type:"Repeated" name:"Config"`
	RegionId        string                         `position:"Query" name:"RegionId"`
}

func (req *CreateResourceQueueRequest) Invoke(client goaliyun.Client) (*CreateResourceQueueResponse, error) {
	resp := &CreateResourceQueueResponse{}
	err := client.Request("emr", "CreateResourceQueue", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateResourceQueueConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type CreateResourceQueueResponse struct {
	RequestId goaliyun.String
}

type CreateResourceQueueConfigList []CreateResourceQueueConfig

func (list *CreateResourceQueueConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateResourceQueueConfig)
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
