package v1

import (
	"github.com/gin-gonic/gin"
)

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}


// @Summary 獲取多個標籤
// @Produce  json
// @Param name query string false "標籤名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.TagSwagger  "成功"
// @Failure 400 {object} errcode.Error "請求失敗"
// @Failure 500 {object} errcode.Error "伺服器錯誤"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary 更新標籤
// @Produce  json
// @Param id path int true "標籤 ID"
// @Param name body string false "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "請求失敗"
// @Failure 500 {object} errcode.Error "伺服器錯誤"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新標籤
// @Produce  json
// @Param id path int true "標籤 ID"
// @Param name body string false "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "伺服器錯誤"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除標籤
// @Produce  json
// @Param id path int true "標籤 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "伺服器錯誤"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}