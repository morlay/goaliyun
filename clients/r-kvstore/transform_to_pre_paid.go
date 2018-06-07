package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type TransformToPrePaidRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *TransformToPrePaidRequest) Invoke(client goaliyun.Client) (*TransformToPrePaidResponse, error) {
	resp := &TransformToPrePaidResponse{}
	err := client.Request("r-kvstore", "TransformToPrePaid", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TransformToPrePaidResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
	EndTime   goaliyun.String
}
