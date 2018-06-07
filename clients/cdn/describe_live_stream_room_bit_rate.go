package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRoomBitRateRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRoomBitRateRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRoomBitRateResponse, error) {
	resp := &DescribeLiveStreamRoomBitRateResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRoomBitRate", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRoomBitRateResponse struct {
	RequestId                goaliyun.String
	FrameRateAndBitRateInfos DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList
}

type DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo struct {
	StreamUrl      goaliyun.String
	VideoFrameRate goaliyun.Float
	AudioFrameRate goaliyun.Float
	BitRate        goaliyun.Float
	Time           goaliyun.String
}

type DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList []DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo

func (list *DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo)
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
