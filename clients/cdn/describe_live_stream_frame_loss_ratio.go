package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamFrameLossRatioRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamFrameLossRatioRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamFrameLossRatioResponse, error) {
	resp := &DescribeLiveStreamFrameLossRatioResponse{}
	err := client.Request("cdn", "DescribeLiveStreamFrameLossRatio", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamFrameLossRatioResponse struct {
	RequestId           goaliyun.String
	FrameLossRatioInfos DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList
}

type DescribeLiveStreamFrameLossRatioFrameLossRatioInfo struct {
	StreamUrl      goaliyun.String
	FrameLossRatio goaliyun.Float
	Time           goaliyun.String
}

type DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList []DescribeLiveStreamFrameLossRatioFrameLossRatioInfo

func (list *DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameLossRatioFrameLossRatioInfo)
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
