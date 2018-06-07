package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterNetInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterNetInfoRequest) Invoke(client goaliyun.Client) (*DescribeClusterNetInfoResponse, error) {
	resp := &DescribeClusterNetInfoResponse{}
	err := client.Request("polardb", "DescribeClusterNetInfo", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterNetInfoResponse struct {
	RequestId            goaliyun.String
	DBClusterNetworkType goaliyun.String
	DBInstanceNetInfos   DescribeClusterNetInfoDBInstanceNetInfoList
}

type DescribeClusterNetInfoDBInstanceNetInfo struct {
	ConnectionString goaliyun.String
	IPAddress        goaliyun.String
	IPType           goaliyun.String
	Port             goaliyun.String
	VPCId            goaliyun.String
	VSwitchId        goaliyun.String
}

type DescribeClusterNetInfoDBInstanceNetInfoList []DescribeClusterNetInfoDBInstanceNetInfo

func (list *DescribeClusterNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterNetInfoDBInstanceNetInfo)
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
