package model

import (
	"casbintest/global"
	"errors"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strings"
)

//casbin_rule表中
//  p_type --类型，可以是p策略，g角色等等
//  v0 --角色 roleName/roleId   sub
//  v1 --Path 路径              obj
//  v2 --Method 请求方式         act
//  v3 --允许读/写 read/write
//  v4 --允许/拒绝 allow/deny
//  v5 --
//
//ptype=g的时候v1=角色
//ptype=p的时候, v0=subject, v1=obj, v2=action

type CasbinModel struct {
	PType  string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	RoleId string `json:"role_id" gorm:"column:v0" description:"角色ID"`
	ProId  string `json:"pro_id" gorm:"column:v1" description:"项目ID"`
	Path   string `json:"path" gorm:"column:v2" description:"api路径"`
	Method string `json:"method" gorm:"column:v3" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Create() error {
	if success := global.Enforcer.AddPolicy(c.RoleId, c.ProId, c.Path, c.Method); success == false {
		return errors.New("存在相同的API，添加失败")
	}
	return nil
}

func (c *CasbinModel) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(c).Where("v1 = ? AND v2 = ?", c.Path, c.Method).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (c *CasbinModel) List() [][]string {

	policy := global.Enforcer.GetFilteredPolicy(0, c.RoleId)
	return policy
}

// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.Enforcer
func Casbin() *casbin.Enforcer {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=123 sslmode=disable TimeZone=Asia/Shanghai")
	fmt.Println("11", err)
	adapter := gormadapter.NewAdapterByDB(db)
	enforcer := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	enforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = enforcer.LoadPolicy()
	return enforcer
}

// @function: ClearCasbin
// @description: 清除匹配的权限
// @param: v int, p ...string
// @return: bool
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)

}

// @function: ParamsMatch
// @description: 自定义规则函数
// @param: fullNameKey1 string, key2 string
// @return: bool
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @function: ParamsMatchFunc
// @description: 自定义规则函数
// @param: args ...interface{}
// @return: interface{}, error
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
