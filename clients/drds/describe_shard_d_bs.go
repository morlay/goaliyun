package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeShardDBsRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeShardDBsRequest) Invoke(client goaliyun.Client) (*DescribeShardDBsResponse, error) {
	resp := &DescribeShardDBsResponse{}
	err := client.Request("drds", "DescribeShardDBs", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeShardDBsResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeShardDBsDbIntancePairList
}

type DescribeShardDBsDbIntancePair struct {
	SubDbName    goaliyun.String
	InstanceName goaliyun.String
}

type DescribeShardDBsDbIntancePairList []DescribeShardDBsDbIntancePair

func (list *DescribeShardDBsDbIntancePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeShardDBsDbIntancePair)
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
