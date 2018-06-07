package dds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    string `position:"Query" name:"DBInstanceStorage"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	OrderType            string `position:"Query" name:"OrderType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceSpecResponse, error) {
	resp := &ModifyDBInstanceSpecResponse{}
	err := client.Request("dds", "ModifyDBInstanceSpec", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
