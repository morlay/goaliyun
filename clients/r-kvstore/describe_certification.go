package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DescribeCertificationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Parameters           string `position:"Query" name:"Parameters"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCertificationRequest) Invoke(client goaliyun.Client) (*DescribeCertificationResponse, error) {
	resp := &DescribeCertificationResponse{}
	err := client.Request("r-kvstore", "DescribeCertification", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCertificationResponse struct {
	RequestId       goaliyun.String
	NoCertification bool
}
