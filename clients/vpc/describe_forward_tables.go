package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeForwardTablesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeForwardTablesRequest) Invoke(client goaliyun.Client) (*DescribeForwardTablesResponse, error) {
	resp := &DescribeForwardTablesResponse{}
	err := client.Request("vpc", "DescribeForwardTables", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeForwardTablesResponse struct {
	RequestId     goaliyun.String
	TotalCount    goaliyun.Integer
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	ForwardTables DescribeForwardTablesForwardTableList
}

type DescribeForwardTablesForwardTable struct {
	NatGatewayId   goaliyun.String
	ForwardTableId goaliyun.String
	CreationTime   goaliyun.String
	ForwardEntrys  DescribeForwardTablesForwardEntryList
}

type DescribeForwardTablesForwardEntry struct {
	ForwardTableId goaliyun.String
	ForwardEntryId goaliyun.String
	ExternalIp     goaliyun.String
	ExternalPort   goaliyun.String
	IpProtocol     goaliyun.String
	InternalIp     goaliyun.String
	InternalPort   goaliyun.String
	Status         goaliyun.String
}

type DescribeForwardTablesForwardTableList []DescribeForwardTablesForwardTable

func (list *DescribeForwardTablesForwardTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTablesForwardTable)
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

type DescribeForwardTablesForwardEntryList []DescribeForwardTablesForwardEntry

func (list *DescribeForwardTablesForwardEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTablesForwardEntry)
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
