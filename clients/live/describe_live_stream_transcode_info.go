package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamTranscodeInfoRequest struct {
	SecurityToken       string `position:"Query" name:"SecurityToken"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	DomainTranscodeName string `position:"Query" name:"DomainTranscodeName"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamTranscodeInfoRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamTranscodeInfoResponse, error) {
	resp := &DescribeLiveStreamTranscodeInfoResponse{}
	err := client.Request("live", "DescribeLiveStreamTranscodeInfo", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	RequestId           goaliyun.String
	DomainTranscodeList DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfo struct {
	TranscodeApp              goaliyun.String
	TranscodeName             goaliyun.String
	TranscodeTemplate         goaliyun.String
	CustomTranscodeParameters DescribeLiveStreamTranscodeInfoCustomTranscodeParameters
}

type DescribeLiveStreamTranscodeInfoCustomTranscodeParameters struct {
	VideoBitrate goaliyun.Integer
	FPS          goaliyun.Integer
	Height       goaliyun.Integer
	Width        goaliyun.Integer
	TemplateType goaliyun.String
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList []DescribeLiveStreamTranscodeInfoDomainTranscodeInfo

func (list *DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamTranscodeInfoDomainTranscodeInfo)
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
