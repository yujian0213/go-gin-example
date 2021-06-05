package week02
// 业务层获取到错误直接往上层抛
type UserService struct {}
func (s *UserService) FindUserByID(userID int64) (*User, error) {
	dao := UserDao{}
	return dao.Get(userID)
}
