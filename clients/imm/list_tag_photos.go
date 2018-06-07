package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagPhotosRequest struct {
	TagName  string `position:"Query" name:"TagName"`
	MaxKeys  string `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListTagPhotosRequest) Invoke(client goaliyun.Client) (*ListTagPhotosResponse, error) {
	resp := &ListTagPhotosResponse{}
	err := client.Request("imm", "ListTagPhotos", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagPhotosResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Photos     ListTagPhotosPhotosItemList
}

type ListTagPhotosPhotosItem struct {
	SrcUri   goaliyun.String
	TagScore goaliyun.Float
}

type ListTagPhotosPhotosItemList []ListTagPhotosPhotosItem

func (list *ListTagPhotosPhotosItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagPhotosPhotosItem)
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
