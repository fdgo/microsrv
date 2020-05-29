package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	//Insert()
	//Update()
	//Delete()
	//Query()
}
func Insert()  {
	//打开链接
	db,err := sql.Open("mysql","root:000000@tcp(190.27.239.127:3306)/haixian")
	if err != nil {
		fmt.Println("open fail！")
		db.Close()
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping fail！")
		db.Close()
		return
	}
	fmt.Println("connect success!")
	defer db.Close()

	//预处理sql
	stmt, err := db.Prepare("insert into article values(default,?,?,?,?,?,?,?,?)")
	defer func() {
		if stmt!=nil{
			stmt.Close()
		}
	}()
	if err != nil{
		fmt.Println("stmt fail!")
		return
	}
	r,err := stmt.Exec("xiaoshan","2018-12-13 16:40:57",589,"acon","aimg",4555,"2019-12-13 16:40:57","2019-11-13")
	if err != nil{
		fmt.Println("r fail!")
		return
	}
	count,err := r.RowsAffected()
	if err != nil{
		fmt.Println("count fail!")
		return
	}
	if count >0 {
		fmt.Println("add success!")
	}else{
		fmt.Println("add fail!")
	}

	id , _ := r.LastInsertId()
	fmt.Println(id)
}

//相同数据修改认为修改失败！
func Update()  {
	//打开链接
	db,err := sql.Open("mysql","root:000000@tcp(190.27.239.127:3306)/haixian")
	if err != nil {
		fmt.Println("open fail！")
		db.Close()
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping fail！")
		db.Close()
		return
	}
	fmt.Println("connect success!")
	defer db.Close()

	//预处理sql
	stmt, err := db.Prepare("update article set acontent=? where id=?")
	defer func() {
		if stmt!=nil{
			stmt.Close()
		}
	}()
	if err != nil{
		fmt.Println("stmt fail!")
		return
	}
	r,err := stmt.Exec("good",7)
	if err != nil{
		fmt.Println("r fail!")
		return
	}
	count,err := r.RowsAffected()
	if err != nil{
		fmt.Println("count fail!")
		return
	}
	if count >0 {
		fmt.Println("update success!")
	}else{
		fmt.Println("update fail!")
	}
}

func Delete()  {
	//打开链接
	db,err := sql.Open("mysql","root:000000@tcp(190.27.239.127:3306)/haixian")
	if err != nil {
		fmt.Println("open fail！")
		db.Close()
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping fail！")
		db.Close()
		return
	}
	fmt.Println("connect success!")
	defer db.Close()

	//预处理sql
	stmt, err := db.Prepare("delete from article where id=?")
	defer func() {
		if stmt!=nil{
			stmt.Close()
		}
	}()
	if err != nil{
		fmt.Println("stmt fail!")
		return
	}
	r,err := stmt.Exec(7)
	if err != nil{
		fmt.Println("r fail!")
		return
	}
	count,err := r.RowsAffected()
	if err != nil{
		fmt.Println("count fail!")
		return
	}
	if count >0 {
		fmt.Println("delete success!")
	}else{
		fmt.Println("delete fail!")
	}
}

func Query()  {
	//打开链接
	db,err := sql.Open("mysql","root:000000@tcp(190.27.239.127:3306)/haixian")
	if err != nil {
		fmt.Println("open fail！")
		db.Close()
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping fail！")
		db.Close()
		return
	}
	fmt.Println("connect success!")
	defer db.Close()


	stmt, err := db.Prepare("select * from article")
	if err != nil{
		fmt.Println("stmt fail!")
		stmt.Close()
		return
	}
	defer stmt.Close()

	rows,err := stmt.Query()
	if err != nil{
		fmt.Println("Query fail!")
		rows.Close()
		return
	}
	defer rows.Close()
	for rows.Next(){
		var id int
		var arti_name string
		var atime string
		var acount int
		var acontent string
		var aimg string
		var article_type_id int
		var begin_time string
		var update_time string
		rows.Scan(&id,&arti_name,&atime,&acount,&acontent,&aimg,&article_type_id,&begin_time,&update_time)
		fmt.Println(id,arti_name,atime,acount,acontent,aimg,article_type_id,begin_time,update_time)
	}

}