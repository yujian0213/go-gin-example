package upload

import (
	"fmt"
	setting "go-gin-example/pkg"
	"go-gin-example/pkg/file"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string  {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}
func GetImageName(name string) string  {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name,ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}
func CheckImageSize( f multipart.File) bool  {
	size,err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}
func CheckImage(src string) error {
	dir,err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err:%v",err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
