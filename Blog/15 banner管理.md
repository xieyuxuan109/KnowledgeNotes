## 代码 标准增删改查
```
package banner_api

import (
	"blog_server/common"
	"blog_server/common/res"
	"blog_server/global"
	"blog_server/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BannerApi struct{}
type BannerCreateRequest struct {
	Cover string `json:"cover" binding:"required"`
	Href  string `json:"herf"`
	Show  bool   `json:"show"`
}
type BannerListRequest struct {
	common.PageInfo
	Show bool `json:"show"`
}

func (BannerApi) BannerCreateView(c *gin.Context) {
	var cr BannerCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	err = global.DB.Create(&models.BannerModel{
		Cover: cr.Cover,
		Href:  cr.Href,
		Show:  cr.Show,
	}).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OkWithMsg("添加banner成功", c)
}
func (BannerApi) BannerDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", cr.IDList)
	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	res.OkWithMsg(fmt.Sprintf("删除banner%d个，成功%d个", len(cr.IDList), len(list)), c)
}
func (BannerApi) BannerListView(c *gin.Context) {
	var cr BannerListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	list, count, _ := common.ListQuery(models.BannerModel{
		Show: cr.Show,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
func (BannerApi) BannerUpdateView(c *gin.Context) {
	var id models.IDRequest
	err := c.ShouldBindUri(&id)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var cr BannerCreateRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var model models.BannerModel
	err = global.DB.Take(&model, id.ID).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	global.DB.Model(&model).Updates(map[string]any{
		"Cover": cr.Cover,
		"Href":  cr.Href,
		"Show":  cr.Show,
	})
	res.OkWithMsg("banner更新成功", c)
}
```