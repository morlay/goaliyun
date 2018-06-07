package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DetectFaceRequest struct {
	SrcUris  string `position:"Query" name:"SrcUris"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DetectFaceRequest) Invoke(client goaliyun.Client) (*DetectFaceResponse, error) {
	resp := &DetectFaceResponse{}
	err := client.Request("imm", "DetectFace", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetectFaceResponse struct {
	RequestId      goaliyun.String
	SuccessDetails DetectFaceSuccessDetailsItemList
	FailDetails    DetectFaceFailDetailsItemList
	SrcUris        DetectFaceSrcUriList
}

type DetectFaceSuccessDetailsItem struct {
	SrcUri  goaliyun.String
	PhotoId goaliyun.String
	Faces   DetectFaceFacesItemList
}

type DetectFaceFacesItem struct {
	FaceId        goaliyun.String
	FaceRectangle DetectFaceFaceRectangle
	FaceAttribute DetectFaceFaceAttribute
}

type DetectFaceFaceRectangle struct {
	Top    goaliyun.String
	Left   goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
}

type DetectFaceFaceAttribute struct {
	Gender      DetectFaceGender
	Age         DetectFaceAge
	HeadPose    DetectFaceHeadPose
	EyeStatus   DetectFaceEyeStatus
	Blur        DetectFaceBlur
	FaceQuality DetectFaceFaceQuality
}

type DetectFaceGender struct {
	Value goaliyun.String
}

type DetectFaceAge struct {
	Value goaliyun.Integer
}

type DetectFaceHeadPose struct {
	PitchAngle goaliyun.Float
	RollAngle  goaliyun.Float
	YawAngle   goaliyun.Float
}

type DetectFaceEyeStatus struct {
	LeftEyeStatus  DetectFaceLeftEyeStatus
	RightEyeStatus DetectFaceRightEyeStatus
}

type DetectFaceLeftEyeStatus struct {
	NormalGlassEyeOpen  goaliyun.Float
	NoGlassEyeClose     goaliyun.Float
	Occlusion           goaliyun.Float
	NoGlassEyeOpen      goaliyun.Float
	NormalGlassEyeClose goaliyun.Float
	DarkGlasses         goaliyun.Float
}

type DetectFaceRightEyeStatus struct {
	NormalGlassEyeOpen  goaliyun.Float
	NoGlassEyeClose     goaliyun.Float
	Occlusion           goaliyun.Float
	NoGlassEyeOpen      goaliyun.Float
	NormalGlassEyeClose goaliyun.Float
	DarkGlasses         goaliyun.Float
}

type DetectFaceBlur struct {
	Blurness DetectFaceBlurness
}

type DetectFaceBlurness struct {
	Value     goaliyun.Float
	Threshold goaliyun.Float
}

type DetectFaceFaceQuality struct {
	Value     goaliyun.Float
	Threshold goaliyun.Float
}

type DetectFaceFailDetailsItem struct {
	SrcUri goaliyun.String
	Reason goaliyun.String
}

type DetectFaceSuccessDetailsItemList []DetectFaceSuccessDetailsItem

func (list *DetectFaceSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceSuccessDetailsItem)
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

type DetectFaceFailDetailsItemList []DetectFaceFailDetailsItem

func (list *DetectFaceFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceFailDetailsItem)
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

type DetectFaceSrcUriList []goaliyun.String

func (list *DetectFaceSrcUriList) UnmarshalJSON(data []byte) error {
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

type DetectFaceFacesItemList []DetectFaceFacesItem

func (list *DetectFaceFacesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceFacesItem)
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
