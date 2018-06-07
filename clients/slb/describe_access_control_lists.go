package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessControlListsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AclName              string `position:"Query" name:"AclName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessControlListsRequest) Invoke(client goaliyun.Client) (*DescribeAccessControlListsResponse, error) {
	resp := &DescribeAccessControlListsResponse{}
	err := client.Request("slb", "DescribeAccessControlLists", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessControlListsResponse struct {
	RequestId goaliyun.String
	Acls      DescribeAccessControlListsAclList
}

type DescribeAccessControlListsAcl struct {
	AclId   goaliyun.String
	AclName goaliyun.String
}

type DescribeAccessControlListsAclList []DescribeAccessControlListsAcl

func (list *DescribeAccessControlListsAclList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessControlListsAcl)
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
