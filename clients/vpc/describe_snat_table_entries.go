package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSnatTableEntriesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnatTableEntriesRequest) Invoke(client goaliyun.Client) (*DescribeSnatTableEntriesResponse, error) {
	resp := &DescribeSnatTableEntriesResponse{}
	err := client.Request("vpc", "DescribeSnatTableEntries", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnatTableEntriesResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageSize         goaliyun.Integer
	SnatTableEntries DescribeSnatTableEntriesSnatTableEntryList
}

type DescribeSnatTableEntriesSnatTableEntry struct {
	SnatTableId     goaliyun.String
	SnatEntryId     goaliyun.String
	SourceVSwitchId goaliyun.String
	SourceCIDR      goaliyun.String
	SnatIp          goaliyun.String
	Status          goaliyun.String
}

type DescribeSnatTableEntriesSnatTableEntryList []DescribeSnatTableEntriesSnatTableEntry

func (list *DescribeSnatTableEntriesSnatTableEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnatTableEntriesSnatTableEntry)
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
