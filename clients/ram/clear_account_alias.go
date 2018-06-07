package ram

import (
	"github.com/morlay/goaliyun"
)

type ClearAccountAliasRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ClearAccountAliasRequest) Invoke(client goaliyun.Client) (*ClearAccountAliasResponse, error) {
	resp := &ClearAccountAliasResponse{}
	err := client.Request("ram", "ClearAccountAlias", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ClearAccountAliasResponse struct {
	RequestId goaliyun.String
}
