package function

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/tagphi/czdb-search-golang/pkg/db"
	"github.com/tuotoo/qrcode"
	"log"
	"math"
	"math/big"
	randNew "math/rand"
	"mime/multipart"
	"os"
	"path"
	"regexp"
	"shop/config"
	"strconv"
	"strings"
	"time"
)

// GetIpAddress 获取IP地理位置，纯真社区IP库
func GetIpAddress(ip string) string {
	region, err := db.Search(ip, config.Cz88Ip)
	if err != nil {
		fmt.Println("IP位置获取失败：", err)
		return ""
	} else {
		re := regexp.MustCompile(`\s+`)
		newStr := re.ReplaceAllString(region, "")
		//fmt.Println(newStr)
		return newStr
	}
}

// Md5 md5加密
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// RandomNumberString 获取随机数字
func RandomNumberString(len int) string {
	var numbers = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()
	for i := 1; i <= len; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {

		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}

// RandNumSimple 取两个数字之间的随机数
func RandNumSimple(min, max int) int {
	if min >= max {
		return min
	}
	// 直接使用 rand.Intn，适用于 Go 1.20+
	return randNew.Intn(max-min+1) + min
}

// CreateARandomString 创建随机字符串
func CreateARandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// CutPinyiin 中文转拼音
func CutPinyiin(hans string) string {
	var result string
	if hans != "" {
		china := pinyin.NewArgs()
		res := pinyin.Pinyin(hans, china)
		for _, v := range res {
			result += v[0]
		}
	} else {
		result = ""
	}
	return result
}

// SlicePage 计算数组分页
func SlicePage(page, pageSize, nums int64) (sliceStart, sliceEnd int64) {
	if page <= 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 50 //设置一页默认显示的记录数
	}
	//if pageSize > nums {
	//	return 0, nums
	//}
	// 总页数
	pageCount := int64(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize
	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}

// QrDecode 二维码解码
func QrDecode(qrFile string) string {
	fi, err := os.Open(qrFile)
	if err != nil {
		log.Println("打开二维码失败", err)
		return ""
	}
	defer fi.Close()
	res, err := qrcode.Decode(fi)
	if err != nil {
		log.Println("二维码解码失败", err)
		return ""
	}
	//fmt.Println(res.Content)
	return res.Content
}

// AesEncrypt 加密
// 生成的密钥需要 (十六进制): 例子：4a5e3d8c1b9f7a2e0d6c4f8a9b3e1d7c
func AesEncrypt(data, sKey string) (string, error) {
	resData := []byte(data)
	key := []byte(sKey)
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//判断加密块的大小
	blockSize := block.BlockSize()
	//填充 判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(resData)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	encryptBytes := append(resData, padText...)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	res := base64.StdEncoding.EncodeToString(crypted)
	return res, nil
}

// AesDecrypt 解密
// 生成的密钥需要 (十六进制): 例子：4a5e3d8c1b9f7a2e0d6c4f8a9b3e1d7c
func AesDecrypt(data, sKey string) (string, error) {
	resData, _ := base64.StdEncoding.DecodeString(data)
	key := []byte(sKey)
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(resData))
	//执行解密
	blockMode.CryptBlocks(crypted, resData)
	//去除填充
	length := len(crypted)
	if length == 0 {
		return "", errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(crypted[length-1])
	crypted = crypted[:(length - unPadding)]
	return string(crypted), nil
}

// StrToUnix 时间转时间戳
func StrToUnix(timeStr, layout string) time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, _ := time.ParseInLocation(layout, timeStr, local)
	//timeUnix := tt.Unix()
	return tt
}

