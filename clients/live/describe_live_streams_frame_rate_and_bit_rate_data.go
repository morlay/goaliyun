package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamsFrameRateAndBitRateDataRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamsFrameRateAndBitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamsFrameRateAndBitRateDataResponse, error) {
	resp := &DescribeLiveStreamsFrameRateAndBitRateDataResponse{}
	err := client.Request("live", "DescribeLiveStreamsFrameRateAndBitRateData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamsFrameRateAndBitRateDataResponse struct {
	RequestId                goaliyun.String
	FrameRateAndBitRateInfos DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      goaliyun.String
	VideoFrameRate goaliyun.Float
	AudioFrameRate goaliyun.Float
	BitRate        goaliyun.Float
	Time           goaliyun.String
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo)
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
