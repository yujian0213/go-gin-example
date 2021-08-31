package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/models"
	setting "go-gin-example/pkg"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/util"
	"net/http"
)
//文章标签列表
func GetTags(c *gin.Context)  {
	name := c.Query("name")
	//condition
	condition := make(map[string]interface{})
	if name != "" {
		condition["name"] = name
	}
	var state  = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		condition["state"] = state
	}
	//data
	data := make(map[string]interface{})
	data["list"] = models.GetTags(util.GetPage(c),setting.AppSetting.PageSize,condition)
	data["total"] = models.GetTagTotal(condition)

	//return result
	c.JSON(http.StatusOK,gin.H{
		"code":e.SUCCESS,
		"msg":e.GetMsg(e.SUCCESS),
		"data":data,
	})

}
// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")
	vaild := validation.Validation{}
	vaild.Required(name,"name").Message("名称不能为空")
	vaild.MaxSize(name,100,"name").Message("名称最长为100个字符")
	vaild.Required(createdBy,"created_by").Message("创建人不能为空")
	vaild.MaxSize(createdBy,100,"created_by").Message("创建人最长为100个字符")
	vaild.Range(state,0,1,"state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {//验证ok
		if !models.ExistsTagByName(name) {//该数据不存在
			code = e.SUCCESS
			models.AddTag(name,state,createdBy)
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})

}
// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Query("id")).MustInt()
	fmt.Printf("参数%d",id)
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state  = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id,data)
		}else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})
}
//删除文章标签
func DeleteTag(c *gin.Context)  {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id,1,"id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})

}
