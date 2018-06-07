package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamLimitInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	LimitDomain   string `position:"Query" name:"LimitDomain"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamLimitInfoRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamLimitInfoResponse, error) {
	resp := &DescribeLiveStreamLimitInfoResponse{}
	err := client.Request("cdn", "DescribeLiveStreamLimitInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamLimitInfoResponse struct {
	RequestId      goaliyun.String
	UserLimitLists DescribeLiveStreamLimitInfoUserLimitModeList
}

type DescribeLiveStreamLimitInfoUserLimitMode struct {
	LimitDomain       goaliyun.String
	LimitNum          goaliyun.String
	LimitTranscodeNum goaliyun.String
}

type DescribeLiveStreamLimitInfoUserLimitModeList []DescribeLiveStreamLimitInfoUserLimitMode

func (list *DescribeLiveStreamLimitInfoUserLimitModeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamLimitInfoUserLimitMode)
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
