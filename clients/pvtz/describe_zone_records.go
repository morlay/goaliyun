package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZoneRecordsRequest struct {
	PageSize     int64  `position:"Query" name:"PageSize"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Tag          string `position:"Query" name:"Tag"`
	Keyword      string `position:"Query" name:"Keyword"`
	Lang         string `position:"Query" name:"Lang"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeZoneRecordsRequest) Invoke(client goaliyun.Client) (*DescribeZoneRecordsResponse, error) {
	resp := &DescribeZoneRecordsResponse{}
	err := client.Request("pvtz", "DescribeZoneRecords", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZoneRecordsResponse struct {
	RequestId  goaliyun.String
	TotalItems goaliyun.Integer
	TotalPages goaliyun.Integer
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	Records    DescribeZoneRecordsRecordList
}

type DescribeZoneRecordsRecord struct {
	RecordId goaliyun.Integer
	Rr       goaliyun.String
	Type     goaliyun.String
	Ttl      goaliyun.Integer
	Priority goaliyun.Integer
	Value    goaliyun.String
	Status   goaliyun.String
}

type DescribeZoneRecordsRecordList []DescribeZoneRecordsRecord

func (list *DescribeZoneRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneRecordsRecord)
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
