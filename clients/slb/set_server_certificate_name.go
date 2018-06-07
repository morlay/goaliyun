package slb

import (
	"github.com/morlay/goaliyun"
)

type SetServerCertificateNameRequest struct {
	Access_key_id         string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId   string `position:"Query" name:"ServerCertificateId"`
	ServerCertificateName string `position:"Query" name:"ServerCertificateName"`
	Tags                  string `position:"Query" name:"Tags"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *SetServerCertificateNameRequest) Invoke(client goaliyun.Client) (*SetServerCertificateNameResponse, error) {
	resp := &SetServerCertificateNameResponse{}
	err := client.Request("slb", "SetServerCertificateName", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetServerCertificateNameResponse struct {
	RequestId goaliyun.String
}
