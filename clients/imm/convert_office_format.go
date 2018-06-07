package imm

import (
	"github.com/morlay/goaliyun"
)

type ConvertOfficeFormatRequest struct {
	ImageSpec       string `position:"Query" name:"ImageSpec"`
	SrcType         string `position:"Query" name:"SrcType"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	ModelId         string `position:"Query" name:"ModelId"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	MaxSheetRow     int64  `position:"Query" name:"MaxSheetRow"`
	MaxSheetCount   int64  `position:"Query" name:"MaxSheetCount"`
	EndPage         int64  `position:"Query" name:"EndPage"`
	SheetOnePage    string `position:"Query" name:"SheetOnePage"`
	StartPage       int64  `position:"Query" name:"StartPage"`
	MaxSheetCol     int64  `position:"Query" name:"MaxSheetCol"`
	TgtType         string `position:"Query" name:"TgtType"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	TgtUri          string `position:"Query" name:"TgtUri"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ConvertOfficeFormatRequest) Invoke(client goaliyun.Client) (*ConvertOfficeFormatResponse, error) {
	resp := &ConvertOfficeFormatResponse{}
	err := client.Request("imm", "ConvertOfficeFormat", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConvertOfficeFormatResponse struct {
	RequestId  goaliyun.String
	TaskId     goaliyun.String
	TgtLoc     goaliyun.String
	Status     goaliyun.String
	CreateTime goaliyun.String
	Percent    goaliyun.Integer
}
