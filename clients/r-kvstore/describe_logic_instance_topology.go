package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLogicInstanceTopologyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeLogicInstanceTopologyRequest) Invoke(client goaliyun.Client) (*DescribeLogicInstanceTopologyResponse, error) {
	resp := &DescribeLogicInstanceTopologyResponse{}
	err := client.Request("r-kvstore", "DescribeLogicInstanceTopology", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLogicInstanceTopologyResponse struct {
	RequestId      goaliyun.String
	InstanceId     goaliyun.String
	RedisProxyList DescribeLogicInstanceTopologyNodeInfoList
	RedisShardList DescribeLogicInstanceTopologyNodeInfoList
}

type DescribeLogicInstanceTopologyNodeInfo struct {
	NodeId     goaliyun.String
	Connection goaliyun.String
	Bandwidth  goaliyun.String
	Capacity   goaliyun.String
	NodeType   goaliyun.String
}

type DescribeLogicInstanceTopologyNodeInfoList []DescribeLogicInstanceTopologyNodeInfo

func (list *DescribeLogicInstanceTopologyNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLogicInstanceTopologyNodeInfo)
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
