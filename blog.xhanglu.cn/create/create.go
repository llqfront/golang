package create

import (
	. "blog.xhanglu.cn/tip"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"time"
)

func Create(c *gin.Context) {
	uname := c.PostForm("uname")
	email := c.PostForm("email")
	content := c.PostForm("content")
	insertTime := time.Now().Unix()

	db, err := sql.Open("mysql", "root:123123@/TESTSQL?charset=utf8")
	defer db.Close()
	CheckErr(err)
	if uname == "" {
		//用户名为空的处理
		c.JSON(200, gin.H{"status": 300, "msg": "用户名不能为空"})
		return
	}
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		//用户名为空的处理
		c.JSON(200, gin.H{"status": 300, "msg": "请你输入正确邮箱地址"})
		return
	}
	//插入数据
	stmt, err := db.Prepare("INSERT book SET uname=?,email=?,content=?,insert_time=?")
	CheckErr(err)
	if _, err := stmt.Exec(uname, content, email, insertTime); err == nil {
		c.JSON(200, gin.H{"status": 200, "msg": "操作成功"})
	}
}
