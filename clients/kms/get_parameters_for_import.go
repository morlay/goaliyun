package kms

import (
	"github.com/morlay/goaliyun"
)

type GetParametersForImportRequest struct {
	KeyId             string `position:"Query" name:"KeyId"`
	STSToken          string `position:"Query" name:"STSToken"`
	WrappingAlgorithm string `position:"Query" name:"WrappingAlgorithm"`
	WrappingKeySpec   string `position:"Query" name:"WrappingKeySpec"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *GetParametersForImportRequest) Invoke(client goaliyun.Client) (*GetParametersForImportResponse, error) {
	resp := &GetParametersForImportResponse{}
	err := client.Request("kms", "GetParametersForImport", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetParametersForImportResponse struct {
	KeyId           goaliyun.String
	RequestId       goaliyun.String
	ImportToken     goaliyun.String
	PublicKey       goaliyun.String
	TokenExpireTime goaliyun.String
}
