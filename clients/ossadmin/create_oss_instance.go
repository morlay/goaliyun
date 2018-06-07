package ossadmin

import (
	"github.com/morlay/goaliyun"
)

type CreateOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateOssInstanceRequest) Invoke(client goaliyun.Client) (*CreateOssInstanceResponse, error) {
	resp := &CreateOssInstanceResponse{}
	err := client.Request("ossadmin", "CreateOssInstance", "2015-03-02", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateOssInstanceResponse struct {
	RequestId      goaliyun.String
	AliUid         goaliyun.Integer
	InstanceId     goaliyun.String
	InstacneStatus goaliyun.String
	InstanceName   goaliyun.String
	StartTime      goaliyun.String
	EndTime        goaliyun.String
}
