package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int,error)  {
	content,err := ioutil.ReadAll(f)
	return len(content),err
}
func GetExt(fileName string) string  {
	return path.Ext(fileName)
}
func CheckNotExists(src string) bool  {
	_,err := os.Stat(src)
	return os.IsNotExist(err)
}
func CheckPermission(src string) bool  {
	_,err := os.Stat(src)
	return os.IsPermission(err)
}
func MKDir(src string) error  {
	err := os.MkdirAll(src,os.ModePerm)
	if err != nil{
		return err
	}
	return nil
}
func IsNotExistMKDir(src string) error  {
	if notExist := CheckNotExists(src);notExist == true {
		if err:= MKDir(src);err != nil {
			return err
		}
	}
	return nil
}
func Open(name string,flag int,perm os.FileMode) (*os.File,error)  {
	f,err := os.OpenFile(name,flag,perm)
	if err != nil {
		return nil,err
	}
	return f,nil
}
