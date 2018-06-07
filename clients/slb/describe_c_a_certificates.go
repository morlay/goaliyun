package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCACertificatesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCACertificatesRequest) Invoke(client goaliyun.Client) (*DescribeCACertificatesResponse, error) {
	resp := &DescribeCACertificatesResponse{}
	err := client.Request("slb", "DescribeCACertificates", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCACertificatesResponse struct {
	RequestId      goaliyun.String
	CACertificates DescribeCACertificatesCACertificateList
}

type DescribeCACertificatesCACertificate struct {
	RegionId          goaliyun.String
	CACertificateId   goaliyun.String
	CACertificateName goaliyun.String
	Fingerprint       goaliyun.String
	ResourceGroupId   goaliyun.String
	CreateTime        goaliyun.String
	CreateTimeStamp   goaliyun.Integer
}

type DescribeCACertificatesCACertificateList []DescribeCACertificatesCACertificate

func (list *DescribeCACertificatesCACertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCACertificatesCACertificate)
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
