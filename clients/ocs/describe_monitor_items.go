package ocs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMonitorItemsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMonitorItemsRequest) Invoke(client goaliyun.Client) (*DescribeMonitorItemsResponse, error) {
	resp := &DescribeMonitorItemsResponse{}
	err := client.Request("ocs", "DescribeMonitorItems", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMonitorItemsResponse struct {
	RequestId                     goaliyun.String
	GetOcsMonitorItemsResponseDTO DescribeMonitorItemsGetOcsMonitorItemsResponseDTO
}

type DescribeMonitorItemsGetOcsMonitorItemsResponseDTO struct {
	MonitorItems DescribeMonitorItemsMonitorItemList
}

type DescribeMonitorItemsMonitorItem struct {
	MonitorKey goaliyun.String
	Unit       goaliyun.String
}

type DescribeMonitorItemsMonitorItemList []DescribeMonitorItemsMonitorItem

func (list *DescribeMonitorItemsMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorItemsMonitorItem)
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
