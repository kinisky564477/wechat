package wx

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/zjxpcyc/wechat/core"
)

func checkJSONResult(res map[string]interface{}) error {
	log.Info("接口返回结果: ", res)

	errcode, _ := res["errcode"]
	errmsg, _ := res["errmsg"]
	if errcode == nil {
		return nil
	}

	err, _ := errcode.(float64)
	errNum := int(err)

	if errNum == 0 {
		return nil
	}

	msg, _ := errmsg.(string)
	log.Error("接口返回错误: " + strconv.Itoa(errNum) + "-" + msg)
	return errors.New(strconv.Itoa(errNum) + "-" + msg)
}

// JsTicketSignature 计算 js-ticket signature
func JsTicketSignature(url, noncestr, ticket, timestamp string) string {
	willSign := []string{
		"noncestr=" + noncestr,
		"timestamp=" + timestamp,
		"url=" + url,
		"jsapi_ticket=" + ticket,
	}
	sort.Strings(willSign)
	str2Sign := strings.Join(willSign, "&")

	return core.Sha1(str2Sign)
}
