package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPhotoStoresRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListPhotoStoresRequest) Invoke(client goaliyun.Client) (*ListPhotoStoresResponse, error) {
	resp := &ListPhotoStoresResponse{}
	err := client.Request("cloudphoto", "ListPhotoStores", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPhotoStoresResponse struct {
	Code        goaliyun.String
	Message     goaliyun.String
	RequestId   goaliyun.String
	Action      goaliyun.String
	PhotoStores ListPhotoStoresPhotoStoreList
}

type ListPhotoStoresPhotoStore struct {
	Id               goaliyun.Integer
	IdStr            goaliyun.String
	Name             goaliyun.String
	Remark           goaliyun.String
	AutoCleanEnabled bool
	AutoCleanDays    goaliyun.Integer
	DefaultQuota     goaliyun.Integer
	Ctime            goaliyun.Integer
	Mtime            goaliyun.Integer
	Buckets          ListPhotoStoresBucketList
}

type ListPhotoStoresBucket struct {
	Name   goaliyun.String
	Region goaliyun.String
	State  goaliyun.String
}

type ListPhotoStoresPhotoStoreList []ListPhotoStoresPhotoStore

func (list *ListPhotoStoresPhotoStoreList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoStoresPhotoStore)
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

type ListPhotoStoresBucketList []ListPhotoStoresBucket

func (list *ListPhotoStoresBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoStoresBucket)
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
