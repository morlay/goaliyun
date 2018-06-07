package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRealTimeReqHitRateDataRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRealTimeReqHitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainRealTimeReqHitRateDataResponse, error) {
	resp := &DescribeDomainRealTimeReqHitRateDataResponse{}
	err := client.Request("cdn", "DescribeDomainRealTimeReqHitRateData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRealTimeReqHitRateDataResponse struct {
	RequestId goaliyun.String
	Data      DescribeDomainRealTimeReqHitRateDataReqHitRateDataModelList
}

type DescribeDomainRealTimeReqHitRateDataReqHitRateDataModel struct {
	ReqHitRate goaliyun.Float
	TimeStamp  goaliyun.String
}

type DescribeDomainRealTimeReqHitRateDataReqHitRateDataModelList []DescribeDomainRealTimeReqHitRateDataReqHitRateDataModel

func (list *DescribeDomainRealTimeReqHitRateDataReqHitRateDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeReqHitRateDataReqHitRateDataModel)
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
