package models

const TableNameSupplier = "supplier"

// Supplier mapped from table <supplier>
type Supplier struct {
	SupplierID  int32   `gorm:"column:supplier_id;type:integer;primaryKey;autoIncrement:true" json:"supplier_id"`
	Name        string  `gorm:"column:name;type:character varying(100);not null" json:"name"`
	ContactName *string `gorm:"column:contact_name;type:character varying(50)" json:"contact_name"`
	Email       *string `gorm:"column:email;type:character varying(100)" json:"email"`
	Phone       *string `gorm:"column:phone;type:character varying(20)" json:"phone"`
	Address     *string `gorm:"column:address;type:text" json:"address"`
}

// TableName Supplier's table name
func (*Supplier) TableName() string {
	return TableNameSupplier
}
