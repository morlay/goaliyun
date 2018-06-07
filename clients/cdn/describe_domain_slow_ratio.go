package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainSlowRatioRequest struct {
	StartTime     string `position:"Query" name:"StartTime"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainSlowRatioRequest) Invoke(client goaliyun.Client) (*DescribeDomainSlowRatioResponse, error) {
	resp := &DescribeDomainSlowRatioResponse{}
	err := client.Request("cdn", "DescribeDomainSlowRatio", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainSlowRatioResponse struct {
	EndTime                  goaliyun.String
	DataInterval             goaliyun.Integer
	PageNumber               goaliyun.Integer
	PageSize                 goaliyun.Integer
	TotalCount               goaliyun.Integer
	StartTime                goaliyun.String
	SlowRatioDataPerInterval DescribeDomainSlowRatioSlowRatioDataList
}

type DescribeDomainSlowRatioSlowRatioData struct {
	TotalUsers   goaliyun.Integer
	SlowUsers    goaliyun.Integer
	SlowRatio    goaliyun.Float
	RegionNameZh goaliyun.String
	RegionNameEn goaliyun.String
	IspNameZh    goaliyun.String
	IspNameEn    goaliyun.String
	Time         goaliyun.String
}

type DescribeDomainSlowRatioSlowRatioDataList []DescribeDomainSlowRatioSlowRatioData

func (list *DescribeDomainSlowRatioSlowRatioDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSlowRatioSlowRatioData)
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
