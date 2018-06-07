package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyInstanceNetExpireTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	ClassicExpiredDays   int64  `position:"Query" name:"ClassicExpiredDays"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceNetExpireTimeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceNetExpireTimeResponse, error) {
	resp := &ModifyInstanceNetExpireTimeResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceNetExpireTime", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceNetExpireTimeResponse struct {
	RequestId    goaliyun.String
	InstanceId   goaliyun.String
	NetInfoItems ModifyInstanceNetExpireTimeNetInfoItemList
}

type ModifyInstanceNetExpireTimeNetInfoItem struct {
	DBInstanceNetType goaliyun.String
	Port              goaliyun.String
	ExpiredTime       goaliyun.String
	ConnectionString  goaliyun.String
	IPAddress         goaliyun.String
}

type ModifyInstanceNetExpireTimeNetInfoItemList []ModifyInstanceNetExpireTimeNetInfoItem

func (list *ModifyInstanceNetExpireTimeNetInfoItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyInstanceNetExpireTimeNetInfoItem)
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
