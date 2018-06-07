package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRelayPushBitRateRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRelayPushBitRateRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRelayPushBitRateResponse, error) {
	resp := &DescribeLiveStreamRelayPushBitRateResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRelayPushBitRate", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRelayPushBitRateResponse struct {
	RequestId                 goaliyun.String
	RelayPushBitRateModelList DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList
}

type DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel struct {
	Time          goaliyun.String
	VedioFrame    goaliyun.String
	VedioTimstamp goaliyun.String
	AudioFrame    goaliyun.String
	AudioTimstamp goaliyun.String
	RelayDomain   goaliyun.String
}

type DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList []DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel

func (list *DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel)
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
