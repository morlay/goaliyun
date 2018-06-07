package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRoomUserNumberRequest struct {
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

func (req *DescribeLiveStreamRoomUserNumberRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRoomUserNumberResponse, error) {
	resp := &DescribeLiveStreamRoomUserNumberResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRoomUserNumber", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRoomUserNumberResponse struct {
	RequestId       goaliyun.String
	TotalUserNumber goaliyun.Integer
	OnlineUserInfo  DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo struct {
	StreamUrl  goaliyun.String
	UserNumber goaliyun.Integer
	Time       goaliyun.String
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList []DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo)
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
