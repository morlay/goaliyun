package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetUserAgentAcessRestrictionRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	UserAgent     string `position:"Query" name:"UserAgent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetUserAgentAcessRestrictionRequest) Invoke(client goaliyun.Client) (*SetUserAgentAcessRestrictionResponse, error) {
	resp := &SetUserAgentAcessRestrictionResponse{}
	err := client.Request("cdn", "SetUserAgentAcessRestriction", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetUserAgentAcessRestrictionResponse struct {
	RequestId goaliyun.String
}
