package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnDomainLogsRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	LogDay        string `position:"Query" name:"LogDay"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnDomainLogsRequest) Invoke(client goaliyun.Client) (*DescribeCdnDomainLogsResponse, error) {
	resp := &DescribeCdnDomainLogsResponse{}
	err := client.Request("cdn", "DescribeCdnDomainLogs", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnDomainLogsResponse struct {
	RequestId      goaliyun.String
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	TotalCount     goaliyun.Integer
	DomainLogModel DescribeCdnDomainLogsDomainLogModel
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       goaliyun.String
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogName   goaliyun.String
	LogPath   goaliyun.String
	LogSize   goaliyun.Integer
	StartTime goaliyun.String
	EndTime   goaliyun.String
}

type DescribeCdnDomainLogsDomainLogDetailList []DescribeCdnDomainLogsDomainLogDetail

func (list *DescribeCdnDomainLogsDomainLogDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainLogsDomainLogDetail)
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
