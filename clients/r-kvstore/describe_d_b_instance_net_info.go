package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceNetInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceNetInfoRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceNetInfoResponse, error) {
	resp := &DescribeDBInstanceNetInfoResponse{}
	err := client.Request("r-kvstore", "DescribeDBInstanceNetInfo", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceNetInfoResponse struct {
	RequestId           goaliyun.String
	InstanceNetworkType goaliyun.String
	NetInfoItems        DescribeDBInstanceNetInfoInstanceNetInfoList
}

type DescribeDBInstanceNetInfoInstanceNetInfo struct {
	ConnectionString  goaliyun.String
	IPAddress         goaliyun.String
	Port              goaliyun.String
	VPCId             goaliyun.String
	VSwitchId         goaliyun.String
	DBInstanceNetType goaliyun.String
	IPType            goaliyun.String
	ExpiredTime       goaliyun.String
	Upgradeable       goaliyun.String
}

type DescribeDBInstanceNetInfoInstanceNetInfoList []DescribeDBInstanceNetInfoInstanceNetInfo

func (list *DescribeDBInstanceNetInfoInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoInstanceNetInfo)
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
