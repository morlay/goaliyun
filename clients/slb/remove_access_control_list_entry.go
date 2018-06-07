package slb

import (
	"github.com/morlay/goaliyun"
)

type RemoveAccessControlListEntryRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AclEntrys            string `position:"Query" name:"AclEntrys"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RemoveAccessControlListEntryRequest) Invoke(client goaliyun.Client) (*RemoveAccessControlListEntryResponse, error) {
	resp := &RemoveAccessControlListEntryResponse{}
	err := client.Request("slb", "RemoveAccessControlListEntry", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveAccessControlListEntryResponse struct {
	RequestId goaliyun.String
}
