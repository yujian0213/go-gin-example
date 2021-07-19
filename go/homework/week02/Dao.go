package week02

import (
	"fmt"
	"github.com/pkg/errors"
)

type Dao interface {
	Get(id int64) interface{}
	List() interface{}
	Create()
	Update()
	Delete(id uint64)

}
type User struct {
	Id int64 `json:"Id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
type UserDao struct {}
// Dao层获取到底层错误，使用errors的Wrap进行包装
var ErrorSqlNoRows = errors.New("sql.ErrNoRows")
func (u *UserDao) Get(id int64) (*User,error) {
	user := User{}
	err := db.Where("id = ?",id).Find(&user).Error
	if errors.Is(err,ErrorSqlNoRows){
		return &user, errors.Wrap(err,fmt.Sprintf("find user null of user id=%v",id))
	}
	return &user,nil
}