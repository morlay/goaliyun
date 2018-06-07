package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTemplateListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateIds          string `position:"Query" name:"TemplateIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryTemplateListRequest) Invoke(client goaliyun.Client) (*QueryTemplateListResponse, error) {
	resp := &QueryTemplateListResponse{}
	err := client.Request("mts", "QueryTemplateList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTemplateListResponse struct {
	RequestId    goaliyun.String
	TemplateList QueryTemplateListTemplateList
	NonExistTids QueryTemplateListNonExistTidList
}

type QueryTemplateListTemplate struct {
	Id          goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Container   QueryTemplateListContainer
	Video       QueryTemplateListVideo
	Audio       QueryTemplateListAudio
	TransConfig QueryTemplateListTransConfig
	MuxConfig   QueryTemplateListMuxConfig
}

type QueryTemplateListContainer struct {
	Format goaliyun.String
}

type QueryTemplateListVideo struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Bitrate    goaliyun.String
	Crf        goaliyun.String
	Width      goaliyun.String
	Height     goaliyun.String
	Fps        goaliyun.String
	Gop        goaliyun.String
	Preset     goaliyun.String
	ScanMode   goaliyun.String
	Bufsize    goaliyun.String
	Maxrate    goaliyun.String
	PixFmt     goaliyun.String
	Degrain    goaliyun.String
	Qscale     goaliyun.String
	Remove     goaliyun.String
	Crop       goaliyun.String
	Pad        goaliyun.String
	MaxFps     goaliyun.String
	BitrateBnd QueryTemplateListBitrateBnd
}

type QueryTemplateListBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type QueryTemplateListAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Remove     goaliyun.String
}

type QueryTemplateListTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type QueryTemplateListMuxConfig struct {
	Segment QueryTemplateListSegment
	Gif     QueryTemplateListGif
}

type QueryTemplateListSegment struct {
	Duration goaliyun.String
}

type QueryTemplateListGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type QueryTemplateListTemplateList []QueryTemplateListTemplate

func (list *QueryTemplateListTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTemplateListTemplate)
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

type QueryTemplateListNonExistTidList []goaliyun.String

func (list *QueryTemplateListNonExistTidList) UnmarshalJSON(data []byte) error {
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
