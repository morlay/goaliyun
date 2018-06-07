package dds

import (
	"github.com/morlay/goaliyun"
)

type CreateDBInstanceRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int64  `position:"Query" name:"DBInstanceStorage"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	CouponNo              string `position:"Query" name:"CouponNo"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	NetworkType           string `position:"Query" name:"NetworkType"`
	StorageEngine         string `position:"Query" name:"StorageEngine"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	Engine                string `position:"Query" name:"Engine"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	BusinessInfo          string `position:"Query" name:"BusinessInfo"`
	Period                int64  `position:"Query" name:"Period"`
	RestoreTime           string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId       string `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	BackupId              string `position:"Query" name:"BackupId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	AccountPassword       string `position:"Query" name:"AccountPassword"`
	VpcId                 string `position:"Query" name:"VpcId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	ChargeType            string `position:"Query" name:"ChargeType"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateDBInstanceRequest) Invoke(client goaliyun.Client) (*CreateDBInstanceResponse, error) {
	resp := &CreateDBInstanceResponse{}
	err := client.Request("dds", "CreateDBInstance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDBInstanceResponse struct {
	RequestId    goaliyun.String
	OrderId      goaliyun.String
	DBInstanceId goaliyun.String
}
