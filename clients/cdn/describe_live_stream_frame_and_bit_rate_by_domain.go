package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamFrameAndBitRateByDomainRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamFrameAndBitRateByDomainRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamFrameAndBitRateByDomainResponse, error) {
	resp := &DescribeLiveStreamFrameAndBitRateByDomainResponse{}
	err := client.Request("cdn", "DescribeLiveStreamFrameAndBitRateByDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamFrameAndBitRateByDomainResponse struct {
	RequestId                goaliyun.String
	Count                    goaliyun.Integer
	PageNumber               goaliyun.Integer
	PageSize                 goaliyun.Integer
	FrameRateAndBitRateInfos DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo struct {
	StreamUrl      goaliyun.String
	VideoFrameRate goaliyun.Float
	AudioFrameRate goaliyun.Float
	BitRate        goaliyun.Float
	Time           goaliyun.String
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList []DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo

func (list *DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo)
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
