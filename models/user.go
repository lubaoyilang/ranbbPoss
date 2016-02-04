package model
import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"errors"
)

type User struct {
	Uid           string `orm:"pk;unique;size(35)"`
	Mobile        string `orm:"unique;size(15)"`
	Password      string `orm:"size(35)"`
	Realname      string `orm:"size(35)"`
	Idcard        string `orm:"unique;size(20)"`
	Alipayaccount string `orm:"unique;size(35)"`
	Alipayname    string `orm:"size(35)"`
	Active        int    `orm:"default(1);size(1)"`
	Asset         int64  `orm:"default(0);size(10)"`
	Rate          int64  `orm:"default(0);size(10)"`
	Income        int64  `orm:"default(0);size(10)"`
	VerifyAmount  int64  `orm:"column(verifyAmount);type(BIGINT);size(10)"`
	Total         int64  `orm:"default(0);size(10)"`
	Createtime    int64  `orm:"default(0);size(10)"`
	Updatetime    int64  `orm:"default(0);size(10)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserById(UID string) (*User,error){
	user := &User{Uid:UID}
	o := orm.NewOrm()
	err := o.Read(user)
	return user,err
}
func UpdateUser(user * User)  error{
	o := orm.NewOrm()
	o.Begin()
	v := User{Uid: user.Uid}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(user); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return o.Commit()
		}else{
			return o.Rollback()
		}
	}
	return errors.New("can not find user")
}