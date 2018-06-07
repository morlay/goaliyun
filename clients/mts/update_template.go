package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateTemplateRequest struct {
	Container            string `position:"Query" name:"Container"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MuxConfig            string `position:"Query" name:"MuxConfig"`
	Video                string `position:"Query" name:"Video"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TemplateId           string `position:"Query" name:"TemplateId"`
	Name                 string `position:"Query" name:"Name"`
	TransConfig          string `position:"Query" name:"TransConfig"`
	Audio                string `position:"Query" name:"Audio"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateTemplateRequest) Invoke(client goaliyun.Client) (*UpdateTemplateResponse, error) {
	resp := &UpdateTemplateResponse{}
	err := client.Request("mts", "UpdateTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateTemplateResponse struct {
	RequestId goaliyun.String
	Template  UpdateTemplateTemplate
}

type UpdateTemplateTemplate struct {
	Id          goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Container   UpdateTemplateContainer
	Video       UpdateTemplateVideo
	Audio       UpdateTemplateAudio
	TransConfig UpdateTemplateTransConfig
	MuxConfig   UpdateTemplateMuxConfig
}

type UpdateTemplateContainer struct {
	Format goaliyun.String
}

type UpdateTemplateVideo struct {
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
	BitrateBnd UpdateTemplateBitrateBnd
}

type UpdateTemplateBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type UpdateTemplateAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Remove     goaliyun.String
}

type UpdateTemplateTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type UpdateTemplateMuxConfig struct {
	Segment UpdateTemplateSegment
	Gif     UpdateTemplateGif
}

type UpdateTemplateSegment struct {
	Duration goaliyun.String
}

type UpdateTemplateGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}
