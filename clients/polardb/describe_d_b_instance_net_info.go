package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceNetInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceNetInfoRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceNetInfoResponse, error) {
	resp := &DescribeDBInstanceNetInfoResponse{}
	err := client.Request("polardb", "DescribeDBInstanceNetInfo", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceNetInfoResponse struct {
	RequestId           goaliyun.String
	InstanceNetworkType goaliyun.String
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	ConnectionString goaliyun.String
	IPAddress        goaliyun.String
	IPType           goaliyun.String
	Port             goaliyun.String
	VPCId            goaliyun.String
	VSwitchId        goaliyun.String
}

type DescribeDBInstanceNetInfoDBInstanceNetInfoList []DescribeDBInstanceNetInfoDBInstanceNetInfo

func (list *DescribeDBInstanceNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceNetInfo)
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
