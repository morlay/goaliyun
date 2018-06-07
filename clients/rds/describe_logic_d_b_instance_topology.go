package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLogicDBInstanceTopologyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeLogicDBInstanceTopologyRequest) Invoke(client goaliyun.Client) (*DescribeLogicDBInstanceTopologyResponse, error) {
	resp := &DescribeLogicDBInstanceTopologyResponse{}
	err := client.Request("rds", "DescribeLogicDBInstanceTopology", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLogicDBInstanceTopologyResponse struct {
	DBInstanceId          goaliyun.Integer
	DBInstanceName        goaliyun.String
	DBInstanceStatus      goaliyun.Integer
	DBInstanceStatusDesc  goaliyun.String
	DBInstanceConnType    goaliyun.String
	DBInstanceDescription goaliyun.String
	Engine                goaliyun.String
	EngineVersion         goaliyun.String
	Items                 DescribeLogicDBInstanceTopologyLogicDBInstanceParameterList
}

type DescribeLogicDBInstanceTopologyLogicDBInstanceParameter struct {
	DBInstanceID          goaliyun.Integer
	DBInstanceName        goaliyun.String
	DBInstanceStatus      goaliyun.Integer
	DBInstanceStatusDesc  goaliyun.String
	DBInstanceConnType    goaliyun.String
	DBInstanceDescription goaliyun.String
	Engine                goaliyun.String
	EngineVersion         goaliyun.String
	CharacterType         goaliyun.String
}

type DescribeLogicDBInstanceTopologyLogicDBInstanceParameterList []DescribeLogicDBInstanceTopologyLogicDBInstanceParameter

func (list *DescribeLogicDBInstanceTopologyLogicDBInstanceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLogicDBInstanceTopologyLogicDBInstanceParameter)
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
