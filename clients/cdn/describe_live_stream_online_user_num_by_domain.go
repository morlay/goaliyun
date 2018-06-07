package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamOnlineUserNumByDomainRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamOnlineUserNumByDomainRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamOnlineUserNumByDomainResponse, error) {
	resp := &DescribeLiveStreamOnlineUserNumByDomainResponse{}
	err := client.Request("cdn", "DescribeLiveStreamOnlineUserNumByDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamOnlineUserNumByDomainResponse struct {
	RequestId       goaliyun.String
	TotalUserNumber goaliyun.Integer
	Count           goaliyun.Integer
	PageNumber      goaliyun.Integer
	PageSize        goaliyun.Integer
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  goaliyun.String
	UserNumber goaliyun.Integer
	Time       goaliyun.String
}

type DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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
