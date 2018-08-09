package get

import (
	. "blog.xhanglu.cn/tip"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Message struct {
	Id         int
	Uname      string
	Email      string
	Content    string
	InsertTime string
}

func Get(c *gin.Context) {
	var (
		status int
		// desc string
		id          int
		insert_time int
		uname       string
		content     string
		email       string
	)
	db, err := sql.Open("mysql", "root:123123@/TESTSQL?charset=utf8")
	defer db.Close()
	// err = db.Ping()  //sql.Open无法断定数据库是否正常连接，所以调用db.Ping()进行判定
	// if err != nil {
	//     status = 300
	//     desc = "数据库连接失败"
	// }else{
	//     status = 200
	//     desc = "数据库连接成功"
	// }
	// rows, err := db.Query("SELECT * FROM book")
	// page := c.DefaultQuery("page", "0")
	// pageSize := c.Query("pageSize")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	pageNo := (page - 1) * pageSize
	// pageNo, err = (strconv.Atoi(page) - 1) * strconv.Atoi(pageSize)
	// fmt.Printf("pageNo type:%T\n", pageNo)
	// pageSize = c.Query("pageSize")
	// fmt.Printf(c.Param("name"))
	// rows, err := db.Query("SELECT * FROM book limit 10, 2")
	// 方法三  order by ASC/DESC
	rows, err := db.Query("select * from book where id > ? order by id limit ?", pageNo, pageSize)
	// 方法二
	// rows, err := db.Query("select * from book where id > ? limit ?", pageNo, pageSize)
	// 方法二
	// rows, err := db.Query("select * from book limit ?, ?", pageNo, pageSize)

	var msgSlice []*Message
	for rows.Next() {
		err = rows.Scan(&id, &uname, &content, &email, &insert_time)
		CheckErr(err)
		msg := new(Message)
		msg.Id = id
		msg.InsertTime = time.Unix(int64(insert_time), 0).Format("2006-01-02 15:04:00")
		msg.Uname = uname
		msg.Content = content
		msg.Email = email
		//fmt.Fprintf(w, id)
		msgSlice = append(msgSlice, msg)
	}
	CheckErr(err)
	c.JSON(200, gin.H{"status": status, "msg": msgSlice})
}
