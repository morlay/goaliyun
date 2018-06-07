package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeReplicaInitializeProgressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeReplicaInitializeProgressRequest) Invoke(client goaliyun.Client) (*DescribeReplicaInitializeProgressResponse, error) {
	resp := &DescribeReplicaInitializeProgressResponse{}
	err := client.Request("r-kvstore", "DescribeReplicaInitializeProgress", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeReplicaInitializeProgressResponse struct {
	RequestId goaliyun.String
	Items     DescribeReplicaInitializeProgressItemsItemList
}

type DescribeReplicaInitializeProgressItemsItem struct {
	ReplicaId   goaliyun.String
	Status      goaliyun.String
	Progress    goaliyun.String
	FinishTime  goaliyun.String
	CurrentStep goaliyun.String
}

type DescribeReplicaInitializeProgressItemsItemList []DescribeReplicaInitializeProgressItemsItem

func (list *DescribeReplicaInitializeProgressItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaInitializeProgressItemsItem)
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
