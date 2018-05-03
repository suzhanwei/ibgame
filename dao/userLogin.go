package dbdao

import (
	"github.com/go-xorm/xorm"
	"reflect"
)

func init() {
	DbStructs["user_login"] = reflect.TypeOf(UserLogin{})
	DbDaos["user_login"] = reflect.TypeOf(UserLoginDao{})
}

type UserLogin struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Name     string `xorm:"not null default '' index VARCHAR(100)"`
	Uid      string `xorm:"not null default '' VARCHAR(40)"`
	Password string `xorm:"not null default '' VARCHAR(40)"`
	Ctime    int64  `xorm:"not null default 0 BIGINT(20)"`
}

type UserLoginDao struct {
	DbBaseDao
}

func NewUserLoginDao(v ...interface{}) *UserLoginDao {
	this := new(UserLoginDao)
	this.UpdateEngine(v...)
	return this
}

func (this *UserLoginDao) Get(mId Param) (ret []UserLogin, err error) {
	ret = make([]UserLogin, 0)
	err = this.getByPrimaryKey(mId, 0, 0).Find(&ret)
	return
}
func (this *UserLoginDao) GetLimit(mId Param, pn, rn int) (ret []UserLogin, err error) {
	ret = make([]UserLogin, 0)
	err = this.getByPrimaryKey(mId, pn, rn).Find(&ret)
	return
}
func (this *UserLoginDao) GetCount(mId Param) (ret int64, err error) {
	ret, err = this.getByPrimaryKey(mId, 0, 0).Count(new(UserLogin))
	return
}
func (this *UserLoginDao) getByPrimaryKey(mId Param, pn, rn int) (session *xorm.Session) {
	session = this.GetSession()

	this.buildQuery(session, mId, "id")

	if rn > 0 {
		session = session.Limit(rn, pn)
	}
	return
}

func (this *UserLoginDao) GetByIdxName(mName Param) (ret []UserLogin, err error) {
	ret = make([]UserLogin, 0)
	err = this.getByIdxNameWithParams(mName, 0, 0).Find(&ret)
	return
}
func (this *UserLoginDao) GetByIdxNameCount(mName Param) (ret int64, err error) {
	ret, err = this.getByIdxNameWithParams(mName, 0, 0).Count(new(UserLogin))
	return
}
func (this *UserLoginDao) GetByIdxNameLimit(mName Param, pn, rn int) (ret []UserLogin, err error) {
	ret = make([]UserLogin, 0)
	err = this.getByIdxNameWithParams(mName, pn, rn).Find(&ret)
	return
}
func (this *UserLoginDao) getByIdxNameWithParams(mName Param, pn, rn int) (session *xorm.Session) {
	session = this.GetSession()

	this.buildQuery(session, mName, "name")

	if rn > 0 {
		session = session.Limit(rn, pn)
	}
	return
}
