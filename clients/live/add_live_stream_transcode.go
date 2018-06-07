package live

import (
	"github.com/morlay/goaliyun"
)

type AddLiveStreamTranscodeRequest struct {
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddLiveStreamTranscodeRequest) Invoke(client goaliyun.Client) (*AddLiveStreamTranscodeResponse, error) {
	resp := &AddLiveStreamTranscodeResponse{}
	err := client.Request("live", "AddLiveStreamTranscode", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveStreamTranscodeResponse struct {
	RequestId goaliyun.String
}
