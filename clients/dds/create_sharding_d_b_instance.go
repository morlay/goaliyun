package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateShardingDBInstanceRequest struct {
	ResourceOwnerId       int64                                     `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string                                    `position:"Query" name:"ClientToken"`
	EngineVersion         string                                    `position:"Query" name:"EngineVersion"`
	NetworkType           string                                    `position:"Query" name:"NetworkType"`
	ReplicaSets           *CreateShardingDBInstanceReplicaSetList   `position:"Query" type:"Repeated" name:"ReplicaSet"`
	StorageEngine         string                                    `position:"Query" name:"StorageEngine"`
	SecurityToken         string                                    `position:"Query" name:"SecurityToken"`
	Engine                string                                    `position:"Query" name:"Engine"`
	DBInstanceDescription string                                    `position:"Query" name:"DBInstanceDescription"`
	Period                int64                                     `position:"Query" name:"Period"`
	RestoreTime           string                                    `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string                                    `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId       string                                    `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount          string                                    `position:"Query" name:"OwnerAccount"`
	ConfigServers         *CreateShardingDBInstanceConfigServerList `position:"Query" type:"Repeated" name:"ConfigServer"`
	OwnerId               int64                                     `position:"Query" name:"OwnerId"`
	Mongoss               *CreateShardingDBInstanceMongosList       `position:"Query" type:"Repeated" name:"Mongos"`
	SecurityIPList        string                                    `position:"Query" name:"SecurityIPList"`
	VSwitchId             string                                    `position:"Query" name:"VSwitchId"`
	AccountPassword       string                                    `position:"Query" name:"AccountPassword"`
	VpcId                 string                                    `position:"Query" name:"VpcId"`
	ZoneId                string                                    `position:"Query" name:"ZoneId"`
	ChargeType            string                                    `position:"Query" name:"ChargeType"`
	RegionId              string                                    `position:"Query" name:"RegionId"`
}

func (req *CreateShardingDBInstanceRequest) Invoke(client goaliyun.Client) (*CreateShardingDBInstanceResponse, error) {
	resp := &CreateShardingDBInstanceResponse{}
	err := client.Request("dds", "CreateShardingDBInstance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateShardingDBInstanceReplicaSet struct {
	_class  string `name:"_class"`
	Storage int64  `name:"Storage"`
}

type CreateShardingDBInstanceConfigServer struct {
	_class  string `name:"_class"`
	Storage int64  `name:"Storage"`
}

type CreateShardingDBInstanceMongos struct {
	_class string `name:"_class"`
}

type CreateShardingDBInstanceResponse struct {
	RequestId    goaliyun.String
	OrderId      goaliyun.String
	DBInstanceId goaliyun.String
}

type CreateShardingDBInstanceReplicaSetList []CreateShardingDBInstanceReplicaSet

func (list *CreateShardingDBInstanceReplicaSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceReplicaSet)
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

type CreateShardingDBInstanceConfigServerList []CreateShardingDBInstanceConfigServer

func (list *CreateShardingDBInstanceConfigServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceConfigServer)
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

type CreateShardingDBInstanceMongosList []CreateShardingDBInstanceMongos

func (list *CreateShardingDBInstanceMongosList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceMongos)
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
