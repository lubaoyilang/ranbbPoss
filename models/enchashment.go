package model
import (
	"github.com/astaxie/beego/orm"
)

type Enchashment struct {
	Id  		int 	`orm:"column(id);pk;size(11);auto"`
	UID 		string 	`orm:"column(UID);index;size(36)"`
	Amount 		int64 	`orm:"column(Amount);size(16);default(0)"`
	CreateTime  int64  	`orm:"column(createTime);default(0);size(16)"`
	UpdateTime  int64  	`orm:"column(updateTime);default(0);size(16)"`
	Operator 	string  `orm:"column(operator);index;size(30)"`
	State		int  	`orm:column(state);index;size(1);default(0)`
	Memo		string 	`orm:column(memo);type(TEXT)`
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