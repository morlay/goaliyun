package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRegionDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRegionDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainRegionDataResponse, error) {
	resp := &DescribeDomainRegionDataResponse{}
	err := client.Request("cdn", "DescribeDomainRegionData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRegionDataResponse struct {
	RequestId    goaliyun.String
	DomainName   goaliyun.String
	DataInterval goaliyun.String
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	Value        DescribeDomainRegionDataRegionProportionDataList
}

type DescribeDomainRegionDataRegionProportionData struct {
	Region          goaliyun.String
	Proportion      goaliyun.String
	RegionEname     goaliyun.String
	AvgObjectSize   goaliyun.String
	AvgResponseTime goaliyun.String
	Bps             goaliyun.String
	ByteHitRate     goaliyun.String
	Qps             goaliyun.String
	ReqErrRate      goaliyun.String
	ReqHitRate      goaliyun.String
	AvgResponseRate goaliyun.String
	TotalBytes      goaliyun.String
	BytesProportion goaliyun.String
	TotalQuery      goaliyun.String
}

type DescribeDomainRegionDataRegionProportionDataList []DescribeDomainRegionDataRegionProportionData

func (list *DescribeDomainRegionDataRegionProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRegionDataRegionProportionData)
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
