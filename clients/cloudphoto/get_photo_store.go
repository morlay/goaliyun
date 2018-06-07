package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetPhotoStoreRequest struct {
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetPhotoStoreRequest) Invoke(client goaliyun.Client) (*GetPhotoStoreResponse, error) {
	resp := &GetPhotoStoreResponse{}
	err := client.Request("cloudphoto", "GetPhotoStore", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPhotoStoreResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	RequestId  goaliyun.String
	Action     goaliyun.String
	PhotoStore GetPhotoStorePhotoStore
}

type GetPhotoStorePhotoStore struct {
	Id                goaliyun.Integer
	IdStr             goaliyun.String
	Name              goaliyun.String
	Remark            goaliyun.String
	AutoCleanEnabled  bool
	AutoCleanDays     goaliyun.Integer
	DefaultQuota      goaliyun.Integer
	DefaultTrashQuota goaliyun.Integer
	Ctime             goaliyun.Integer
	Mtime             goaliyun.Integer
	Buckets           GetPhotoStoreBucketList
}

type GetPhotoStoreBucket struct {
	Name   goaliyun.String
	Region goaliyun.String
	State  goaliyun.String
	Acl    goaliyun.String
}

type GetPhotoStoreBucketList []GetPhotoStoreBucket

func (list *GetPhotoStoreBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotoStoreBucket)
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
