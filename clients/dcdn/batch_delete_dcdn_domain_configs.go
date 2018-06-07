package dcdn

import (
	"github.com/morlay/goaliyun"
)

type BatchDeleteDcdnDomainConfigsRequest struct {
	FunctionNames string `position:"Query" name:"FunctionNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *BatchDeleteDcdnDomainConfigsRequest) Invoke(client goaliyun.Client) (*BatchDeleteDcdnDomainConfigsResponse, error) {
	resp := &BatchDeleteDcdnDomainConfigsResponse{}
	err := client.Request("dcdn", "BatchDeleteDcdnDomainConfigs", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchDeleteDcdnDomainConfigsResponse struct {
	RequestId goaliyun.String
}
