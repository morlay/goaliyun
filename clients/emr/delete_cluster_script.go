package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteClusterScriptRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteClusterScriptRequest) Invoke(client goaliyun.Client) (*DeleteClusterScriptResponse, error) {
	resp := &DeleteClusterScriptResponse{}
	err := client.Request("emr", "DeleteClusterScript", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteClusterScriptResponse struct {
	RequestId goaliyun.String
}
