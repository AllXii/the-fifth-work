package controller

import (
	"gin/留言板/dao"
	"gin/留言板/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发帖子
func DoPost(c *gin.Context){
	var post models.Post
	// 接收数据
	c.BindJSON(&post)
	//储存数据
	models.CreatePost(&post)
	//返回响应
	c.JSON(http.StatusOK, post)
}

// 删除帖子
func DeletePost(c *gin.Context){
	//post id
	var id map[string]interface{}
	c.BindJSON(&id)
	//删除指定的帖子
	models.DeletePost(id)
	//该帖子下的评论的删除
	models.DeleteComment(id, "all")
	//子评论的删除
	models.DeleteCommentReply(id, "all")
	//返回响应
	c.JSON(http.StatusOK, gin.H{"id": "帖子删除"})
}

// 发评论
func DoComment(c *gin.Context){
	var comment models.Comment
	// 接收post回来的json的数据
	c.BindJSON(&comment)
	// 储存数据
	models.CreateComment(&comment)
	//返回响应
	c.JSON(http.StatusOK, comment)
}

// 删除评论
func DeleteComment(c *gin.Context){
	var id map[string]interface{}
	c.BindJSON(&id)
	//删除指定的数据
	models.DeleteComment(id, "single")
	//子评论的删除
	models.DeleteCommentReply(id, "floor")
	//返回响应
	c.JSON(http.StatusOK, gin.H{"id": "评论删除"})
}

// 点赞
func DoLike(c *gin.Context){
    var id map[string]interface{}
	var comment models.Comment
    c.BindJSON(&id)
    //查询数据
    dao.DB.Where("id = ?", id["id"]).First(&comment)
    //修改数据
    comment.Like++
	dao.DB.Model(&comment).Update("like", comment.Like)
    //返回响应
    c.JSON(http.StatusOK, gin.H{
    	"id": comment.Like,
	})
}

// 回复评论
func DoCommentReply(c *gin.Context){
    var commentReply models.Comment_reply
    //接收数据
    c.BindJSON(&commentReply)
    //储存数据
    models.CreateCommentReply(&commentReply)
    //返回响应
    c.JSON(http.StatusOK, commentReply)
}

//删除回复
func DeleteCommentReply(c *gin.Context){
	var id map[string]interface{}
	//接收数据
	c.BindJSON(&id)
	//删除数据
	models.DeleteCommentReply(id, "single")
	//返回响应
	c.JSON(http.StatusOK, gin.H{"id": "回复删除"})
}


