package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainHttpsDataRequest struct {
	DomainType    string `position:"Query" name:"DomainType"`
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	Cls           string `position:"Query" name:"Cls"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainHttpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainHttpsDataResponse, error) {
	resp := &DescribeDomainHttpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainHttpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainHttpsDataResponse struct {
	RequestId            goaliyun.String
	DomainNames          goaliyun.String
	DataInterval         goaliyun.String
	HttpsStatisticsInfos DescribeDomainHttpsDataHttpsStatisticsInfoList
}

type DescribeDomainHttpsDataHttpsStatisticsInfo struct {
	Time               goaliyun.String
	L1HttpsBps         goaliyun.Float
	L1HttpsInnerBps    goaliyun.Float
	L1HttpsOutBps      goaliyun.Float
	L1HttpsQps         goaliyun.Integer
	L1HttpsInnerQps    goaliyun.Integer
	L1HttpsOutQps      goaliyun.Integer
	L1HttpsTtraf       goaliyun.Integer
	HttpsSrcBps        goaliyun.Integer
	HttpsSrcTraf       goaliyun.Integer
	L1HttpsInnerTraf   goaliyun.Integer
	L1HttpsOutTraf     goaliyun.Integer
	HttpsByteHitRate   goaliyun.Float
	HttpsReqHitRate    goaliyun.Float
	L1HttpsHitRate     goaliyun.Float
	L1HttpsInner_acc   goaliyun.Float
	L1HttpsOut_acc     goaliyun.Float
	L1HttpsTacc        goaliyun.Float
	L1DyHttpsBps       goaliyun.Float
	L1DyHttpsInnerBps  goaliyun.Float
	L1DyHttpsOutBps    goaliyun.Float
	L1StHttpsBps       goaliyun.Float
	L1StHttpsInnerBps  goaliyun.Float
	L1StHttpsOutBps    goaliyun.Float
	L1DyHttpsTraf      goaliyun.Float
	L1DyHttpsInnerTraf goaliyun.Float
	L1DyHttpsOutTraf   goaliyun.Float
	L1StHttpsTraf      goaliyun.Float
	L1StHttpsInnerTraf goaliyun.Float
	L1StHttpsOutTraf   goaliyun.Float
	L1DyHttpsQps       goaliyun.Float
	L1DyHttpsInnerQps  goaliyun.Float
	L1DyHttpsOutQps    goaliyun.Float
	L1StHttpsQps       goaliyun.Float
	L1StHttpsInnerQps  goaliyun.Float
	L1StHttpsOutQps    goaliyun.Float
	L1DyHttpsAcc       goaliyun.Float
	L1DyHttpsInnerAcc  goaliyun.Float
	L1DyHttpsOutAcc    goaliyun.Float
	L1StHttpsAcc       goaliyun.Float
	L1StHttpsInnerAcc  goaliyun.Float
	L1StHttpsOutAcc    goaliyun.Float
}

type DescribeDomainHttpsDataHttpsStatisticsInfoList []DescribeDomainHttpsDataHttpsStatisticsInfo

func (list *DescribeDomainHttpsDataHttpsStatisticsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpsDataHttpsStatisticsInfo)
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
