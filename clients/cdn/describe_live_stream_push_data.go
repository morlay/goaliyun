package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamPushDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamPushDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamPushDataResponse, error) {
	resp := &DescribeLiveStreamPushDataResponse{}
	err := client.Request("cdn", "DescribeLiveStreamPushData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamPushDataResponse struct {
	RequestId           goaliyun.String
	PushStreamModelList DescribeLiveStreamPushDataPushStreamModelList
}

type DescribeLiveStreamPushDataPushStreamModel struct {
	Time          goaliyun.String
	Stream        goaliyun.String
	FrameRate     goaliyun.Float
	BitRate       goaliyun.Float
	FrameLossRate goaliyun.Float
	ServerAddr    goaliyun.String
	ClientAddr    goaliyun.String
}

type DescribeLiveStreamPushDataPushStreamModelList []DescribeLiveStreamPushDataPushStreamModel

func (list *DescribeLiveStreamPushDataPushStreamModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamPushDataPushStreamModel)
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
