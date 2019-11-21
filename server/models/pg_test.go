package models

import (
	"testing"
	"time"

	// "github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Birthday  time.Time
	Age       int
	Name      string     `gorm:"size:255"`                  // string默认长度为255, 使用这种tag重设。
	Num       int        `gorm:"AUTO_INCREMENT"`            // 自增
	IgnoreMe  int        `gorm:"-"`                         // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}
type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

// func (user *User) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("ID", uuid.New())
// 	return nil
// }

func TestPG(t *testing.T) {
	db, err := InitPG()
	if err != nil {
		t.Log(err)
		return
	}

	db.SingularTable(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Language{})

	t.Log("table user exist:", db.HasTable(&User{}))
	t.Log("table Language exist:", db.HasTable(&Language{}))
	t.Log("table user_languages exist:", db.HasTable("user_languages"))

	// user1 := User{Name: "CreateUser", Age: 18, Birthday: time.Now(), Languages: []Language{Language{ID: 1}}}
	user := User{Name: "CreateUser", Age: 18, Birthday: time.Now()}

	if !db.NewRecord(user) {
		t.Error("User should be new record before create")
		t.Log("user exist")
	}

	if count := db.Save(&user).RowsAffected; count != 1 {
		t.Error("There should be one record be affected when create record")
	} else {
		t.Log("save ok")
	}

	if db.NewRecord(user) || db.NewRecord(&user) {
		t.Error("User should not new record after save")
	} else {
		t.Log("000")
	}

	var newUser User
	if err := db.First(&newUser, user.ID).Error; err != nil {
		t.Errorf("No error should happen, but got %v", err)
	} else {
		t.Log("first")
	}
	t.Log(newUser)
	if newUser.Age != 18 {
		t.Errorf("User's Age should be saved (int)")
	} else {
		t.Log("18")
	}
}

/*


 */
