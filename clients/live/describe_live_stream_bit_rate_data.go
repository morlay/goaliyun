package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamBitRateDataRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamBitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamBitRateDataResponse, error) {
	resp := &DescribeLiveStreamBitRateDataResponse{}
	err := client.Request("live", "DescribeLiveStreamBitRateData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamBitRateDataResponse struct {
	RequestId                goaliyun.String
	FrameRateAndBitRateInfos DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      goaliyun.String
	VideoFrameRate goaliyun.Float
	AudioFrameRate goaliyun.Float
	BitRate        goaliyun.Float
	Time           goaliyun.String
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo)
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
