package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteCACertificateRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCACertificateRequest) Invoke(client goaliyun.Client) (*DeleteCACertificateResponse, error) {
	resp := &DeleteCACertificateResponse{}
	err := client.Request("slb", "DeleteCACertificate", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCACertificateResponse struct {
	RequestId goaliyun.String
}
