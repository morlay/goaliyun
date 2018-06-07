package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnDomainLogsRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DomainName           string `position:"Query" name:"DomainName"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	LogDay               string `position:"Query" name:"LogDay"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnDomainLogsRequest) Invoke(client goaliyun.Client) (*DescribeCdnDomainLogsResponse, error) {
	resp := &DescribeCdnDomainLogsResponse{}
	err := client.Request("vod", "DescribeCdnDomainLogs", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnDomainLogsResponse struct {
	RequestId      goaliyun.String
	PageNo         goaliyun.Integer
	PageSize       goaliyun.Integer
	Total          goaliyun.Integer
	DomainLogModel DescribeCdnDomainLogsDomainLogModel
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       goaliyun.String
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogPath   goaliyun.String
	StartTime goaliyun.String
	EndTime   goaliyun.String
	LogSize   goaliyun.Integer
	LogName   goaliyun.String
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
