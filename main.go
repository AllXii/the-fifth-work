package main

import (
	"gin/留言板/controller"
	"gin/留言板/dao"
	"gin/留言板/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
    err := dao.InitMySQL()
    if err != nil{
    	panic(err)
	}
	//创建表
	dao.DB.AutoMigrate(&models.Post{})
	dao.DB.AutoMigrate(&models.Comment{})
    dao.DB.AutoMigrate(&models.Comment_reply{})
	r := gin.Default()
	// 访问页面
	r.GET("/", nil)
	//实现帖子的发表
	r.POST("/post", controller.DoPost)
	//实现帖子的删除
	r.DELETE("/post", controller.DeletePost)
	// 实现评论的发表
	r.POST("/comment",  controller.DoComment)
	//实现评论的删除
	r.DELETE("/comment", controller.DeleteComment)
	//实现评论点赞
	r.PUT("/comment", controller.DoLike)
    //实现回复
	r.POST("/reply", controller.DoCommentReply)
	//实现回复的删除
    r.DELETE("/reply",controller.DeleteCommentReply)
	r.Run(":8080")
}