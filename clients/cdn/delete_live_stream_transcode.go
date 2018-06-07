package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveStreamTranscodeRequest struct {
	Template      string `position:"Query" name:"Template"`
	App           string `position:"Query" name:"App"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveStreamTranscodeRequest) Invoke(client goaliyun.Client) (*DeleteLiveStreamTranscodeResponse, error) {
	resp := &DeleteLiveStreamTranscodeResponse{}
	err := client.Request("cdn", "DeleteLiveStreamTranscode", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveStreamTranscodeResponse struct {
	RequestId goaliyun.String
}
