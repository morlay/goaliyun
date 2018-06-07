package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamPushErrorsRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamPushErrorsRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamPushErrorsResponse, error) {
	resp := &DescribeLiveStreamPushErrorsResponse{}
	err := client.Request("cdn", "DescribeLiveStreamPushErrors", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamPushErrorsResponse struct {
	RequestId           goaliyun.String
	PushErrorsModelList DescribeLiveStreamPushErrorsPushErrorsModelList
}

type DescribeLiveStreamPushErrorsPushErrorsModel struct {
	ErrorCode goaliyun.String
}

type DescribeLiveStreamPushErrorsPushErrorsModelList []DescribeLiveStreamPushErrorsPushErrorsModel

func (list *DescribeLiveStreamPushErrorsPushErrorsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamPushErrorsPushErrorsModel)
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
