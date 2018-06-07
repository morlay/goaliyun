package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancesByExpireTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	Expired              string `position:"Query" name:"Expired"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ExpirePeriod         int64  `position:"Query" name:"ExpirePeriod"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancesByExpireTimeRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancesByExpireTimeResponse, error) {
	resp := &DescribeDBInstancesByExpireTimeResponse{}
	err := client.Request("rds", "DescribeDBInstancesByExpireTime", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancesByExpireTimeResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList
}

type DescribeDBInstancesByExpireTimeDBInstanceExpireTime struct {
	DBInstanceId          goaliyun.String
	DBInstanceDescription goaliyun.String
	ExpireTime            goaliyun.String
	DBInstanceStatus      goaliyun.String
	LockMode              goaliyun.String
}

type DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList []DescribeDBInstancesByExpireTimeDBInstanceExpireTime

func (list *DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesByExpireTimeDBInstanceExpireTime)
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
