package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetReqAuthConfigRequest struct {
	Key1           string `position:"Query" name:"Key.1"`
	Key2           string `position:"Query" name:"Key.2"`
	AuthRemoteDesc string `position:"Query" name:"AuthRemoteDesc"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	TimeOut        string `position:"Query" name:"TimeOut"`
	AuthType       string `position:"Query" name:"AuthType"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *SetReqAuthConfigRequest) Invoke(client goaliyun.Client) (*SetReqAuthConfigResponse, error) {
	resp := &SetReqAuthConfigResponse{}
	err := client.Request("cdn", "SetReqAuthConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetReqAuthConfigResponse struct {
	RequestId goaliyun.String
}
