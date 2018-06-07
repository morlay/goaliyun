package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRelayPushDataRequest struct {
	RelayDomain   string `position:"Query" name:"RelayDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRelayPushDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRelayPushDataResponse, error) {
	resp := &DescribeLiveStreamRelayPushDataResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRelayPushData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRelayPushDataResponse struct {
	RequestId                goaliyun.String
	RelayPushDetailModelList DescribeLiveStreamRelayPushDataRelayPushDetailModelList
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModel struct {
	Time          goaliyun.String
	Stream        goaliyun.String
	FrameRate     goaliyun.Float
	BitRate       goaliyun.Float
	FrameLossRate goaliyun.Float
	ServerAddr    goaliyun.String
	ClientAddr    goaliyun.String
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModelList []DescribeLiveStreamRelayPushDataRelayPushDetailModel

func (list *DescribeLiveStreamRelayPushDataRelayPushDetailModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushDataRelayPushDetailModel)
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
