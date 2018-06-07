package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamHlsOnlineUserNumByDomainRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamHlsOnlineUserNumByDomainRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamHlsOnlineUserNumByDomainResponse, error) {
	resp := &DescribeLiveStreamHlsOnlineUserNumByDomainResponse{}
	err := client.Request("cdn", "DescribeLiveStreamHlsOnlineUserNumByDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamHlsOnlineUserNumByDomainResponse struct {
	RequestId       goaliyun.String
	TotalUserNumber goaliyun.Integer
	Count           goaliyun.Integer
	PageNumber      goaliyun.Integer
	PageSize        goaliyun.Integer
	OnlineUserInfo  DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  goaliyun.String
	UserNumber goaliyun.Integer
	Time       goaliyun.String
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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
