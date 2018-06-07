package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyNotificationConfigurationRequest struct {
	ResourceOwnerAccount string                                               `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                                               `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string                                               `position:"Query" name:"NotificationArn"`
	NotificationTypes    *ModifyNotificationConfigurationNotificationTypeList `position:"Query" type:"Repeated" name:"NotificationType"`
	OwnerId              int64                                                `position:"Query" name:"OwnerId"`
	RegionId             string                                               `position:"Query" name:"RegionId"`
}

func (req *ModifyNotificationConfigurationRequest) Invoke(client goaliyun.Client) (*ModifyNotificationConfigurationResponse, error) {
	resp := &ModifyNotificationConfigurationResponse{}
	err := client.Request("ess", "ModifyNotificationConfiguration", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyNotificationConfigurationResponse struct {
	RequestId goaliyun.String
}

type ModifyNotificationConfigurationNotificationTypeList []string

func (list *ModifyNotificationConfigurationNotificationTypeList) UnmarshalJSON(data []byte) error {
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
