package models

import "gin/留言板/dao"

type Post struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Name string `json:"name"`
	Content string `json:"content"`
}

// 帖子的增加
func CreatePost(post *Post)(err error){
	if err = dao.DB.Create(&post).Error; err != nil{
		return err
	}
	return
}

//帖子的删除
func DeletePost(id map[string]interface{})(err error){
	if err = dao.DB.Where("id = ?", id["id"]).Delete(Post{}).Error; err != nil{
		return err
	}
	return
}