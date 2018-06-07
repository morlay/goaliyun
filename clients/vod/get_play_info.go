package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetPlayInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StreamType           string `position:"Query" name:"StreamType"`
	Formats              string `position:"Query" name:"Formats"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	VideoId              string `position:"Query" name:"VideoId"`
	PlayerVersion        string `position:"Query" name:"PlayerVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Rand                 string `position:"Query" name:"Rand"`
	ReAuthInfo           string `position:"Query" name:"ReAuthInfo"`
	OutputType           string `position:"Query" name:"OutputType"`
	Definition           string `position:"Query" name:"Definition"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetPlayInfoRequest) Invoke(client goaliyun.Client) (*GetPlayInfoResponse, error) {
	resp := &GetPlayInfoResponse{}
	err := client.Request("vod", "GetPlayInfo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPlayInfoResponse struct {
	RequestId    goaliyun.String
	PlayInfoList GetPlayInfoPlayInfoList
	VideoBase    GetPlayInfoVideoBase
}

type GetPlayInfoPlayInfo struct {
	Width            goaliyun.Integer
	Height           goaliyun.Integer
	Size             goaliyun.Integer
	PlayURL          goaliyun.String
	Bitrate          goaliyun.String
	Definition       goaliyun.String
	Duration         goaliyun.String
	Format           goaliyun.String
	Fps              goaliyun.String
	Encrypt          goaliyun.Integer
	Plaintext        goaliyun.String
	Complexity       goaliyun.String
	StreamType       goaliyun.String
	Rand             goaliyun.String
	JobId            goaliyun.String
	PreprocessStatus goaliyun.String
}

type GetPlayInfoVideoBase struct {
	CoverURL     goaliyun.String
	Duration     goaliyun.String
	Status       goaliyun.String
	Title        goaliyun.String
	VideoId      goaliyun.String
	MediaType    goaliyun.String
	CreationTime goaliyun.String
}

type GetPlayInfoPlayInfoList []GetPlayInfoPlayInfo

func (list *GetPlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPlayInfoPlayInfo)
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
