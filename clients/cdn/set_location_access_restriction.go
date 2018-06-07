package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetLocationAccessRestrictionRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Location      string `position:"Query" name:"Location"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetLocationAccessRestrictionRequest) Invoke(client goaliyun.Client) (*SetLocationAccessRestrictionResponse, error) {
	resp := &SetLocationAccessRestrictionResponse{}
	err := client.Request("cdn", "SetLocationAccessRestriction", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLocationAccessRestrictionResponse struct {
	RequestId goaliyun.String
}
