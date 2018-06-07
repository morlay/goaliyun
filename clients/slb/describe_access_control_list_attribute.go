package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessControlListAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AclEntryComment      string `position:"Query" name:"AclEntryComment"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessControlListAttributeRequest) Invoke(client goaliyun.Client) (*DescribeAccessControlListAttributeResponse, error) {
	resp := &DescribeAccessControlListAttributeResponse{}
	err := client.Request("slb", "DescribeAccessControlListAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessControlListAttributeResponse struct {
	RequestId        goaliyun.String
	AclId            goaliyun.String
	AclName          goaliyun.String
	AclEntrys        DescribeAccessControlListAttributeAclEntryList
	RelatedListeners DescribeAccessControlListAttributeRelatedListenerList
}

type DescribeAccessControlListAttributeAclEntry struct {
	AclEntryIP      goaliyun.String
	AclEntryComment goaliyun.String
}

type DescribeAccessControlListAttributeRelatedListener struct {
	LoadBalancerId goaliyun.String
	ListenerPort   goaliyun.Integer
	AclType        goaliyun.String
	Protocol       goaliyun.String
}

type DescribeAccessControlListAttributeAclEntryList []DescribeAccessControlListAttributeAclEntry

func (list *DescribeAccessControlListAttributeAclEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessControlListAttributeAclEntry)
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

type DescribeAccessControlListAttributeRelatedListenerList []DescribeAccessControlListAttributeRelatedListener

func (list *DescribeAccessControlListAttributeRelatedListenerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessControlListAttributeRelatedListener)
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
