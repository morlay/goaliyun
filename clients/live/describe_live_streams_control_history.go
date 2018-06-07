package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamsControlHistoryRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamsControlHistoryRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamsControlHistoryResponse, error) {
	resp := &DescribeLiveStreamsControlHistoryResponse{}
	err := client.Request("live", "DescribeLiveStreamsControlHistory", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamsControlHistoryResponse struct {
	RequestId   goaliyun.String
	ControlInfo DescribeLiveStreamsControlHistoryLiveStreamControlInfoList
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfo struct {
	StreamName goaliyun.String
	ClientIP   goaliyun.String
	Action     goaliyun.String
	TimeStamp  goaliyun.String
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfoList []DescribeLiveStreamsControlHistoryLiveStreamControlInfo

func (list *DescribeLiveStreamsControlHistoryLiveStreamControlInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsControlHistoryLiveStreamControlInfo)
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
