package models

const TableNameUserSetting = "user_settings"

// UserSetting mapped from table <user_settings>
type UserSetting struct {
	SettingID    int32   `gorm:"column:setting_id;type:integer;primaryKey;autoIncrement:true" json:"setting_id"`
	CustomerID   *int32  `gorm:"column:customer_id;type:integer" json:"customer_id"`
	SettingKey   string  `gorm:"column:setting_key;type:character varying(50);not null" json:"setting_key"`
	SettingValue *string `gorm:"column:setting_value;type:text" json:"setting_value"`
}

// TableName UserSetting's table name
func (*UserSetting) TableName() string {
	return TableNameUserSetting
}
