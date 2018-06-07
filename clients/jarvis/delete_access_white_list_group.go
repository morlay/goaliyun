package jarvis

import (
	"github.com/morlay/goaliyun"
)

type DeleteAccessWhiteListGroupRequest struct {
	GroupIdList string `position:"Query" name:"GroupIdList"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	Lang        string `position:"Query" name:"Lang"`
	SourceCode  string `position:"Query" name:"SourceCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteAccessWhiteListGroupRequest) Invoke(client goaliyun.Client) (*DeleteAccessWhiteListGroupResponse, error) {
	resp := &DeleteAccessWhiteListGroupResponse{}
	err := client.Request("jarvis", "DeleteAccessWhiteListGroup", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAccessWhiteListGroupResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
