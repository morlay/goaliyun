package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainQoSRtRequest struct {
	Ip            string `position:"Query" name:"Ip"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainQoSRtRequest) Invoke(client goaliyun.Client) (*DescribeDomainQoSRtResponse, error) {
	resp := &DescribeDomainQoSRtResponse{}
	err := client.Request("cdn", "DescribeDomainQoSRt", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainQoSRtResponse struct {
	DomainName goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
	Ip         goaliyun.String
	Content    DescribeDomainQoSRtDataList
}

type DescribeDomainQoSRtData struct {
	More5s goaliyun.String
	Time   goaliyun.String
	More3s goaliyun.String
}

type DescribeDomainQoSRtDataList []DescribeDomainQoSRtData

func (list *DescribeDomainQoSRtDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainQoSRtData)
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
