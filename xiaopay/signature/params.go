package signature

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Params map[string]string

// map本来已经是引用类型了，所以不需要 *Params
func (p Params) SetString(k, s string) Params {
	p[k] = s
	return p
}

func (p Params) GetString(k string) string {
	s, _ := p[k]
	return s
}

func (p Params) SetInt(k string, i int) Params {
	p[k] = strconv.Itoa(i)
	return p
}

func (p Params) GetInt(k string) int {
	i, _ := strconv.Atoi(p.GetString(k))
	return i
}

func (p Params) SetInt64(k string, i int64) Params {
	p[k] = strconv.FormatInt(i, 10)
	return p
}

func (p Params) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}

// 判断key是否存在
func (p Params) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}

func StructToParams(obj interface{}, strname string) Params {
	//结构体转成xml格式
	//if xmlByteData, err := xml.Marshal(obj); err == nil {
	//	strData := string(xmlByteData)
	//	fmt.Println(strData)
	//	//转成xml格式
	//	fmt.Println(strings.Replace(strData, "RedPacketQueryRequest", "xml", -1))
	//}
	//结构体转成xml格式化的格式
	xmlIndentByteData, err2 := xml.MarshalIndent(obj, "", "  ")
	if err2 != nil {
		return nil
	}
	strData := string(xmlIndentByteData)
	//fmt.Println(strData)
	//fmt.Println(strings.Replace(strData, strname, "xml", -1))

	return XmlToMap(strings.Replace(strData, strname, "xml", -1))
}

func XmlToMap(xmlStr string) Params {
	params := make(Params)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	var (
		key   string
		value string
	)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {

		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			//fmt.Println("11111:",token)
			key = token.Name.Local
		case xml.CharData: // 标签内容
			//fmt.Println("22222:",token)
			content := string([]byte(token))
			value = strings.Replace(content, " ", "", -1)
		default:
			//fmt.Println("33333:",token)
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
	}
	return params
}
