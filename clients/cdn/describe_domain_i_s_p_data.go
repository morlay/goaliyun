package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainISPDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainISPDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainISPDataResponse, error) {
	resp := &DescribeDomainISPDataResponse{}
	err := client.Request("cdn", "DescribeDomainISPData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainISPDataResponse struct {
	RequestId    goaliyun.String
	DomainName   goaliyun.String
	DataInterval goaliyun.String
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	Value        DescribeDomainISPDataISPProportionDataList
}

type DescribeDomainISPDataISPProportionData struct {
	ISP             goaliyun.String
	Proportion      goaliyun.String
	IspEname        goaliyun.String
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

type DescribeDomainISPDataISPProportionDataList []DescribeDomainISPDataISPProportionData

func (list *DescribeDomainISPDataISPProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainISPDataISPProportionData)
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
