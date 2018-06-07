package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRelayPushErrorsRequest struct {
	RelayDomain   string `position:"Query" name:"RelayDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRelayPushErrorsRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRelayPushErrorsResponse, error) {
	resp := &DescribeLiveStreamRelayPushErrorsResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRelayPushErrors", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRelayPushErrorsResponse struct {
	RequestId                goaliyun.String
	RelayPushErrorsModelList DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList
}

type DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel struct {
	ErrorCode goaliyun.String
}

type DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList []DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel

func (list *DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel)
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
