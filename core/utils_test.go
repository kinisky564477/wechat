package core_test

import (
	"testing"

	"github.com/zjxpcyc/wechat/core"
)

func TestMap2XMLString(t *testing.T) {
	testData := map[string]interface{}{
		"appid":     "wx2421b1c4370ec43b",
		"attach":    "支付测试",
		"total_fee": 1,
	}

	expected1 := "<xml><appid>wx2421b1c4370ec43b</appid><attach>支付测试</attach><total_fee>1</total_fee></xml>"
	expected2 := "<xml><appid><![CDATA[wx2421b1c4370ec43b]]></appid><attach><![CDATA[支付测试]]></attach><total_fee>1</total_fee></xml>"

	res1 := core.Map2XMLString(testData)
	if res1 != expected1 {
		t.Fatalf("Transfrom map to xml string fail, %s", res1)
	}

	res2 := core.Map2XMLString(testData, "appid", "attach")
	if res2 != expected2 {
		t.Fatalf("Transfrom map to xml string fail, %s", res2)
	}
}

func TestMD5(t *testing.T) {
	cases := []map[string]string{
		map[string]string{
			"target":   "yansen",
			"salt":     "",
			"expected": "e1c91b6b6117f93c1c8734a22acffc2d",
		},
		map[string]string{
			"target":   "yansen",
			"salt":     "issohandsome",
			"expected": "2f0943124b2bf8cab4a6783401c2c4a4",
		},
	}

	for _, c := range cases {
		if c["expected"] != core.MD5(c["target"], c["salt"]) {
			t.Fatalf("校验 MD5 未加盐失败")
		}
	}
}

func TestHmacSha256(t *testing.T) {
	cases := []map[string]string{
		map[string]string{
			"target":   "yansen",
			"key":      "issohandsome",
			"expected": "0c3d3dd9656d1d8f3faecbf14745b13dbd37d7035833ee97836b98f127a0b9b1",
		},
	}

	for _, c := range cases {
		if c["expected"] != core.HmacSHA256(c["target"], c["key"]) {
			t.Fatalf("校验 HmacSha256 失败")
		}
	}
}

func TestGetSignOfPay(t *testing.T) {
	data := map[string]interface{}{
		"appid":       "wxd930ea5d5a258f4f",
		"mch_id":      "10000100",
		"device_info": 1000,
		"body":        "test",
		"nonce_str":   "ibuaiVcKdpRxkhJA",
		"key":         "192006250b4c09247ec02edce69f6a2d",
		"no_value":    "",
	}

	expectedMD5 := "9A0A8659F005D6984697E2CA0A9CF3B7"
	expectedHmacSHA256 := "6A9AE1657590FD6257D693A078E1C3E4BB6BA4DC30B23E0EE2496E54170DACD6"

	res1 := core.GetSignOfPay(data)
	if res1["sign"] != expectedMD5 {
		t.Fatalf("计算支付 Sign (md5) 失败")
	}

	res2 := core.GetSignOfPay(data, true)
	if res2["sign"] != expectedHmacSHA256 {
		t.Fatalf("计算支付 Sign (hmac-sha256) 失败")
	}
}

func TestRandLimitString(t *testing.T) {
	case1 := core.RandLimitString()
	case2 := core.RandLimitString()
	case3 := core.RandLimitString(8)

	if case1 == case2 {
		t.Fatalf("TestRandLimitString Fail - all is same")
	}

	if len(case3) != 8 {
		t.Fatalf("计算 8 长度随机数失败")
	}
}

func TestRandomIntn(t *testing.T) {
	res1 := core.RandomIntn(6, 26)
	res2 := core.RandomIntn(6, 26)

	same := true
	for i := 0; i < 26; i++ {
		if res1[i] != res2[i] {
			same = false
			break
		}
	}

	if same {
		t.Fatalf("TestRandomIntn fail-%v", res1, res2)
	}
}
