package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type ClickToDialRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecordFlag           string `position:"Query" name:"RecordFlag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CallerShowNumber     string `position:"Query" name:"CallerShowNumber"`
	SessionTimeout       int64  `position:"Query" name:"SessionTimeout"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	OutId                string `position:"Query" name:"OutId"`
	AsrFlag              string `position:"Query" name:"AsrFlag"`
	AsrModelId           string `position:"Query" name:"AsrModelId"`
	CallerNumber         string `position:"Query" name:"CallerNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ClickToDialRequest) Invoke(client goaliyun.Client) (*ClickToDialResponse, error) {
	resp := &ClickToDialResponse{}
	err := client.Request("dyvmsapi", "ClickToDial", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ClickToDialResponse struct {
	RequestId goaliyun.String
	CallId    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
