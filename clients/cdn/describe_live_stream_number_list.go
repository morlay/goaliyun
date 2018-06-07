package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamNumberListRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamNumberListRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamNumberListResponse, error) {
	resp := &DescribeLiveStreamNumberListResponse{}
	err := client.Request("cdn", "DescribeLiveStreamNumberList", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamNumberListResponse struct {
	RequestId         goaliyun.String
	DomainName        goaliyun.String
	StreamNumberInfos DescribeLiveStreamNumberListStreamNumberInfoList
}

type DescribeLiveStreamNumberListStreamNumberInfo struct {
	Number goaliyun.Integer
	Time   goaliyun.String
}

type DescribeLiveStreamNumberListStreamNumberInfoList []DescribeLiveStreamNumberListStreamNumberInfo

func (list *DescribeLiveStreamNumberListStreamNumberInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamNumberListStreamNumberInfo)
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
