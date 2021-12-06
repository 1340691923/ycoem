package youdao

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type YouDaoFanyi struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

type DictWeb struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type DictBasic struct {
	UsPhonetic string   `json:"us-phonetic"`
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	UkSpeech   string   `json:"uk-speech"`
	UsSpeech   string   `json:"us-speech"`
	Explains   []string `json:"explains"`
}

type DictResp struct {
	ErrorCode    string                 `json:"errorCode"`
	Query        string                 `json:"query"`
	Translation  []string               `json:"translation"`
	Basic        DictBasic              `json:"basic"`
	Web          []DictWeb              `json:"web,omitempty"`
	Lang         string                 `json:"l"`
	Dict         map[string]interface{} `json:"dict,omitempty"`
	Webdict      map[string]interface{} `json:"webdict,omitempty"`
	TSpeakUrl    string                 `json:"tSpeakUrl,omitempty"`
	SpeakUrl     string                 `json:"speakUrl,omitempty"`
	ReturnPhrase []string               `json:"returnPhrase,omitempty"`
}

func NewYouDaoFanyi(appKey, appSecret string) *YouDaoFanyi {
	return &YouDaoFanyi{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (this *YouDaoFanyi) Fanyi(q, from, to string) (err error) {
	input := this.truncate(q)
	u1 := uuid.NewV4()
	stamp := time.Now().Unix()

	instr := this.AppKey + input + u1.String() + strconv.FormatInt(stamp, 10) + this.AppSecret
	sig := sha256.Sum256([]byte(instr))
	var sigstr string = this.HexBuffToString(sig[:])

	data := make(url.Values, 0)
	data["q"] = []string{q}
	data["from"] = []string{from}
	data["to"] = []string{to}
	data["appKey"] = []string{this.AppKey}
	data["salt"] = []string{u1.String()}
	data["sign"] = []string{sigstr}
	data["signType"] = []string{"v3"}
	data["curtime"] = []string{strconv.FormatInt(stamp, 10)}

	var resp *http.Response
	resp, err = http.PostForm("https://openapi.youdao.com/api", data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(body))
	var jsonObj DictResp
	err = json.Unmarshal(body, &jsonObj)
	if err != nil {
		return
	}
	this.show(&jsonObj, os.Stdout)
	return
}

func (this *YouDaoFanyi) truncate(q string) string {
	res := make([]byte, 10)
	qlen := len([]rune(q))
	if qlen <= 20 {
		return q
	} else {
		temp := []byte(q)
		copy(res, temp[:10])
		lenstr := strconv.Itoa(qlen)
		res = append(res, lenstr...)
		res = append(res, temp[qlen-10:qlen]...)
		return string(res)
	}
}

func (this *YouDaoFanyi) HexBuffToString(buff []byte) string {
	var ret string
	for _, value := range buff {
		str := strconv.FormatUint(uint64(value), 16)
		if len([]rune(str)) == 1 {
			ret = ret + "0" + str
		} else {
			ret = ret + str
		}
	}
	return ret
}

func (this *YouDaoFanyi) show(resp *DictResp, w io.Writer) {
	if resp.ErrorCode != "0" {
		fmt.Fprintln(w, "请输入正确的数据")
	}
	fmt.Fprintln(w, "@", resp.Query)

	if resp.Basic.UkPhonetic != "" {
		fmt.Fprintln(w, "英:", "[", resp.Basic.UkPhonetic, "]")
	}
	if resp.Basic.UsPhonetic != "" {
		fmt.Fprintln(w, "美:", "[", resp.Basic.UsPhonetic, "]")
	}

	fmt.Fprintln(w, "[翻译]")
	for key, item := range resp.Translation {
		fmt.Fprintln(w, "\t", key+1, ".", item)
	}
	fmt.Fprintln(w, "[延伸]")
	for key, item := range resp.Basic.Explains {
		fmt.Fprintln(w, "\t", key+1, ".", item)
	}

	fmt.Fprintln(w, "[网络]")
	for key, item := range resp.Web {
		fmt.Fprintln(w, "\t", key+1, ".", item.Key)
		fmt.Fprint(w, "\t翻译:")
		for _, val := range item.Value {
			fmt.Fprint(w, val, ",")
		}
		fmt.Fprint(w, "\n")
	}
}
