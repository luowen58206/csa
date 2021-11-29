package upload

import (
	"blog_service.com/m/global"
	"blog_service.com/m/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

/*
	该方法用于针对上传后的文件名格式化
	简单来讲，将文件名 MD5 后再进行写入
	防止直接把原始名 称就暴露出去了
 */

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 获取文件名称、先是通过获取文件后缀并筛选出原始文件名进行MD5加密
//最后返回经过加密处理后的文件名
func GetFileName (name string) string  {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetFileExt 获取文件后缀、主要是通过调用 path.ext 方法进行循环查找 . 符号
//最后通过切片索引返回对应的文件后缀的mingc
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存地址、这里直接返回配置文件中的保存目录即可
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

/*
	·检查保存目录是否存在，通过调用 os.Stat 方法获取文件的描述信息 FileInfo
	·并调用 os.IsNotExist 方法进行判断，
	·，其原理是利用 os.Stat 方法所返回 的 error 值与系统中所定义的 oserror.ErrNotExist 进行判断，以此达到校验效果
 */

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

/*
	·检查文件后缀是否包含在约定的后缀配置项中，
	·需要的是所上传的文件的 后缀有可能是大写、小写、大小写等，
	·因此我们需要调用 strings.ToUpper 方法统一转为大 写（固定的格式）来进行匹配。
*/

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

/*
	·检查文件权限是否足够，与 CheckSavePath 方法原理一致，
	·是利用 oserror.ErrPermission 进行判断。
*/

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}