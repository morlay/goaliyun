package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SearchTemplateRequest) Invoke(client goaliyun.Client) (*SearchTemplateResponse, error) {
	resp := &SearchTemplateResponse{}
	err := client.Request("mts", "SearchTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchTemplateResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	TemplateList SearchTemplateTemplateList
}

type SearchTemplateTemplate struct {
	Id          goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Container   SearchTemplateContainer
	Video       SearchTemplateVideo
	Audio       SearchTemplateAudio
	TransConfig SearchTemplateTransConfig
	MuxConfig   SearchTemplateMuxConfig
}

type SearchTemplateContainer struct {
	Format goaliyun.String
}

type SearchTemplateVideo struct {
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
	BitrateBnd SearchTemplateBitrateBnd
}

type SearchTemplateBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type SearchTemplateAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Remove     goaliyun.String
}

type SearchTemplateTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type SearchTemplateMuxConfig struct {
	Segment SearchTemplateSegment
	Gif     SearchTemplateGif
}

type SearchTemplateSegment struct {
	Duration goaliyun.String
}

type SearchTemplateGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type SearchTemplateTemplateList []SearchTemplateTemplate

func (list *SearchTemplateTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchTemplateTemplate)
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
