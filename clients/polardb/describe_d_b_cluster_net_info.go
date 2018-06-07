package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBClusterNetInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBClusterNetInfoRequest) Invoke(client goaliyun.Client) (*DescribeDBClusterNetInfoResponse, error) {
	resp := &DescribeDBClusterNetInfoResponse{}
	err := client.Request("polardb", "DescribeDBClusterNetInfo", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBClusterNetInfoResponse struct {
	RequestId          goaliyun.String
	ClusterNetworkType goaliyun.String
	DBClusterNetInfos  DescribeDBClusterNetInfoDBClusterNetInfoList
}

type DescribeDBClusterNetInfoDBClusterNetInfo struct {
	ConnectionString goaliyun.String
	IPAddress        goaliyun.String
	IPType           goaliyun.String
	Port             goaliyun.String
	VPCId            goaliyun.String
	VSwitchId        goaliyun.String
}

type DescribeDBClusterNetInfoDBClusterNetInfoList []DescribeDBClusterNetInfoDBClusterNetInfo

func (list *DescribeDBClusterNetInfoDBClusterNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterNetInfoDBClusterNetInfo)
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
