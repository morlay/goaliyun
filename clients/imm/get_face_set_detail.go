package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetFaceSetDetailRequest struct {
	Marker          string `position:"Query" name:"Marker"`
	Project         string `position:"Query" name:"Project"`
	SetId           string `position:"Query" name:"SetId"`
	ReturnAttribute string `position:"Query" name:"ReturnAttribute"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetFaceSetDetailRequest) Invoke(client goaliyun.Client) (*GetFaceSetDetailResponse, error) {
	resp := &GetFaceSetDetailResponse{}
	err := client.Request("imm", "GetFaceSetDetail", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetFaceSetDetailResponse struct {
	RequestId   goaliyun.String
	SetId       goaliyun.String
	NextMarker  goaliyun.String
	FaceDetails GetFaceSetDetailFaceDetailsItemList
}

type GetFaceSetDetailFaceDetailsItem struct {
	FaceId        goaliyun.String
	SrcUri        goaliyun.String
	PhotoId       goaliyun.String
	GroupId       goaliyun.String
	UnGroupReason goaliyun.String
	FaceRectangle GetFaceSetDetailFaceRectangle
	FaceAttribute GetFaceSetDetailFaceAttribute
}

type GetFaceSetDetailFaceRectangle struct {
	Top    goaliyun.Integer
	Left   goaliyun.Integer
	Width  goaliyun.Integer
	Height goaliyun.Integer
}

type GetFaceSetDetailFaceAttribute struct {
	Gender      GetFaceSetDetailGender
	Age         GetFaceSetDetailAge
	HeadPose    GetFaceSetDetailHeadPose
	EyeStatus   GetFaceSetDetailEyeStatus
	Blur        GetFaceSetDetailBlur
	FaceQuality GetFaceSetDetailFaceQuality
}

type GetFaceSetDetailGender struct {
	Value goaliyun.String
}

type GetFaceSetDetailAge struct {
	Value goaliyun.String
}

type GetFaceSetDetailHeadPose struct {
	PitchAngle goaliyun.Float
	RollAngle  goaliyun.Float
	YawAngle   goaliyun.Float
}

type GetFaceSetDetailEyeStatus struct {
	LeftEyeStatus  GetFaceSetDetailLeftEyeStatus
	RightEyeStatus GetFaceSetDetailRightEyeStatus
}

type GetFaceSetDetailLeftEyeStatus struct {
	NormalGlassEyeOpen  goaliyun.Float
	NoGlassEyeClose     goaliyun.Float
	Occlusion           goaliyun.Float
	NoGlassEyeOpen      goaliyun.Float
	NormalGlassEyeClose goaliyun.Float
	DarkGlasses         goaliyun.Float
}

type GetFaceSetDetailRightEyeStatus struct {
	NormalGlassEyeOpen  goaliyun.Float
	NoGlassEyeClose     goaliyun.Float
	Occlusion           goaliyun.Float
	NoGlassEyeOpen      goaliyun.Float
	NormalGlassEyeClose goaliyun.Float
	DarkGlasses         goaliyun.Float
}

type GetFaceSetDetailBlur struct {
	Blurness GetFaceSetDetailBlurness
}

type GetFaceSetDetailBlurness struct {
	Value     goaliyun.Float
	Threshold goaliyun.Float
}

type GetFaceSetDetailFaceQuality struct {
	Value     goaliyun.Float
	Threshold goaliyun.Float
}

type GetFaceSetDetailFaceDetailsItemList []GetFaceSetDetailFaceDetailsItem

func (list *GetFaceSetDetailFaceDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetFaceSetDetailFaceDetailsItem)
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