// SaveImageFile 保存上传的图片到服务器
func SaveImageFile(c *gin.Context, filePath string, file *multipart.FileHeader) (fileName, fileUrl string, err error) {
	fileExt := strings.ToLower(path.Ext(file.Filename))
	switch fileExt {
	case ".png", ".jpg", ".gif", ".jpeg":
	default:
		return "", "", fmt.Errorf("不支持的文件类型: %s", fileExt)
	}
	dateDir := "/" + time.Now().Format("20060102")
	fileDir := filePath + dateDir
	if ok := IsFileExist(fileDir); !ok {
		if err = os.MkdirAll(fileDir, 0755); err != nil {
			return "", "", err
		}
	}
	fileName = strconv.FormatInt(time.Now().Unix(), 10) + RandomNumberString(5) + fileExt
	fileSavePath := fileDir + "/" + fileName
	if err = c.SaveUploadedFile(file, fileSavePath); err != nil {
		return "", "", err
	}
	err = os.Chmod(fileDir, 0755)
	if err != nil {
		return "", "", fmt.Errorf("设置目录权限失败: %s", err)
	}
	fileUrl = dateDir + "/" + fileName
	return fileName, fileUrl, nil
}

// SaveBase64ImageFile base64图片写入文件,保存图片
func SaveBase64ImageFile(path string, base64ImageContent string) (bool, string, string) {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64ImageContent)
	if !b {
		return false, "", ""
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	//allData := re.FindAllSubmatch([]byte(base64_image_content), 2)
	//fileType := string(allData[0][1]) //png ，jpeg 后缀获取
	fileType := "txt" //设置后缀获取
	base64Str := re.ReplaceAllString(base64ImageContent, "")
	date := time.Now().Format("20060102")
	if ok := IsFileExist(path + "/" + date); !ok {
		_ = os.Mkdir(path+"/"+date, 0666)
	}
	// umask 会影响 Mkdir 的实际权限，显式 chmod 确保 Nginx 可读
	_ = os.Chmod(path+"/"+date, 0755)
	var fileName = strconv.FormatInt(time.Now().Unix(), 10) + RandomNumberString(5) + "." + fileType
	var file string = path + "/" + date + "/" + fileName
	byteData, _ := base64.StdEncoding.DecodeString(base64Str)
	err := os.WriteFile(file, byteData, 0666)
	if err != nil {
		return false, "", ""
	}
	// 显式 chmod 确保 Nginx 可读
	_ = os.Chmod(file, 0644)
	return true, fileName, file
}

// IsFileExist 判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsURL 正则表达式用于匹配常见的网址格式
func IsURL(str string) bool {
	urlPattern := `^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`
	match, _ := regexp.MatchString(urlPattern, str)
	return match
}

// UploadFileToQiniu 文件上传到七牛云
func UploadFileToQiniu(localFilePath, newFileName, Domain string) (string, error) {
	AccessKey := "-nYPTkUjszQQQjzpxPR69J-kuhlYC6CxbLto1-L4"
	SecretKey := "JU_mdCWPtWyRpcoKIwSBA8tQYgudR3eRn9rZpawB"
	Bucket := "card163"
	// 生成上传凭证
	putPolicy := storage.PutPolicy{Scope: Bucket}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	// 配置上传参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 根据你的存储区域选择
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github",
		},
	}
	// 上传文件
	err := formUploader.PutFile(context.Background(), &ret, upToken, newFileName, localFilePath, &putExtra)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	// 返回文件的公开访问URL
	publicAccessURL := fmt.Sprintf("%s/%s", Domain, ret.Key)
	return publicAccessURL, nil
}

// CalculateTimeDifference 计算两个时间之间的秒数
func CalculateTimeDifference(t1Str, t2Str string) int {
	//t1Str := "2026-04-27 16:14:16"
	//t2Str := "2026-04-27 16:14:45"
	layout := "2006-01-02 15:04:05"
	t1, err := time.Parse(layout, t1Str)
	if err != nil {
		panic(err)
	}
	t2, _ := time.Parse(layout, t2Str)
	diff := t2.Sub(t1)
	seconds := int(diff.Seconds())
	//fmt.Println("相差秒数:", seconds) // 29
	return seconds
}
