package ram

import (
	"github.com/morlay/goaliyun"
)

type SetAccountAliasRequest struct {
	AccountAlias string `position:"Query" name:"AccountAlias"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SetAccountAliasRequest) Invoke(client goaliyun.Client) (*SetAccountAliasResponse, error) {
	resp := &SetAccountAliasResponse{}
	err := client.Request("ram", "SetAccountAlias", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetAccountAliasResponse struct {
	RequestId goaliyun.String
}
