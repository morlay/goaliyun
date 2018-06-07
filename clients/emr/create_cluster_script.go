package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateClusterScriptRequest struct {
	Args            string `position:"Query" name:"Args"`
	Path            string `position:"Query" name:"Path"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	NodeIdList      string `position:"Query" name:"NodeIdList"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateClusterScriptRequest) Invoke(client goaliyun.Client) (*CreateClusterScriptResponse, error) {
	resp := &CreateClusterScriptResponse{}
	err := client.Request("emr", "CreateClusterScript", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateClusterScriptResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
