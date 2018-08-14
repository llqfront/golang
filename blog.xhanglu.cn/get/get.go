package get

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	. "blog.xhanglu.cn/tip"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
		// status int
		// desc        string
		id          int
		insert_time int
		uname       string
		content     string
		email       string
	)
	db, err := sql.Open("mysql", "root:123123@/TESTSQL?charset=utf8")
	err = db.Ping() //sql.Open无法断定数据库是否正常连接，所以调用db.Ping()进行判定
	defer db.Close()
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
	// rows, err := db.Query("select * from book where id > ? order by id limit ?", pageNo, pageSize)
	// count, err := db.Query("select found_rows()")

	// var count int64

	// fmt.Print(count)
	//

	// 方法二
	// rows, err := db.Query("select * from book where id > ? limit ?", pageNo, pageSize)
	// 方法二
	rows, err := db.Query("select * from book limit ?, ?", pageNo, pageSize)

	var dataSlice []*Message
	for rows.Next() {
		err = rows.Scan(&id, &uname, &content, &email, &insert_time)
		CheckErr(err)
		data := new(Message)
		data.Id = id
		data.InsertTime = time.Unix(int64(insert_time), 0).Format("2006-01-02 15:04:05")
		data.Uname = uname
		data.Content = content
		data.Email = email
		//fmt.Fprintf(w, id)
		dataSlice = append(dataSlice, data)
	}
	var count int
	err = db.QueryRow("select count(*) from book").Scan(&count)
	CheckErr(err)
	// "count": "1", 总数 ok
	// "size": "20",  一页多少条 ok
	// "page": "1",  当前页码 ok
	c.JSON(200, gin.H{"status": http.StatusOK, "result": dataSlice, "msg": "操作成功", "count": count, "page": page, "size": pageSize})
}
