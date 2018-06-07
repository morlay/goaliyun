package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamsBlockListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int64  `position:"Query" name:"PageNum"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamsBlockListRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamsBlockListResponse, error) {
	resp := &DescribeLiveStreamsBlockListResponse{}
	err := client.Request("live", "DescribeLiveStreamsBlockList", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamsBlockListResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
	PageNum    goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalNum   goaliyun.Integer
	TotalPage  goaliyun.Integer
	StreamUrls DescribeLiveStreamsBlockListStreamUrlList
}

type DescribeLiveStreamsBlockListStreamUrlList []goaliyun.String

func (list *DescribeLiveStreamsBlockListStreamUrlList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
