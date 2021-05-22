package models

import (
	"gin/留言板/dao"
)

type Comment struct {
	ID uint `json:"id"`
	PostId uint `json:"post_id"`
	Content string `json:"content"`
	Name string `json:"name"`
	Like uint `json:"like"`
}

// comment增加操作
func CreateComment(comment *Comment)(err error){
	if err = dao.DB.Create(&comment).Error; err != nil {
		return err
	}
    return
}

// comment指定删除操作
func DeleteComment(id map[string]interface{}, obj string)(err error){
	if obj == "single"{
	    if err = dao.DB.Where("id = ?", id["id"]).Delete(Comment{}).Error; err != nil{
		    return err
	    }
	}else if obj == "all" {
		if err = dao.DB.Where("post_id = ?", id["id"]).Delete(Comment{}).Error; err != nil{
			return err
		}
	}
	return
}