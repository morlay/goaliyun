package slb

import (
	"github.com/morlay/goaliyun"
)

type SetCACertificateNameRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CACertificateName    string `position:"Query" name:"CACertificateName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetCACertificateNameRequest) Invoke(client goaliyun.Client) (*SetCACertificateNameResponse, error) {
	resp := &SetCACertificateNameResponse{}
	err := client.Request("slb", "SetCACertificateName", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetCACertificateNameResponse struct {
	RequestId goaliyun.String
}
