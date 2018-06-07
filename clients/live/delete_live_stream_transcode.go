package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveStreamTranscodeRequest struct {
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveStreamTranscodeRequest) Invoke(client goaliyun.Client) (*DeleteLiveStreamTranscodeResponse, error) {
	resp := &DeleteLiveStreamTranscodeResponse{}
	err := client.Request("live", "DeleteLiveStreamTranscode", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveStreamTranscodeResponse struct {
	RequestId goaliyun.String
}
