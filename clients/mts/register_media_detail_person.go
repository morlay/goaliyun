package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RegisterMediaDetailPersonRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Images               string `position:"Query" name:"Images"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Category             string `position:"Query" name:"Category"`
	PersonName           string `position:"Query" name:"PersonName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RegisterMediaDetailPersonRequest) Invoke(client goaliyun.Client) (*RegisterMediaDetailPersonResponse, error) {
	resp := &RegisterMediaDetailPersonResponse{}
	err := client.Request("mts", "RegisterMediaDetailPerson", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegisterMediaDetailPersonResponse struct {
	RequestId            goaliyun.String
	RegisteredPersonages RegisterMediaDetailPersonRegisteredPersonageList
	FailedImages         RegisterMediaDetailPersonFailedImageList
}

type RegisterMediaDetailPersonRegisteredPersonage struct {
	PersonName goaliyun.String
	FaceId     goaliyun.String
	Target     goaliyun.String
	Quality    goaliyun.String
	Gender     goaliyun.String
	ImageId    goaliyun.String
	ImageFile  RegisterMediaDetailPersonImageFile
}

type RegisterMediaDetailPersonImageFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type RegisterMediaDetailPersonFailedImage struct {
	Code       goaliyun.String
	Success    goaliyun.String
	ImageFile1 RegisterMediaDetailPersonImageFile1
}

type RegisterMediaDetailPersonImageFile1 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type RegisterMediaDetailPersonRegisteredPersonageList []RegisterMediaDetailPersonRegisteredPersonage

func (list *RegisterMediaDetailPersonRegisteredPersonageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RegisterMediaDetailPersonRegisteredPersonage)
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

type RegisterMediaDetailPersonFailedImageList []RegisterMediaDetailPersonFailedImage

func (list *RegisterMediaDetailPersonFailedImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RegisterMediaDetailPersonFailedImage)
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
