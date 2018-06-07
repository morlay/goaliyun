package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamsOnlineListRequest struct {
	StreamType    string `position:"Query" name:"StreamType"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int64  `position:"Query" name:"PageNum"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamsOnlineListRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamsOnlineListResponse, error) {
	resp := &DescribeLiveStreamsOnlineListResponse{}
	err := client.Request("live", "DescribeLiveStreamsOnlineList", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamsOnlineListResponse struct {
	RequestId  goaliyun.String
	PageNum    goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalNum   goaliyun.Integer
	TotalPage  goaliyun.Integer
	OnlineInfo DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfo struct {
	DomainName    goaliyun.String
	AppName       goaliyun.String
	StreamName    goaliyun.String
	PublishTime   goaliyun.String
	PublishUrl    goaliyun.String
	PublishDomain goaliyun.String
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList []DescribeLiveStreamsOnlineListLiveStreamOnlineInfo

func (list *DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsOnlineListLiveStreamOnlineInfo)
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
