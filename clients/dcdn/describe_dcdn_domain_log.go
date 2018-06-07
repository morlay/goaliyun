package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainLogRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainLogRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainLogResponse, error) {
	resp := &DescribeDcdnDomainLogResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainLog", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainLogResponse struct {
	RequestId        goaliyun.String
	DomainName       goaliyun.String
	DomainLogDetails DescribeDcdnDomainLogDomainLogDetailList
}

type DescribeDcdnDomainLogDomainLogDetail struct {
	LogCount  goaliyun.Integer
	PageInfos DescribeDcdnDomainLogPageInfoDetailList
	LogInfos  DescribeDcdnDomainLogLogInfoDetailList
}

type DescribeDcdnDomainLogPageInfoDetail struct {
	PageIndex goaliyun.Integer
	PageSize  goaliyun.Integer
	Total     goaliyun.Integer
}

type DescribeDcdnDomainLogLogInfoDetail struct {
	LogName   goaliyun.String
	LogPath   goaliyun.String
	LogSize   goaliyun.Integer
	StartTime goaliyun.String
	EndTime   goaliyun.String
}

type DescribeDcdnDomainLogDomainLogDetailList []DescribeDcdnDomainLogDomainLogDetail

func (list *DescribeDcdnDomainLogDomainLogDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainLogDomainLogDetail)
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

type DescribeDcdnDomainLogPageInfoDetailList []DescribeDcdnDomainLogPageInfoDetail

func (list *DescribeDcdnDomainLogPageInfoDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainLogPageInfoDetail)
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

type DescribeDcdnDomainLogLogInfoDetailList []DescribeDcdnDomainLogLogInfoDetail

func (list *DescribeDcdnDomainLogLogInfoDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainLogLogInfoDetail)
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
