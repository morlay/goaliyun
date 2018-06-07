package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainOnlineUserNumberRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainOnlineUserNumberRequest) Invoke(client goaliyun.Client) (*DescribeDomainOnlineUserNumberResponse, error) {
	resp := &DescribeDomainOnlineUserNumberResponse{}
	err := client.Request("cdn", "DescribeDomainOnlineUserNumber", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainOnlineUserNumberResponse struct {
	RequestId                    goaliyun.String
	LiveStreamOnlineUserNumInfos DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo struct {
	Time       goaliyun.String
	UserNumber goaliyun.Integer
}

type DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList []DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo

func (list *DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo)
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
