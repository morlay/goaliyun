package slb

import (
	"github.com/morlay/goaliyun"
)

type UploadCACertificateRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CACertificate        string `position:"Query" name:"CACertificate"`
	CACertificateName    string `position:"Query" name:"CACertificateName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UploadCACertificateRequest) Invoke(client goaliyun.Client) (*UploadCACertificateResponse, error) {
	resp := &UploadCACertificateResponse{}
	err := client.Request("slb", "UploadCACertificate", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadCACertificateResponse struct {
	RequestId         goaliyun.String
	CACertificateId   goaliyun.String
	CACertificateName goaliyun.String
	Fingerprint       goaliyun.String
	ResourceGroupId   goaliyun.String
	CreateTime        goaliyun.String
	CreateTimeStamp   goaliyun.Integer
}
