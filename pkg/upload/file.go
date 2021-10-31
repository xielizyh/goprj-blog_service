package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 返回经过处理的文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name)
	// 截断后缀
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	// 获取文件的描述信息
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 检查文件后缀是否允许
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// CheckMaxSize 检查文件大小是否超限
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// CheckPermission 检查文件权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建上传文件所使用的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile 保存上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// 拷贝内容
	_, err = io.Copy(out, src)
	return err
}
