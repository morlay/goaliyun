package slb

import (
	"github.com/morlay/goaliyun"
)

type UploadServerCertificateRequest struct {
	Access_key_id           string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ServerCertificate       string `position:"Query" name:"ServerCertificate"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	AliCloudCertificateName string `position:"Query" name:"AliCloudCertificateName"`
	AliCloudCertificateId   string `position:"Query" name:"AliCloudCertificateId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Tags                    string `position:"Query" name:"Tags"`
	PrivateKey              string `position:"Query" name:"PrivateKey"`
	ResourceGroupId         string `position:"Query" name:"ResourceGroupId"`
	ServerCertificateName   string `position:"Query" name:"ServerCertificateName"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *UploadServerCertificateRequest) Invoke(client goaliyun.Client) (*UploadServerCertificateResponse, error) {
	resp := &UploadServerCertificateResponse{}
	err := client.Request("slb", "UploadServerCertificate", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadServerCertificateResponse struct {
	RequestId               goaliyun.String
	ServerCertificateId     goaliyun.String
	Fingerprint             goaliyun.String
	ServerCertificateName   goaliyun.String
	RegionId                goaliyun.String
	RegionIdAlias           goaliyun.String
	AliCloudCertificateId   goaliyun.String
	AliCloudCertificateName goaliyun.String
	IsAliCloudCertificate   goaliyun.Integer
	ResourceGroupId         goaliyun.String
	CreateTime              goaliyun.String
	CreateTimeStamp         goaliyun.Integer
}
