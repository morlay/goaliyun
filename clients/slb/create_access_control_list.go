package slb

import (
	"github.com/morlay/goaliyun"
)

type CreateAccessControlListRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AclName              string `position:"Query" name:"AclName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateAccessControlListRequest) Invoke(client goaliyun.Client) (*CreateAccessControlListResponse, error) {
	resp := &CreateAccessControlListResponse{}
	err := client.Request("slb", "CreateAccessControlList", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccessControlListResponse struct {
	RequestId goaliyun.String
	AclId     goaliyun.String
}
