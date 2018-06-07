package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPublicKeysRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListPublicKeysRequest) Invoke(client goaliyun.Client) (*ListPublicKeysResponse, error) {
	resp := &ListPublicKeysResponse{}
	err := client.Request("ram", "ListPublicKeys", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPublicKeysResponse struct {
	RequestId  goaliyun.String
	PublicKeys ListPublicKeysPublicKeyList
}

type ListPublicKeysPublicKey struct {
	PublicKeyId goaliyun.String
	Status      goaliyun.String
	CreateDate  goaliyun.String
}

type ListPublicKeysPublicKeyList []ListPublicKeysPublicKey

func (list *ListPublicKeysPublicKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPublicKeysPublicKey)
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
