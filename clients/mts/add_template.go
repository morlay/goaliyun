package mts

import (
	"github.com/morlay/goaliyun"
)

type AddTemplateRequest struct {
	Container            string `position:"Query" name:"Container"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	TransConfig          string `position:"Query" name:"TransConfig"`
	MuxConfig            string `position:"Query" name:"MuxConfig"`
	Video                string `position:"Query" name:"Video"`
	Audio                string `position:"Query" name:"Audio"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddTemplateRequest) Invoke(client goaliyun.Client) (*AddTemplateResponse, error) {
	resp := &AddTemplateResponse{}
	err := client.Request("mts", "AddTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddTemplateResponse struct {
	RequestId goaliyun.String
	Template  AddTemplateTemplate
}

type AddTemplateTemplate struct {
	Id          goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Container   AddTemplateContainer
	Video       AddTemplateVideo
	Audio       AddTemplateAudio
	TransConfig AddTemplateTransConfig
	MuxConfig   AddTemplateMuxConfig
}

type AddTemplateContainer struct {
	Format goaliyun.String
}

type AddTemplateVideo struct {
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
	BitrateBnd AddTemplateBitrateBnd
}

type AddTemplateBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type AddTemplateAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Remove     goaliyun.String
	Volume     AddTemplateVolume
}

type AddTemplateVolume struct {
	Level  goaliyun.String
	Method goaliyun.String
}

type AddTemplateTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type AddTemplateMuxConfig struct {
	Segment AddTemplateSegment
	Gif     AddTemplateGif
}

type AddTemplateSegment struct {
	Duration goaliyun.String
}

type AddTemplateGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}
