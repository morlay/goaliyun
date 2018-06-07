package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetTagPhotosRequest struct {
	TagName  string `position:"Query" name:"TagName"`
	MaxKeys  string `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetTagPhotosRequest) Invoke(client goaliyun.Client) (*GetTagPhotosResponse, error) {
	resp := &GetTagPhotosResponse{}
	err := client.Request("imm", "GetTagPhotos", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetTagPhotosResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Photos     GetTagPhotosPhotosItemList
}

type GetTagPhotosPhotosItem struct {
	SrcUri   goaliyun.String
	TagScore goaliyun.Float
}

type GetTagPhotosPhotosItemList []GetTagPhotosPhotosItem

func (list *GetTagPhotosPhotosItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetTagPhotosPhotosItem)
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
