package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddLiveStreamTranscodeRequest struct {
	Template      string `position:"Query" name:"Template"`
	App           string `position:"Query" name:"App"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	Domain        string `position:"Query" name:"Domain"`
	Record        string `position:"Query" name:"Record"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Snapshot      string `position:"Query" name:"Snapshot"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddLiveStreamTranscodeRequest) Invoke(client goaliyun.Client) (*AddLiveStreamTranscodeResponse, error) {
	resp := &AddLiveStreamTranscodeResponse{}
	err := client.Request("cdn", "AddLiveStreamTranscode", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveStreamTranscodeResponse struct {
	RequestId goaliyun.String
}
