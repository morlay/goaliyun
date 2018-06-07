package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetEditingProjectMaterialsRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Type                 string `position:"Query" name:"Type"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetEditingProjectMaterialsRequest) Invoke(client goaliyun.Client) (*GetEditingProjectMaterialsResponse, error) {
	resp := &GetEditingProjectMaterialsResponse{}
	err := client.Request("vod", "GetEditingProjectMaterials", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetEditingProjectMaterialsResponse struct {
	RequestId    goaliyun.String
	MaterialList GetEditingProjectMaterialsMaterialList
}

type GetEditingProjectMaterialsMaterial struct {
	MaterialId   goaliyun.String
	Title        goaliyun.String
	Tags         goaliyun.String
	Status       goaliyun.String
	Size         goaliyun.Integer
	Duration     goaliyun.Float
	Description  goaliyun.String
	CreationTime goaliyun.String
	ModifiedTime goaliyun.String
	CoverURL     goaliyun.String
	CateId       goaliyun.Integer
	CateName     goaliyun.String
	Source       goaliyun.String
	Snapshots    GetEditingProjectMaterialsSnapshotList
	Sprites      GetEditingProjectMaterialsSpriteList
}

type GetEditingProjectMaterialsMaterialList []GetEditingProjectMaterialsMaterial

func (list *GetEditingProjectMaterialsMaterialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEditingProjectMaterialsMaterial)
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

type GetEditingProjectMaterialsSnapshotList []goaliyun.String

func (list *GetEditingProjectMaterialsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type GetEditingProjectMaterialsSpriteList []goaliyun.String

func (list *GetEditingProjectMaterialsSpriteList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
