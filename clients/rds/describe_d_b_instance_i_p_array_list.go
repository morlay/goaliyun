package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceIPArrayListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	WhitelistNetworkType string `position:"Query" name:"WhitelistNetworkType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceIPArrayListRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceIPArrayListResponse, error) {
	resp := &DescribeDBInstanceIPArrayListResponse{}
	err := client.Request("rds", "DescribeDBInstanceIPArrayList", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceIPArrayListResponse struct {
	RequestId goaliyun.String
	Items     DescribeDBInstanceIPArrayListDBInstanceIPArrayList
}

type DescribeDBInstanceIPArrayListDBInstanceIPArray struct {
	DBInstanceIPArrayName      goaliyun.String
	DBInstanceIPArrayAttribute goaliyun.String
	SecurityIPList             goaliyun.String
	WhitelistNetworkType       goaliyun.String
}

type DescribeDBInstanceIPArrayListDBInstanceIPArrayList []DescribeDBInstanceIPArrayListDBInstanceIPArray

func (list *DescribeDBInstanceIPArrayListDBInstanceIPArrayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceIPArrayListDBInstanceIPArray)
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
