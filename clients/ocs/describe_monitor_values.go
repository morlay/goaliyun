package ocs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMonitorValuesRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	MonitorKeys          string `position:"Query" name:"MonitorKeys"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMonitorValuesRequest) Invoke(client goaliyun.Client) (*DescribeMonitorValuesResponse, error) {
	resp := &DescribeMonitorValuesResponse{}
	err := client.Request("ocs", "DescribeMonitorValues", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMonitorValuesResponse struct {
	RequestId                   goaliyun.String
	GetOcsMonitorValuesResponse DescribeMonitorValuesGetOcsMonitorValuesResponse
}

type DescribeMonitorValuesGetOcsMonitorValuesResponse struct {
	Date           goaliyun.String
	OcsInstanceIds DescribeMonitorValuesOcsMonitorDTOList
}

type DescribeMonitorValuesOcsMonitorDTO struct {
	OcsInstanceId goaliyun.String
	MonitorKeys   DescribeMonitorValuesOcsMonitorKeyDTOList
}

type DescribeMonitorValuesOcsMonitorKeyDTO struct {
	MonitorKey goaliyun.String
	Value      goaliyun.String
	Unit       goaliyun.String
	Date       goaliyun.String
}

type DescribeMonitorValuesOcsMonitorDTOList []DescribeMonitorValuesOcsMonitorDTO

func (list *DescribeMonitorValuesOcsMonitorDTOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorValuesOcsMonitorDTO)
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

type DescribeMonitorValuesOcsMonitorKeyDTOList []DescribeMonitorValuesOcsMonitorKeyDTO

func (list *DescribeMonitorValuesOcsMonitorKeyDTOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorValuesOcsMonitorKeyDTO)
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
