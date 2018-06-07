package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAccessKeysRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListAccessKeysRequest) Invoke(client goaliyun.Client) (*ListAccessKeysResponse, error) {
	resp := &ListAccessKeysResponse{}
	err := client.Request("ram", "ListAccessKeys", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAccessKeysResponse struct {
	RequestId  goaliyun.String
	AccessKeys ListAccessKeysAccessKeyList
}

type ListAccessKeysAccessKey struct {
	AccessKeyId goaliyun.String
	Status      goaliyun.String
	CreateDate  goaliyun.String
}

type ListAccessKeysAccessKeyList []ListAccessKeysAccessKey

func (list *ListAccessKeysAccessKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAccessKeysAccessKey)
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
