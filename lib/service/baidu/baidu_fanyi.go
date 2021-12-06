package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BaiduFanyi struct {
	Appid    int
	Password string
}

func NewBaiduFanyi(appid int, password string) *BaiduFanyi {
	return &BaiduFanyi{
		Appid:    appid,
		Password: password,
	}
}

func (this *BaiduFanyi) Fanyi(q, from, to string) (dist string, err error) {
	tran := NewTranslateModeler(q, from, to, this.Password, this.Appid)
	values := tran.ToValues()
	//百度翻译api接口
	var Url = "http://api.fanyi.baidu.com/api/trans/vip/translate"

	resp, err := http.PostForm(Url, values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(body))

	var res struct {
		From        string `json:"from"`
		To          string `json:"to"`
		TransResult []struct {
			Src string `json:"src"`
			Dst string `json:"dst"`
		} `json:"trans_result"`
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return
	}

	if len(res.TransResult) == 0 {
		return
	}

	dist = res.TransResult[0].Dst

	return
}

type TranslateModel struct {
	Q     string
	From  string
	To    string
	Appid int
	Salt  int
	Sign  string
}

func NewTranslateModeler(q, from, to, password string, appID int) TranslateModel {
	tran := TranslateModel{
		Q:    q,
		From: from,
		To:   to,
	}
	tran.Appid = appID
	tran.Salt = time.Now().Second()
	content := strconv.Itoa(appID) + q + strconv.Itoa(tran.Salt) + password
	sign := SumString(content) //计算sign值
	tran.Sign = sign
	return tran
}

func (tran TranslateModel) ToValues() url.Values {
	values := url.Values{
		"q":     {tran.Q},
		"from":  {tran.From},
		"to":    {tran.To},
		"appid": {strconv.Itoa(tran.Appid)},
		"salt":  {strconv.Itoa(tran.Salt)},
		"sign":  {tran.Sign},
	}
	return values
}

//计算文本的md5值
func SumString(content string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(content))
	bys := md5Ctx.Sum(nil)
	//bys := md5.Sum([]byte(content))//这个md5.Sum返回的是数组,不是切片哦
	value := hex.EncodeToString(bys)
	return value
}

const EN = "en"
const ZH = "zh"
