package kms

import (
	"github.com/morlay/goaliyun"
)

type DeleteKeyMaterialRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteKeyMaterialRequest) Invoke(client goaliyun.Client) (*DeleteKeyMaterialResponse, error) {
	resp := &DeleteKeyMaterialResponse{}
	err := client.Request("kms", "DeleteKeyMaterial", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteKeyMaterialResponse struct {
	RequestId goaliyun.String
}
