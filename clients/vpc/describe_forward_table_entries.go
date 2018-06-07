package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeForwardTableEntriesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeForwardTableEntriesRequest) Invoke(client goaliyun.Client) (*DescribeForwardTableEntriesResponse, error) {
	resp := &DescribeForwardTableEntriesResponse{}
	err := client.Request("vpc", "DescribeForwardTableEntries", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeForwardTableEntriesResponse struct {
	RequestId           goaliyun.String
	TotalCount          goaliyun.Integer
	PageNumber          goaliyun.Integer
	PageSize            goaliyun.Integer
	ForwardTableEntries DescribeForwardTableEntriesForwardTableEntryList
}

type DescribeForwardTableEntriesForwardTableEntry struct {
	ForwardTableId goaliyun.String
	ForwardEntryId goaliyun.String
	ExternalIp     goaliyun.String
	ExternalPort   goaliyun.String
	IpProtocol     goaliyun.String
	InternalIp     goaliyun.String
	InternalPort   goaliyun.String
	Status         goaliyun.String
}

type DescribeForwardTableEntriesForwardTableEntryList []DescribeForwardTableEntriesForwardTableEntry

func (list *DescribeForwardTableEntriesForwardTableEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTableEntriesForwardTableEntry)
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
