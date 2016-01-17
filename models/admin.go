package model
import (
	"github.com/astaxie/beego/orm"
	"encoding/gob"
	"github.com/astaxie/beego"
	"reflect"
)

type Admin struct {
	Uid       string `orm:"column(UID);pk"`
	Name      string `orm:"column(name)"`
	Mobile    string `orm:"column(mobile);size(15);unique"`
	Password  string `orm:"column(password);size(36)" json:"-"`
	Privilege string `orm:"column(privilege);size(100)"`
	State 	  int 	 `orm:"column(state);size(1);default(0)"`
}

func init() {
	orm.RegisterModel(new(Admin))
	gob.Register(new(Admin))
}

func GetAdminByUserName(userName string) (*Admin,error) {
	admin := &Admin{Name:userName}
	o := orm.NewOrm()
	err := o.Read(admin,"Name")
	return admin,err
}

func GetAdminListNotRoot(page,size int)(*[]Admin,error){
	var admins []Admin
	if page == 0 {
		page=1
	}
	if size == 0 {
		size = 10
	}
	o := orm.NewOrm()
	_,err := o.Raw(`select * from admin where privilege != 'root' limit ? OFFSET ?`,size,(page-1)*size).QueryRows(&admins)
	return &admins,err
}

func AddAdmin(admin *Admin) error {
	o := orm.NewOrm()
	o.Begin()
	_,err := o.Insert(admin)
	if err != nil {
		o.Rollback()
		return err
	}
	return o.Commit()
}

func DeleteAdmin(mobile string)error {
	o:= orm.NewOrm()
	o.Begin()
	_,err := o.Raw(`delete from admin where mobile=?`,mobile).Exec()
	if err != nil {
		o.Rollback()
		beego.Error(err.Error())
		return err
	}else{
		return o.Commit()
	}
}

func UpdateAdmin(admin * Admin) error {
	beego.Debug("UpdateAdmin",admin)
	o:= orm.NewOrm()
	o.Begin()
	if len(admin.Password)> 0 {
		_,err := o.Raw(`update admin set password=?,privilege=?,state=?,mobile=? where name=?`,admin.Password,admin.Privilege,admin.State,admin.Mobile,admin.Name).Exec()
		if err != nil {
			o.Rollback()
			beego.Error(err.Error())
			return err
		}else{
			return o.Commit()
		}
	}else{
		_,err := o.Raw(`update admin set privilege=?,state=?,mobile=? where name=?`,admin.Privilege,admin.State,admin.Mobile,admin.Name).Exec()
		if err != nil {
			o.Rollback()
			beego.Error(err.Error())
			return err
		}else{
			return o.Commit()
		}
	}

//	_,err := o.QueryTable("admin").Update(Struct2Map(admin))

}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}