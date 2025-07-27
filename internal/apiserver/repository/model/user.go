package model

import (
	"time"
)

const TableNameUserM = "users"

// UserM 用户表
type UserM struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Password  string    `gorm:"column:password;type:varchar(255);not null;comment:用户密码（加密后）" json:"password"`                                            // 用户密码（加密后）
	Email     string    `gorm:"column:email;type:varchar(100);unique;not null;comment:用户电子邮箱地址" json:"email"`                                                   // 用户电子邮箱地址
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP(3);comment:用户创建时间" json:"createdAt"`                                  // 用户创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:用户最后修改时间" json:"updatedAt"` // 用户最后修改时间
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return TableNameUserM
}

// func (m *UserM) AfterCreate(tx *gorm.DB) error {
// 	m.UserId = GenerateRandomID(m.Id)
// 	return tx.Save(m).Error
// }

// func GenerateRandomID(dbID int64) string {
// 	// 随机生成 4 字节的随机值，作为自增 ID 的后缀
// 	randomBytes := make([]byte, 4)
// 	_, _ = rand.Read(randomBytes)                   // 生成随机数
// 	randomSuffix := hex.EncodeToString(randomBytes) // 转为十六进制

// 	return fmt.Sprintf("ID-%d-%s", dbID, randomSuffix) // 组合生成唯一 ID
// }
