package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	RRKeyWord    string `position:"Query" name:"RRKeyWord"`
	TypeKeyWord  string `position:"Query" name:"TypeKeyWord"`
	ValueKeyWord string `position:"Query" name:"ValueKeyWord"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRecordsRequest) Invoke(client goaliyun.Client) (*DescribeDomainRecordsResponse, error) {
	resp := &DescribeDomainRecordsResponse{}
	err := client.Request("alidns", "DescribeDomainRecords", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRecordsResponse struct {
	RequestId     goaliyun.String
	TotalCount    goaliyun.Integer
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	DomainRecords DescribeDomainRecordsRecordList
}

type DescribeDomainRecordsRecord struct {
	DomainName goaliyun.String
	RecordId   goaliyun.String
	RR         goaliyun.String
	Type       goaliyun.String
	Value      goaliyun.String
	TTL        goaliyun.Integer
	Priority   goaliyun.Integer
	Line       goaliyun.String
	Status     goaliyun.String
	Locked     bool
	Weight     goaliyun.Integer
}

type DescribeDomainRecordsRecordList []DescribeDomainRecordsRecord

func (list *DescribeDomainRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRecordsRecord)
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
