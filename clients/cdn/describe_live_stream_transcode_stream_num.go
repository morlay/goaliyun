package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamTranscodeStreamNumRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamTranscodeStreamNumRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamTranscodeStreamNumResponse, error) {
	resp := &DescribeLiveStreamTranscodeStreamNumResponse{}
	err := client.Request("cdn", "DescribeLiveStreamTranscodeStreamNum", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamTranscodeStreamNumResponse struct {
	RequestId         goaliyun.String
	Total             goaliyun.Integer
	TranscodedNumber  goaliyun.Integer
	UntranscodeNumber goaliyun.Integer
}
