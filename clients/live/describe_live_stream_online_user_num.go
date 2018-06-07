package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamOnlineUserNumRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamOnlineUserNumRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamOnlineUserNumResponse, error) {
	resp := &DescribeLiveStreamOnlineUserNumResponse{}
	err := client.Request("live", "DescribeLiveStreamOnlineUserNum", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamOnlineUserNumResponse struct {
	RequestId       goaliyun.String
	TotalUserNumber goaliyun.Integer
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo struct {
	StreamUrl  goaliyun.String
	UserNumber goaliyun.Integer
	Time       goaliyun.String
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList []DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo)
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
