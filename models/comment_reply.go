package models

import "gin/留言板/dao"

type Comment_reply struct {
	ID uint `json:"id"`
	Floor uint `json:"floor"`
	PostId uint `json:"post_id"`
	Name string `json:"name"`
	Content string `json:"content"`
}

func CreateCommentReply(commentReply *Comment_reply)(err error){
	if err = dao.DB.Create(&commentReply).Error; err != nil {
		return err
	}
	return
}

func DeleteCommentReply(id map[string]interface{}, obj string)(err error){
	if obj == "single"{
		if err = dao.DB.Where("id = ?", id["id"]).Delete(Comment_reply{}).Error; err != nil{
			return err
		}
	}else if obj == "floor"{
		if err = dao.DB.Where("floor = ?", id["id"]).Delete(Comment_reply{}).Error; err != nil{
			return err
		}
	}else if obj == "all"{
		if err = dao.DB.Where("post_id", id["id"]).Delete(Comment_reply{}).Error; err != nil{
			return err
		}
	}
	return
}
