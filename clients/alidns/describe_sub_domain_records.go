package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSubDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	SubDomain    string `position:"Query" name:"SubDomain"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Type         string `position:"Query" name:"Type"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeSubDomainRecordsRequest) Invoke(client goaliyun.Client) (*DescribeSubDomainRecordsResponse, error) {
	resp := &DescribeSubDomainRecordsResponse{}
	err := client.Request("alidns", "DescribeSubDomainRecords", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSubDomainRecordsResponse struct {
	RequestId     goaliyun.String
	TotalCount    goaliyun.Integer
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	DomainRecords DescribeSubDomainRecordsRecordList
}

type DescribeSubDomainRecordsRecord struct {
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

type DescribeSubDomainRecordsRecordList []DescribeSubDomainRecordsRecord

func (list *DescribeSubDomainRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSubDomainRecordsRecord)
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
