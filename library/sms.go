package library

import (
	"encoding/json"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId string, accessKeySecret string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: tea.String(accessKeyId),
		// 您的 AccessKey Secret
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendSMS(client *dysmsapi20170525.Client, phone, code string) (err error) {
	template := map[string]string{"code": code}
	j, _ := json.Marshal(template)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("有孚智能科技"),
		TemplateCode:  tea.String("SMS_242565286"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(string(j)),
	}
	runtime := &util.RuntimeOptions{}
	err = func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _e = client.SendSmsWithOptions(sendSmsRequest, runtime)
		return
	}()

	// if tryErr != nil {
	// 	var error = &tea.SDKError{}
	// 	if _t, ok := tryErr.(*tea.SDKError); ok {
	// 		error = _t
	// 	} else {
	// 		error.Message = tea.String(tryErr.Error())
	// 	}
	// 	// 如有需要，请打印 error
	// 	util.AssertAsString(error.Message)
	// }
	return
}
