package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRealTimeByteHitRateDataRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRealTimeByteHitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainRealTimeByteHitRateDataResponse, error) {
	resp := &DescribeDomainRealTimeByteHitRateDataResponse{}
	err := client.Request("cdn", "DescribeDomainRealTimeByteHitRateData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRealTimeByteHitRateDataResponse struct {
	RequestId goaliyun.String
	Data      DescribeDomainRealTimeByteHitRateDataByteHitRateDataModelList
}

type DescribeDomainRealTimeByteHitRateDataByteHitRateDataModel struct {
	ByteHitRate goaliyun.Float
	TimeStamp   goaliyun.String
}

type DescribeDomainRealTimeByteHitRateDataByteHitRateDataModelList []DescribeDomainRealTimeByteHitRateDataByteHitRateDataModel

func (list *DescribeDomainRealTimeByteHitRateDataByteHitRateDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeByteHitRateDataByteHitRateDataModel)
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
