package jarvis

import (
	"github.com/morlay/goaliyun"
)

type DeleteUidWhiteListGroupRequest struct {
	GroupIdList string `position:"Query" name:"GroupIdList"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	Lang        string `position:"Query" name:"Lang"`
	SourceCode  string `position:"Query" name:"SourceCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteUidWhiteListGroupRequest) Invoke(client goaliyun.Client) (*DeleteUidWhiteListGroupResponse, error) {
	resp := &DeleteUidWhiteListGroupResponse{}
	err := client.Request("jarvis", "DeleteUidWhiteListGroup", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteUidWhiteListGroupResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
