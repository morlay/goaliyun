package emr

import (
	"github.com/morlay/goaliyun"
)

type RelateStackVersionRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	StackVersion    string `position:"Query" name:"StackVersion"`
	UserId          string `position:"Query" name:"UserId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RelateStackVersionRequest) Invoke(client goaliyun.Client) (*RelateStackVersionResponse, error) {
	resp := &RelateStackVersionResponse{}
	err := client.Request("emr", "RelateStackVersion", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RelateStackVersionResponse struct {
	RequestId goaliyun.String
	Success   bool
}
