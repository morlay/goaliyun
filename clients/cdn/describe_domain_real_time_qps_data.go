package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRealTimeQpsDataRequest struct {
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRealTimeQpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainRealTimeQpsDataResponse, error) {
	resp := &DescribeDomainRealTimeQpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainRealTimeQpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRealTimeQpsDataResponse struct {
	RequestId goaliyun.String
	Data      DescribeDomainRealTimeQpsDataQpsModelList
}

type DescribeDomainRealTimeQpsDataQpsModel struct {
	Qps       goaliyun.Float
	TimeStamp goaliyun.String
}

type DescribeDomainRealTimeQpsDataQpsModelList []DescribeDomainRealTimeQpsDataQpsModel

func (list *DescribeDomainRealTimeQpsDataQpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeQpsDataQpsModel)
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
