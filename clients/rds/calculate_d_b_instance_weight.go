package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CalculateDBInstanceWeightRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CalculateDBInstanceWeightRequest) Invoke(client goaliyun.Client) (*CalculateDBInstanceWeightResponse, error) {
	resp := &CalculateDBInstanceWeightResponse{}
	err := client.Request("rds", "CalculateDBInstanceWeight", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CalculateDBInstanceWeightResponse struct {
	RequestId goaliyun.String
	Items     CalculateDBInstanceWeightDBInstanceWeightList
}

type CalculateDBInstanceWeightDBInstanceWeight struct {
	DBInstanceId   goaliyun.String
	DBInstanceType goaliyun.String
	Availability   goaliyun.String
	Weight         goaliyun.String
}

type CalculateDBInstanceWeightDBInstanceWeightList []CalculateDBInstanceWeightDBInstanceWeight

func (list *CalculateDBInstanceWeightDBInstanceWeightList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CalculateDBInstanceWeightDBInstanceWeight)
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
