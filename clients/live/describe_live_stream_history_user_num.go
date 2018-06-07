package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamHistoryUserNumRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamHistoryUserNumRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamHistoryUserNumResponse, error) {
	resp := &DescribeLiveStreamHistoryUserNumResponse{}
	err := client.Request("live", "DescribeLiveStreamHistoryUserNum", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamHistoryUserNumResponse struct {
	RequestId              goaliyun.String
	LiveStreamUserNumInfos DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList
}

type DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo struct {
	StreamTime goaliyun.String
	UserNum    goaliyun.String
}

type DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList []DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo

func (list *DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo)
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
