package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type PlayInfoRequest struct {
	PlayDomain           string `position:"Query" name:"PlayDomain"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Formats              string `position:"Query" name:"Formats"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	HlsUriToken          string `position:"Query" name:"HlsUriToken"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Rand                 string `position:"Query" name:"Rand"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *PlayInfoRequest) Invoke(client goaliyun.Client) (*PlayInfoResponse, error) {
	resp := &PlayInfoResponse{}
	err := client.Request("mts", "PlayInfo", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PlayInfoResponse struct {
	RequestId         goaliyun.String
	PlayInfoList      PlayInfoPlayInfoList
	NotFoundCDNDomain PlayInfoNotFoundCDNDomainList
}

type PlayInfoPlayInfo struct {
	Url            goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Width          goaliyun.String
	Height         goaliyun.String
	Bitrate        goaliyun.String
	Fps            goaliyun.String
	Format         goaliyun.String
	Definition     goaliyun.String
	Encryption     goaliyun.String
	Rand           goaliyun.String
	Plaintext      goaliyun.String
	Complexity     goaliyun.String
	ActivityName   goaliyun.String
	EncryptionType goaliyun.String
	DownloadType   goaliyun.String
}

type PlayInfoPlayInfoList []PlayInfoPlayInfo

func (list *PlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayInfoPlayInfo)
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

type PlayInfoNotFoundCDNDomainList []goaliyun.String

func (list *PlayInfoNotFoundCDNDomainList) UnmarshalJSON(data []byte) error {
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
