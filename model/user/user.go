package user

import (
	"github.com/google/uuid"
	"qh-gin-api/global"
)

type User struct {
	global.QGA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                       // 用户UUID
	Username string    `json:"userName" gorm:"index;comment:用户登录名"`                                                    // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                                                               // 用户登录密码
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                              // 用户昵称
	Avatar   string    `json:"avatar" gorm:"default:https://img.scdn.io/i/6881a40048046_1753326592.webp;comment:用户头像"` // 用户头像
	Phone    string    `json:"phone"  gorm:"comment:用户手机号"`                                                            // 用户手机号
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`                                                             // 用户邮箱
	Enable   int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                        //用户是否被冻结 1正常 2冻结
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetNickname() string {
	return u.NickName
}

func (u *User) GetUUID() uuid.UUID {
	return u.UUID
}

func (u *User) GetUserId() uint {
	return u.ID
}

func (u *User) GetUserInfo() any {
	return *u
}
