package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteServerCertificateRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId  string `position:"Query" name:"ServerCertificateId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteServerCertificateRequest) Invoke(client goaliyun.Client) (*DeleteServerCertificateResponse, error) {
	resp := &DeleteServerCertificateResponse{}
	err := client.Request("slb", "DeleteServerCertificate", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteServerCertificateResponse struct {
	RequestId goaliyun.String
}
