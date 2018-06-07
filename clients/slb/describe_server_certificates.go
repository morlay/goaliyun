package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeServerCertificatesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId  string `position:"Query" name:"ServerCertificateId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeServerCertificatesRequest) Invoke(client goaliyun.Client) (*DescribeServerCertificatesResponse, error) {
	resp := &DescribeServerCertificatesResponse{}
	err := client.Request("slb", "DescribeServerCertificates", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeServerCertificatesResponse struct {
	RequestId          goaliyun.String
	ServerCertificates DescribeServerCertificatesServerCertificateList
}

type DescribeServerCertificatesServerCertificate struct {
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

type DescribeServerCertificatesServerCertificateList []DescribeServerCertificatesServerCertificate

func (list *DescribeServerCertificatesServerCertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServerCertificatesServerCertificate)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
