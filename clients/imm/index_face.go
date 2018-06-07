package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type IndexFaceRequest struct {
	SrcUris  string `position:"Query" name:"SrcUris"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	Force    string `position:"Query" name:"Force"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *IndexFaceRequest) Invoke(client goaliyun.Client) (*IndexFaceResponse, error) {
	resp := &IndexFaceResponse{}
	err := client.Request("imm", "IndexFace", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type IndexFaceResponse struct {
	RequestId      goaliyun.String
	SetId          goaliyun.String
	SuccessDetails IndexFaceSuccessDetailsItemList
	FailDetails    IndexFaceFailDetailsItemList
	SrcUris        IndexFaceSrcUriList
}

type IndexFaceSuccessDetailsItem struct {
	SrcUri  goaliyun.String
	PhotoId goaliyun.String
	Faces   IndexFaceFacesItemList
}

type IndexFaceFacesItem struct {
	FaceId        goaliyun.String
	FaceRectangle IndexFaceFaceRectangle
	FaceAttribute IndexFaceFaceAttribute
}

type IndexFaceFaceRectangle struct {
	Top    goaliyun.String
	Left   goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
}

type IndexFaceFaceAttribute struct {
	Gender      IndexFaceGender
	Age         IndexFaceAge
	HeadPose    IndexFaceHeadPose
	EyeStatus   IndexFaceEyeStatus
	Blur        IndexFaceBlur
	FaceQuality IndexFaceFaceQuality
}

type IndexFaceGender struct {
	Value goaliyun.String
}

type IndexFaceAge struct {
	Value goaliyun.Integer
}

type IndexFaceHeadPose struct {
	PitchAngle goaliyun.Float
	RollAngle  goaliyun.Float
	YawAngle   goaliyun.Float
}

type IndexFaceEyeStatus struct {
	LeftEyeStatus  IndexFaceLeftEyeStatus
	RightEyeStatus IndexFaceRightEyeStatus
}

type IndexFaceLeftEyeStatus struct {
	NormalGlassEyeOpen  goaliyun.Float
	NoGlassEyeClose     goaliyun.Float
	Occlusion           goaliyun.Float
	NoGlassEyeOpen      goaliyun.Float
	NormalGlassEyeClose goaliyun.Float
	DarkGlasses         goaliyun.Float
}

type IndexFaceRightEyeStatus struct {
	NormalGlassEyeOpen  goaliyun.Float
	NoGlassEyeClose     goaliyun.Float
	Occlusion           goaliyun.Float
	NoGlassEyeOpen      goaliyun.Float
	NormalGlassEyeClose goaliyun.Float
	DarkGlasses         goaliyun.Float
}

type IndexFaceBlur struct {
	Blurness IndexFaceBlurness
}

type IndexFaceBlurness struct {
	Balue     goaliyun.Float
	Threshold goaliyun.Float
}

type IndexFaceFaceQuality struct {
	Value     goaliyun.Float
	Threshold goaliyun.Float
}

type IndexFaceFailDetailsItem struct {
	SrcUri goaliyun.String
	Reason goaliyun.String
}

type IndexFaceSuccessDetailsItemList []IndexFaceSuccessDetailsItem

func (list *IndexFaceSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceSuccessDetailsItem)
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

type IndexFaceFailDetailsItemList []IndexFaceFailDetailsItem

func (list *IndexFaceFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceFailDetailsItem)
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

type IndexFaceSrcUriList []goaliyun.String

func (list *IndexFaceSrcUriList) UnmarshalJSON(data []byte) error {
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

type IndexFaceFacesItemList []IndexFaceFacesItem

func (list *IndexFaceFacesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceFacesItem)
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
