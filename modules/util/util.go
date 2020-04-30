package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MakeAuthToken() uint64 {
	head := rand.Uint32()
	tail := rand.Uint32()
	token := uint64(head)
	token = token << 31
	token = token | uint64(tail)
	return token
}

func MakeInviteCode(userID uint64) string {
	return string([]byte(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d-%d", userID, time.Now().Unix())))))[:6])
}

func GetBeginTimeOfGivenDay(inTimeSec int64) int64 {
	inTime := time.Unix(inTimeSec, 0)
	y, m, d := inTime.Year(), inTime.Month(), inTime.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
}

func GetBeginTimeOfToday() int64 {
	timeNow := time.Now()
	y, m, d := timeNow.Year(), timeNow.Month(), timeNow.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

/**
 * 生成指定位数验证码(首位不为0)
 */
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		tmpRand := numeric[rand.Intn(r)]
		if i == 0 && tmpRand == 0 {
			i--
			continue
		}
		fmt.Fprintf(&sb, "%d", tmpRand)
	}
	return sb.String()
}

/**
 * uri编码，包含+
 */
func EncodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}

//url:请求地址,response:请求返回的内容
func HttpGet(url string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	response = result.String()
	return
}

//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json，content:请求返回的内容
func HttpPost(url string, data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}

/**
 * 遍历数组内元素是否在str中存在，存在时返回下标，不存在返回-1
 */
func GetArrIdx(arr []string, str string) int {
	idx := -1
	for i := 0; i < len(arr); i++ {
		if strings.Index(str, arr[i]) >= 0 {
			idx = i
			break
		}
	}
	return idx
}

/**
 * 正则提取数字
 */
func GetDigit(str string) int64 {
	reg := regexp.MustCompile(`[\d]`)
	digArr := reg.FindAllString(str, -1)
	result := ""
	for _, i := range digArr {
		result += i
	}
	dig64, _ := strconv.ParseInt(result, 10, 64)
	return dig64
}

/**
 * 根据API路由返回权限标识串
 * @param routePath
 * @returns {string}
 */
func GetPowerStr(routePath string) string {
	_powerStr := "view"
	var powerObj = map[string][]string{
		"view":   {"get", "all", "page", "search", "test"},
		"insert": {"add", "upload", "save"},
		"update": {"update", "active", "bind", "set", "reset", "dispatch"},
		"delete": {"del"},
		"exec":   {"exec", "check", "open", "close", "push"},
	}
	// 去掉第一个/
	routePath = routePath[1:]
	routeArr := strings.Split(routePath, "/")
	last := len(routeArr) - 1
	powerStr := routeArr[last]
	newRouteArr := routeArr[:last]

	for k, _ := range powerObj {
		for _, pValue := range powerObj[k] {
			if strings.Index(powerStr, pValue) != -1 {
				_powerStr = k
				break
			}
		}
	}

	return strings.Replace(strings.Join(newRouteArr, ".")+":"+_powerStr, "api.", "", -1)
}

/**
 * 生成并返回唯一ID
 *
 * @param null
 *
 * @return array
 */
func CreateUuid() (uuid string) {
	mtRand := strconv.Itoa(rand.Intn(1000000000))
	h := md5.New()
	h.Write([]byte(mtRand))
	snippet := hex.EncodeToString(h.Sum(nil))
	uuid = string([]byte(snippet)[0:8]) +
		"-" + string([]byte(snippet)[8:12]) +
		"-" + string([]byte(snippet)[12:16]) +
		"-" + string([]byte(snippet)[16:20]) +
		"-" + string([]byte(snippet)[20:24]) +
		string([]byte(snippet)[24:32])
	return
}