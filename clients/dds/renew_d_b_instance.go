package dds

import (
	"github.com/morlay/goaliyun"
)

type RenewDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RenewDBInstanceRequest) Invoke(client goaliyun.Client) (*RenewDBInstanceResponse, error) {
	resp := &RenewDBInstanceResponse{}
	err := client.Request("dds", "RenewDBInstance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewDBInstanceResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
