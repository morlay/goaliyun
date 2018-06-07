package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type CreateStaticVerificationRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	ReplicaId             string `position:"Query" name:"ReplicaId"`
	DestinationInstanceId string `position:"Query" name:"DestinationInstanceId"`
	SourceInstanceId      string `position:"Query" name:"SourceInstanceId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateStaticVerificationRequest) Invoke(client goaliyun.Client) (*CreateStaticVerificationResponse, error) {
	resp := &CreateStaticVerificationResponse{}
	err := client.Request("r-kvstore", "CreateStaticVerification", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateStaticVerificationResponse struct {
	RequestId goaliyun.String
}
