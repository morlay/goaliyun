package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRealTimeBpsDataRequest struct {
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRealTimeBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainRealTimeBpsDataResponse, error) {
	resp := &DescribeDomainRealTimeBpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainRealTimeBpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRealTimeBpsDataResponse struct {
	RequestId goaliyun.String
	Data      DescribeDomainRealTimeBpsDataBpsModelList
}

type DescribeDomainRealTimeBpsDataBpsModel struct {
	Bps       goaliyun.Float
	TimeStamp goaliyun.String
}

type DescribeDomainRealTimeBpsDataBpsModelList []DescribeDomainRealTimeBpsDataBpsModel

func (list *DescribeDomainRealTimeBpsDataBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeBpsDataBpsModel)
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
