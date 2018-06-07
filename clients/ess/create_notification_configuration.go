package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateNotificationConfigurationRequest struct {
	ResourceOwnerAccount string                                               `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                                               `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string                                               `position:"Query" name:"NotificationArn"`
	NotificationTypes    *CreateNotificationConfigurationNotificationTypeList `position:"Query" type:"Repeated" name:"NotificationType"`
	OwnerId              int64                                                `position:"Query" name:"OwnerId"`
	RegionId             string                                               `position:"Query" name:"RegionId"`
}

func (req *CreateNotificationConfigurationRequest) Invoke(client goaliyun.Client) (*CreateNotificationConfigurationResponse, error) {
	resp := &CreateNotificationConfigurationResponse{}
	err := client.Request("ess", "CreateNotificationConfiguration", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNotificationConfigurationResponse struct {
	RequestId goaliyun.String
}

type CreateNotificationConfigurationNotificationTypeList []string

func (list *CreateNotificationConfigurationNotificationTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
