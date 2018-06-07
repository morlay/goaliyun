package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListResourcePoolRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PoolType        string `position:"Query" name:"PoolType"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListResourcePoolRequest) Invoke(client goaliyun.Client) (*ListResourcePoolResponse, error) {
	resp := &ListResourcePoolResponse{}
	err := client.Request("emr", "ListResourcePool", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListResourcePoolResponse struct {
	RequestId    goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	Total        goaliyun.Integer
	PoolInfoList ListResourcePoolPoolInfoList
}

type ListResourcePoolPoolInfo struct {
	QueueList                 ListResourcePoolQueueList
	EcmResourcePoolConfigList ListResourcePoolEcmResourcePoolConfig2List
	EcmResourcePool           ListResourcePoolEcmResourcePool
}

type ListResourcePoolQueue struct {
	EcmResourcePoolConfigList1 ListResourcePoolEcmResourcePoolConfigList
	EcmResourceQueue           ListResourcePoolEcmResourceQueue
}

type ListResourcePoolEcmResourcePoolConfig struct {
	Id          goaliyun.Integer
	ConfigKey   goaliyun.String
	ConfigValue goaliyun.String
	ConfigType  goaliyun.String
	Category    goaliyun.String
	Status      goaliyun.String
	Note        goaliyun.String
}

type ListResourcePoolEcmResourceQueue struct {
	Id             goaliyun.Integer
	Name           goaliyun.String
	QualifiedName  goaliyun.String
	QueueType      goaliyun.String
	ParentQueueId  goaliyun.Integer
	Leaf           bool
	Status         goaliyun.String
	UserId         goaliyun.String
	ResourcePoolId goaliyun.Integer
}

type ListResourcePoolEcmResourcePoolConfig2 struct {
	Id          goaliyun.Integer
	ConfigKey   goaliyun.String
	ConfigValue goaliyun.String
	ConfigType  goaliyun.String
	Category    goaliyun.String
	Status      goaliyun.String
	Note        goaliyun.String
}

type ListResourcePoolEcmResourcePool struct {
	Id             goaliyun.Integer
	Name           goaliyun.String
	PoolType       goaliyun.String
	Active         bool
	Note           goaliyun.String
	UserId         goaliyun.String
	YarnSiteConfig goaliyun.String
}

type ListResourcePoolPoolInfoList []ListResourcePoolPoolInfo

func (list *ListResourcePoolPoolInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolPoolInfo)
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

type ListResourcePoolQueueList []ListResourcePoolQueue

func (list *ListResourcePoolQueueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolQueue)
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

type ListResourcePoolEcmResourcePoolConfig2List []ListResourcePoolEcmResourcePoolConfig2

func (list *ListResourcePoolEcmResourcePoolConfig2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolEcmResourcePoolConfig2)
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

type ListResourcePoolEcmResourcePoolConfigList []ListResourcePoolEcmResourcePoolConfig

func (list *ListResourcePoolEcmResourcePoolConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolEcmResourcePoolConfig)
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
