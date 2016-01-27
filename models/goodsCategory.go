package model
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)



type GoodsCategroy struct {
	CategroyId int `orm:"column(id);pk;size(11);unique;auto"`
	GoodsId int `orm:"column(goodId);index;size(11)"`
	ShopId int `orm:"column(shopId);index;size(11)"`
	Price int64 `orm:"column(price)"`
	Name string `orm:"column(name);size(36)"`
	EnableTime int64 `orm:"column(enbaleTime);default(0)"`
	TotalNum int `orm:"column(totalNum);size(11)"`
	OutNum int `orm:"column(outNum);size(11)"`
	Memo string `orm:"column(memo);type(text)"`
	LimitPurchaseQuantity int `orm:"column(limitPurchaseQuantity);default(1)"`
}

func init() {
	orm.RegisterModel(new(GoodsCategroy))
}

func GetAllCategorysByGoodsId(id int) (*[]GoodsCategroy,error){
	cs := make([]GoodsCategroy,0)
	o := orm.NewOrm()
	_,err := o.Raw(`select * from goods_categroy where goodId = ?`,id).QueryRows(&cs)
	return &cs,err
}

func GetCategoryById(id int) (v * GoodsCategroy,err error) {
	o := orm.NewOrm()
	v = &GoodsCategroy{CategroyId: id}
	if err = o.Read(v,"CategroyId"); err == nil {
		return v, nil
	}
	return nil, err
}

func AddCategory(category *GoodsCategroy)(int64,error) {
	o := orm.NewOrm()
	o.Begin()
	if id,err := o.Insert(category);err == nil&&id > 0 {
		return id,o.Commit()
	}else {
		o.Rollback()
		return 0,err
	}
}

func UpdateCategory(category * GoodsCategroy) (err error){
	o := orm.NewOrm()
	v := GoodsCategroy{CategroyId: category.CategroyId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(category); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteCateGoryById(id int) (err error){
	o := orm.NewOrm()
	o.Begin()
	v := GoodsCategroy{CategroyId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Delete(&v); err == nil {
			return o.Commit()
		}else{
			return o.Rollback()
		}
	}
	return
}

//func GetCateGroyByGoodsId(goodsId int)([]GoodsCategroy,int,error){
//	cs := make([]GoodsCategroy,0)
//	sess := Engine.NewSession()
//	defer sess.Close()
//	err := sess.Where("goodId=?",goodsId).Find(&cs)
//	if err != nil {
//		return cs,0,err
//	}
//	return cs,len(cs),nil
//}
//
//func GetCateGroyById(cateGroyId int) (* GoodsCategroy,error) {
//	sess := Engine.NewSession()
//	defer sess.Close()
//	categroy := GoodsCategroy{}
//	has,err := sess.Where("id=?",cateGroyId).Get(&categroy)
//	if err != nil || !has{
//		beego.Error(err)
//		return nil,err
//	}
//	return &categroy,nil
//}
//
//func AddCateGroyOutCount(categroy *GoodsCategroy) error {
//	sess := Engine.NewSession()
//	defer sess.Close()
//	sess.Begin()
//	_,err := sess.Exec(`update goods_categroy set outNum=? where id=?`,categroy.OutNum+1,categroy.CategroyId)
//	if err != nil {
//		beego.Error(err)
//		sess.Rollback()
//		return err
//	}
//	return sess.Commit()
//}


