package cloudauth

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitMaterialsRequest struct {
	ResourceOwnerId int64                        `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string                       `position:"Query" name:"SourceIp"`
	Materials       *SubmitMaterialsMaterialList `position:"Query" type:"Repeated" name:"Material"`
	VerifyToken     string                       `position:"Query" name:"VerifyToken"`
	RegionId        string                       `position:"Query" name:"RegionId"`
}

func (req *SubmitMaterialsRequest) Invoke(client goaliyun.Client) (*SubmitMaterialsResponse, error) {
	resp := &SubmitMaterialsResponse{}
	err := client.Request("cloudauth", "SubmitMaterials", "2018-05-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitMaterialsMaterial struct {
	MaterialType string `name:"MaterialType"`
	Value        string `name:"Value"`
}

type SubmitMaterialsResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      SubmitMaterialsData
}

type SubmitMaterialsData struct {
	VerifyStatus SubmitMaterialsVerifyStatus
}

type SubmitMaterialsVerifyStatus struct {
	StatusCode      goaliyun.Integer
	TrustedScore    goaliyun.Float
	SimilarityScore goaliyun.Float
}

type SubmitMaterialsMaterialList []SubmitMaterialsMaterial

func (list *SubmitMaterialsMaterialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMaterialsMaterial)
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
