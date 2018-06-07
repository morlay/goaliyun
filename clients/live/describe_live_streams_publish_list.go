package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamsPublishListRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamsPublishListRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamsPublishListResponse, error) {
	resp := &DescribeLiveStreamsPublishListResponse{}
	err := client.Request("live", "DescribeLiveStreamsPublishList", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamsPublishListResponse struct {
	RequestId   goaliyun.String
	PageNum     goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalNum    goaliyun.Integer
	TotalPage   goaliyun.Integer
	PublishInfo DescribeLiveStreamsPublishListLiveStreamPublishInfoList
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfo struct {
	DomainName    goaliyun.String
	AppName       goaliyun.String
	StreamName    goaliyun.String
	StreamUrl     goaliyun.String
	PublishTime   goaliyun.String
	StopTime      goaliyun.String
	PublishUrl    goaliyun.String
	ClientAddr    goaliyun.String
	EdgeNodeAddr  goaliyun.String
	PublishDomain goaliyun.String
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfoList []DescribeLiveStreamsPublishListLiveStreamPublishInfo

func (list *DescribeLiveStreamsPublishListLiveStreamPublishInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsPublishListLiveStreamPublishInfo)
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
