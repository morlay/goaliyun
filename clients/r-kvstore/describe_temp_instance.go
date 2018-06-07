package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTempInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTempInstanceRequest) Invoke(client goaliyun.Client) (*DescribeTempInstanceResponse, error) {
	resp := &DescribeTempInstanceResponse{}
	err := client.Request("r-kvstore", "DescribeTempInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTempInstanceResponse struct {
	RequestId     goaliyun.String
	TempInstances DescribeTempInstanceTempInstanceList
}

type DescribeTempInstanceTempInstance struct {
	InstanceId     goaliyun.String
	TempInstanceId goaliyun.String
	SnapshotId     goaliyun.String
	CreateTime     goaliyun.String
	Domain         goaliyun.String
	Status         goaliyun.String
	Memory         goaliyun.Integer
	ExpireTime     goaliyun.String
}

type DescribeTempInstanceTempInstanceList []DescribeTempInstanceTempInstance

func (list *DescribeTempInstanceTempInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTempInstanceTempInstance)
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
