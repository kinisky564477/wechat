package component

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/kinisky564477/wechat/core"
)

// CDATA CDATA
type CDATA struct {
	Value string `xml:",cdata"`
}

// PassiveMessageText 文本
type PassiveMessageText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Content      CDATA    `xml:"Content"`
}

// PassiveMessageImage 图片
type PassiveMessageImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	MediaID      CDATA    `xml:"MediaId"`
}

// PassiveMessageVoice 语音
type PassiveMessageVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	MediaID      CDATA    `xml:"MediaId"`
}

// PassiveMessageVideo 视频
type PassiveMessageVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	MediaID      CDATA    `xml:"MediaId"`
	Title        CDATA    `xml:"Title"`
	Description  CDATA    `xml:"Description"`
}

// PassiveMessageMusic 音乐
type PassiveMessageMusic struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Title        CDATA    `xml:"Title"`
	Description  CDATA    `xml:"Description"`
	MusicURL     CDATA    `xml:"MusicURL"`
	HQMusicURL   CDATA    `xml:"HQMusicUrl"`
	ThumbMediaID CDATA    `xml:"ThumbMediaId"`
}

// ArticleItem 图文节点
type ArticleItem struct {
	Title       CDATA `xml:"Title"`
	Description CDATA `xml:"Description"`
	PicURL      CDATA `xml:"PicUrl"`
	URL         CDATA `xml:"Url"`
}

// PassiveMessageNews 图文
type PassiveMessageNews struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CDATA         `xml:"ToUserName"`
	FromUserName CDATA         `xml:"FromUserName"`
	CreateTime   int64         `xml:"CreateTime"`
	MsgType      CDATA         `xml:"MsgType"`
	ArticleCount int           `xml:"MediaId"`
	Articles     []ArticleItem `xml:"Articles>item"`
}

func (t *WxClient) TransformMessage(msg string) (map[string]string, error) {
	xp := &core.XMLParse{}
	return xp.Parse(msg)
}

// ResponseMessageText 回复文本消息
func (t *WxClient) ResponseMessageText(from, to, message string) ([]byte, error) {
	log.Info("(被动)待反馈文本消息: ", fmt.Sprintf("%s", message))

	data := PassiveMessageText{
		ToUserName:   CDATA{to},
		FromUserName: CDATA{from},
		MsgType:      CDATA{"text"},
		CreateTime:   time.Now().Local().Unix(),
		Content:      CDATA{message},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageImage 回复图片消息
func (t *WxClient) ResponseMessageImage(from, to, media string) ([]byte, error) {
	log.Info("(被动)待反馈图片消息: ", fmt.Sprintf("%s", media))

	data := PassiveMessageImage{
		ToUserName:   CDATA{to},
		FromUserName: CDATA{from},
		MsgType:      CDATA{"image"},
		CreateTime:   time.Now().Local().Unix(),
		MediaID:      CDATA{media},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageVoice 回复音频消息
func (t *WxClient) ResponseMessageVoice(from, to, media string) ([]byte, error) {
	log.Info("(被动)待反馈音频消息: ", fmt.Sprintf("音频ID %s", media))

	data := PassiveMessageVoice{
		ToUserName:   CDATA{to},
		FromUserName: CDATA{from},
		MsgType:      CDATA{"voice"},
		CreateTime:   time.Now().Local().Unix(),
		MediaID:      CDATA{media},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageVideo 回复视频消息
func (t *WxClient) ResponseMessageVideo(from, to, media, title, desc string) ([]byte, error) {
	log.Info("(被动)待反馈视频消息: ", fmt.Sprintf("视频ID %s , 标题 %s , 描述 %s ", media, title, desc))

	data := PassiveMessageVideo{
		ToUserName:   CDATA{to},
		FromUserName: CDATA{from},
		MsgType:      CDATA{"video"},
		CreateTime:   time.Now().Local().Unix(),
		MediaID:      CDATA{media},
		Title:        CDATA{title},
		Description:  CDATA{desc},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageMusic 回复音乐消息
func (t *WxClient) ResponseMessageMusic(from, to, music, title, desc string, others ...string) ([]byte, error) {
	log.Info("(被动)待反馈音乐消息: ", fmt.Sprintf("音乐 %s , 标题 %s , 描述 %s, 其他: %v ", music, title, desc, others))

	othLen := len(others)
	hqMusicURL := ""
	thumbMediaID := ""
	if othLen > 0 {
		hqMusicURL = others[0]

		if othLen > 1 {
			thumbMediaID = others[1]
		}
	}

	data := PassiveMessageMusic{
		ToUserName:   CDATA{to},
		FromUserName: CDATA{from},
		MsgType:      CDATA{"music"},
		CreateTime:   time.Now().Local().Unix(),
		Title:        CDATA{title},
		Description:  CDATA{desc},
		MusicURL:     CDATA{music},
		HQMusicURL:   CDATA{hqMusicURL},
		ThumbMediaID: CDATA{thumbMediaID},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageNews 回复图文消息
func (t *WxClient) ResponseMessageNews(from, to string, articles []map[string]string) ([]byte, error) {
	log.Info("(被动)待反馈图文消息: ", fmt.Sprintf("%v", articles))

	num := len(articles)
	items := make([]ArticleItem, 0)
	for _, article := range articles {
		item := ArticleItem{
			Title:       CDATA{article["title"]},
			Description: CDATA{article["desc"]},
			PicURL:      CDATA{article["picurl"]},
			URL:         CDATA{article["url"]},
		}

		items = append(items, item)
	}

	data := PassiveMessageNews{
		ToUserName:   CDATA{to},
		FromUserName: CDATA{from},
		MsgType:      CDATA{"news"},
		CreateTime:   time.Now().Local().Unix(),
		ArticleCount: num,
		Articles:     items,
	}

	return t.getXMLStringOfMessage(data)
}

// getXMLStringOfMessage 获取 xml 字串
func (t *WxClient) getXMLStringOfMessage(message interface{}) ([]byte, error) {
	res, err := xml.MarshalIndent(message, "", "")
	if err != nil {
		log.Error("转换消息xml内容失败: ", err.Error())
		return nil, err
	}

	return res, nil
}
