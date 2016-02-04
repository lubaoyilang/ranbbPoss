package model
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Enchashment struct {
	Id  		int 	`orm:"column(id);pk;size(11);auto"`
	UID 		string 	`orm:"column(UID);index;size(36)"`
	Amount 		int64 	`orm:"column(amount);size(16);default(0)"`
	CreateTime  int64  	`orm:"column(createTime);default(0);size(16)"`
	UpdateTime  int64  	`orm:"column(updateTime);default(0);size(16)"`
	Operator 	string  `orm:"column(operator);index;size(30)"`
	State		int  	`orm:"column(state);index;size(1);default(0)"`
	Memo		string 	`orm:"column(memo);type(TEXT)"`
}

type EnchashmentInfo struct {
	Id  		int `orm:"column(id);pk;size(11);auto"`
	UID 		string 	`orm:"column(UID);index;size(36)"`
	Realname      string `orm:"column(realName);size(35)"`
	Alipayaccount string `orm:"column(aliPayAccount);unique;size(35)"`
	Amount 		int64 	`orm:"column(amount);size(16);default(0)"`
	Asset         int64  `orm:"column(asset)"`
	Rate          int64  `orm:"column(rate)"`
	Income        int64  `orm:"column(income)"`
	VerifyAmount  int64  `orm:"column(verifyAmount);type(BIGINT);size(10)"`
	Total         int64  `orm:"column(total)"`
	CreateTime  int64  	`orm:"column(createTime);default(0);size(16)"`
	State		int  	`orm:"column(state);index;size(1);default(0)"`
}

func init() {
	orm.RegisterModel(new(Enchashment))
}
func GetAllUnDealReq() (*[]Enchashment,error){
	var reqs []Enchashment
	o := orm.NewOrm()
	_,err := o.Raw(`select * from enchashment where state=0 order by createTime`).QueryRows(&reqs)
	return &reqs,err
}
func GetUserEnchashmentList(page int) (*[]EnchashmentInfo,error){
	if page <=0 {
		page = 1
	}
	var reqs []EnchashmentInfo
	o := orm.NewOrm()
	_,err :=  o.Raw(`select * from enchashment e left join user u on e.UID=u.UID  where state!=1  order by e.createTime limit 10 offset ?`,(page-1)*10).QueryRows(&reqs)
	return &reqs,err
}
func GetUserEnchashmentById(EnchashmentId int)(v * EnchashmentInfo,err error) {
	v = &EnchashmentInfo{}
	o := orm.NewOrm()
	err = o.Raw(`select * from enchashment e left join user u on e.UID=u.UID  where state!=1  and id=?`,EnchashmentId).QueryRow(v)
	return
}

func UpdateEnchashmentState(state,id int,operator string) (err error){
	o := orm.NewOrm()
	o.Begin()
	v := &Enchashment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(v); err == nil {
		var num int64
		v.State = state
		v.Operator = operator
		if num, err = o.Update(v); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return o.Commit()
		}else{
			return o.Rollback()
		}
	}
	return
}