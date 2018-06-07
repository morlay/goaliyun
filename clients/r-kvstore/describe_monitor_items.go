package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMonitorItemsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMonitorItemsRequest) Invoke(client goaliyun.Client) (*DescribeMonitorItemsResponse, error) {
	resp := &DescribeMonitorItemsResponse{}
	err := client.Request("r-kvstore", "DescribeMonitorItems", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMonitorItemsResponse struct {
	RequestId    goaliyun.String
	MonitorItems DescribeMonitorItemsKVStoreMonitorItemList
}

type DescribeMonitorItemsKVStoreMonitorItem struct {
	MonitorKey goaliyun.String
	Unit       goaliyun.String
}

type DescribeMonitorItemsKVStoreMonitorItemList []DescribeMonitorItemsKVStoreMonitorItem

func (list *DescribeMonitorItemsKVStoreMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorItemsKVStoreMonitorItem)
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
