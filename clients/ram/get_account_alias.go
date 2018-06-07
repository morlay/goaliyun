package ram

import (
	"github.com/morlay/goaliyun"
)

type GetAccountAliasRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAccountAliasRequest) Invoke(client goaliyun.Client) (*GetAccountAliasResponse, error) {
	resp := &GetAccountAliasResponse{}
	err := client.Request("ram", "GetAccountAlias", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAccountAliasResponse struct {
	RequestId    goaliyun.String
	AccountAlias goaliyun.String
}
