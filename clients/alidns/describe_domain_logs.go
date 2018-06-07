package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainLogsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	GroupId      string `position:"Query" name:"GroupId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainLogsRequest) Invoke(client goaliyun.Client) (*DescribeDomainLogsResponse, error) {
	resp := &DescribeDomainLogsResponse{}
	err := client.Request("alidns", "DescribeDomainLogs", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainLogsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	DomainLogs DescribeDomainLogsDomainLogList
}

type DescribeDomainLogsDomainLog struct {
	ActionTime goaliyun.String
	DomainName goaliyun.String
	Action     goaliyun.String
	Message    goaliyun.String
	ClientIp   goaliyun.String
}

type DescribeDomainLogsDomainLogList []DescribeDomainLogsDomainLog

func (list *DescribeDomainLogsDomainLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainLogsDomainLog)
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
