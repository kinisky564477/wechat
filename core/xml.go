package core

import (
	"errors"
	"regexp"
	"strings"
)

/**
* 针对微信支付接口设计实现的极简 XML 解析器
*
**/

const (
	NodeTag = iota + 1
	NodeValue
)

type XMLCDATA struct {
	Value string `xml:",cdata"`
}

type XMLElement struct {
	TagName  string
	Value    string
	Children []XMLElement
}

type XMLParse struct{}

func (t *XMLParse) Parse(xmlStr string) (map[string]string, error) {
	log.Info("开始解析 XML 字串: " + xmlStr)

	xml, err := t.parse(xmlStr)
	if err != nil {
		log.Error("解析 XML 失败", err.Error())
		return nil, err
	}

	res := make(map[string]string)

	// 忽略根节点
	for _, node := range xml.Children {
		res[node.TagName] = node.Value
	}

	return res, nil
}

func (t *XMLParse) parse(xmlStr string) (*XMLElement, error) {
	pureXML := strings.TrimSpace(xmlStr)
	tagOpen := regexp.MustCompile(`(?msi)<([^>]*)>`)
	tagInfo := tagOpen.FindStringSubmatch(pureXML)
	if tagInfo == nil {
		return nil, errors.New("xml 格式不正确, 没有开始标签")
	}

	ele := XMLElement{}
	ele.TagName = tagInfo[1]

	tagReg := regexp.MustCompile(`(?msi)<` + ele.TagName + `>(.*)</` + ele.TagName + `>`)
	tagVals := tagReg.FindStringSubmatch(pureXML)
	if tagVals == nil {
		return nil, errors.New("xml 格式不正确, 可能没有闭合标签")
	}

	val := strings.TrimSpace(tagVals[1])
	if strings.Index(val, "<![CDATA[") == 0 {
		val = strings.TrimLeft(val, "<![CDATA[")
		val = strings.TrimRight(val, "]]>")
	}
	ele.Value = val

	if tagOpen.Match([]byte(val)) {
		childrenStr := val
		for {
			child, err := t.parse(childrenStr)
			if err != nil {
				return nil, err
			}

			if ele.Children == nil {
				ele.Children = make([]XMLElement, 0)
			}

			ele.Children = append(ele.Children, *child)

			childCloseTag := "</" + child.TagName + ">"
			leftChildren := strings.Split(childrenStr, childCloseTag)
			childrenStr = strings.TrimSpace(leftChildren[1])
			if childrenStr == "" {
				break
			}
		}
	}

	return &ele, nil
}
