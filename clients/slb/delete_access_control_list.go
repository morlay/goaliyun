package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteAccessControlListRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteAccessControlListRequest) Invoke(client goaliyun.Client) (*DeleteAccessControlListResponse, error) {
	resp := &DeleteAccessControlListResponse{}
	err := client.Request("slb", "DeleteAccessControlList", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAccessControlListResponse struct {
	RequestId goaliyun.String
}
