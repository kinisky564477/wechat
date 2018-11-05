package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	sd  int64
	mtx sync.Mutex
)

// RandLimitString 简易随机字串
func RandLimitString(length ...int) string {
	dict := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	dtLen := len(dict)

	limit := 32
	if len(length) > 0 && length[0] > 0 {
		limit = length[0]
	}

	r := rand.New(rand.NewSource(randSeed()))
	res := ""
	for i := 0; i < limit; i++ {
		res = res + dict[r.Intn(dtLen)]
	}

	return res
}

// MD5 md5 加密
func MD5(target string, salt ...string) string {
	h := md5.New()
	io.WriteString(h, target)
	io.WriteString(h, strings.Join(salt, ""))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSHA256 HmacSha256 加密
func HmacSHA256(target, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, target)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//Sha1  sha1 小写
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GetSignOfPay 微信支付计算 SIGN, 默认是 MD5 加密方式
// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
func GetSignOfPay(data map[string]interface{}, hmacSha256 ...bool) map[string]interface{} {
	strToSign := sortAndEncodeForPay(data)

	nonceStr := RandLimitString()
	nonce, ok := data["nonce_str"]
	if !ok {
		data["nonce_str"] = nonceStr
	} else {
		nonceStr = nonce.(string)
	}

	sign := ""
	if len(hmacSha256) > 0 && hmacSha256[0] {
		enKey := ""
		k, ok := data["key"]
		if ok {
			enKey = k.(string)
		}

		sign = strings.ToUpper(HmacSHA256(strToSign, enKey))
	} else {
		sign = strings.ToUpper(MD5(strToSign))
	}

	return map[string]interface{}{
		"nonce_str": nonceStr,
		"sign":      sign,
	}
}

// Map2XMLString map 转 xml 字串
func Map2XMLString(data map[string]interface{}, cdata ...string) string {
	var res string

	// 引入 sortKeys 是为了排序, 方便测试用例
	sortKeys := make([]string, 0)
	for k := range data {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys)

	for i := range sortKeys {
		k := sortKeys[i]
		v := ""
		switch val := data[k].(type) {
		case string:
			v = val
		case int:
			v = strconv.Itoa(val)
		case int64:
			v = strconv.FormatInt(val, 10)
		default:
			log.Error("转换 map 为 xml 失败, 暂时不支持的类型")
		}

		if strSliceIndex(cdata, k) > -1 {
			res = res + fmt.Sprintf("<%s><![CDATA[%s]]></%s>", k, v, k)
		} else {
			res = res + fmt.Sprintf("<%s>%s</%s>", k, v, k)
		}
	}

	return "<xml>" + res + "</xml>"
}

// RandomIntn 随机数
func RandomIntn(length int, max int) []int {
	seed := rand.NewSource(int64(time.Now().Nanosecond() + rand.Intn(6)))
	r := rand.New(seed)

	res := make([]int, 0)
	for i := 0; i < length; i++ {
		res = append(res, r.Intn(max))
	}

	return res
}

// DecodeMiniData 解密小程序返回数据
// https://developers.weixin.qq.com/miniprogram/dev/api/signature.html
func DecodeMiniData(dt, iv, key string) (map[string]interface{}, error) {
	// 转换 iv
	rawIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	aesIV := []byte(rawIv)

	// 转换 session_key
	rawKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	aesKey := []byte(rawKey)

	// 转换加密数据
	data, err := base64.StdEncoding.DecodeString(dt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, aesIV)

	dist := make([]byte, len(data))
	mode.CryptBlocks(dist, data)
	// dist = pKCS7UnPadding(dt)

	var res map[string]interface{}
	if err := json.Unmarshal(dist, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func pKCS7UnPadding(dt []byte) []byte {
	length := len(dt)
	unpadding := int(dt[length-1])
	return dt[:(length - unpadding)]
}

// sortAndEncodeForPay 微信支付计算 sign 步骤之一
func sortAndEncodeForPay(data map[string]interface{}) string {
	s := make([]string, 0)
	for k, val := range data {
		if k == "key" || k == "sign" || val == nil {
			continue
		}

		switch v := val.(type) {
		case string:
			if v == "" {
				continue
			}
			s = append(s, k+"="+v)
		case int:
			vStr := strconv.Itoa(v)
			s = append(s, k+"="+vStr)
		default:
			log.Error("支付接口中存在不支持的类型")
		}
	}

	sort.Strings(s)
	res := strings.Join(s, "&")

	valOfKey, ok := data["key"]
	if ok {
		res = res + "&key=" + valOfKey.(string)
	}

	return res
}

func randSeed() int64 {
	mtx.Lock()
	defer mtx.Unlock()

	if sd >= 100000000 {
		sd = 1
	}

	sd++
	return time.Now().UnixNano() + sd
}

func strSliceIndex(s []string, key string) int {
	if s == nil {
		return -1
	}

	for k, v := range s {
		if v == key {
			return k
		}
	}

	return -1
}
