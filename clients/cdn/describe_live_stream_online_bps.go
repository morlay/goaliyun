package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamOnlineBpsRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamOnlineBpsRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamOnlineBpsResponse, error) {
	resp := &DescribeLiveStreamOnlineBpsResponse{}
	err := client.Request("cdn", "DescribeLiveStreamOnlineBps", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamOnlineBpsResponse struct {
	RequestId                goaliyun.String
	TotalUserNumber          goaliyun.Integer
	FlvBps                   goaliyun.Float
	HlsBps                   goaliyun.Float
	LiveStreamOnlineBpsInfos DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList
}

type DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo struct {
	StreamUrl goaliyun.String
	UpBps     goaliyun.Float
	DownBps   goaliyun.Float
	Time      goaliyun.String
}

type DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList []DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo

func (list *DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo)
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
