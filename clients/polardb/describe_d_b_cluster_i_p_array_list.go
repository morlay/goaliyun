package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBClusterIPArrayListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBClusterIPArrayListRequest) Invoke(client goaliyun.Client) (*DescribeDBClusterIPArrayListResponse, error) {
	resp := &DescribeDBClusterIPArrayListResponse{}
	err := client.Request("polardb", "DescribeDBClusterIPArrayList", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBClusterIPArrayListResponse struct {
	RequestId goaliyun.String
	Items     DescribeDBClusterIPArrayListDBClusterIPArrayList
}

type DescribeDBClusterIPArrayListDBClusterIPArray struct {
	DBClusterIPArrayName      goaliyun.String
	DBClusterIPArrayAttribute goaliyun.String
	SecurityIPList            goaliyun.String
}

type DescribeDBClusterIPArrayListDBClusterIPArrayList []DescribeDBClusterIPArrayListDBClusterIPArray

func (list *DescribeDBClusterIPArrayListDBClusterIPArrayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterIPArrayListDBClusterIPArray)
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
