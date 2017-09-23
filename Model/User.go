package Model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Entity struct {
	ID         string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (entity *Entity) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

/*User Info entity*/
type User struct {
	Entity
	Name         string  `gorm:"not null"`
	Email        string  `gorm:"not null;unique"`
	Profile      Profile `gorm:"ForeignKey:ProfileRefer"`
	SecurityInfo	SecurityInfo `gorm:"ForeignKey:SecurityInfoRefer"`
	BillAddress []Address `gorm:"ForeignKey:BillAddressRefer"`
	ShippingAddress []Address `gorm:"ForeignKey:ShippingAddressRefer"`
	ProfileRefer string
	SecurityInfoRefer string
}

type Profile struct {
	Entity
	Avatar      string
	DateOfBirth time.Time
	Country     string
}

type SecurityInfo struct {
	Entity
	UserName string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

type Address struct {
	Entity
	HouseNo string
	City string
	Street string
	District string
	ShippingAddressRefer string
	BillAddressRefer string
	StoreAddressRefer string
}

/*Store Entity*/
type Store struct {
	Entity
	Phone string
	AddressStore Address `gorm:"ForeignKey:StoreAddressRefer"`
	Admin User `gorm:"ForeignKey:StoreAdminRefer"`
	AboutStore string
	WareHouseOfStore WareHouse `gorm:"ForeignKey:WareHouseRefer"`
	StoreAdminRefer string
	WareHouseRefer string
	StoreAddressRefer string
}

type WareHouse struct {
	Entity
	AddressWare Address `gorm:"ForeignKey:AddressWareRefer"`
	Inventories []Inventory  `gorm:"ForeignKey:InventoriesRefer"`
	AddressWareRefer string
}

type Inventory struct {
	Entity
	ProductIn	Product `gorm:"ForeignKey:ProductRefer"`
	NoInStore int
	PriceIn float64
	PriceOut float64
	MinimumBalance int
	InventoriesRefer string
	ProductRefer string
}

type Product struct {
	Entity
	Name string
	Information string
	Category ProductCategory `gorm:"ForeignKey:CategoryRefer"`
	Supply Store `gorm:"ForeignKey:SupplyRefer"`
	Avatar string
	SupplyRefer string
	CategoryRefer string
}

type ProductCategory struct {
	Entity
	Name string
	Description string
	ProductSubCategories []ProductCategory `gorm:"ForeignKey:ProductSubCategoriesRefers"`
	ProductSubCategoriesRefers string
}

/*Billing entity*/

type PaymentMethod struct {
	Entity
	Name string
	Description string
}

type PaymentInfo struct {
	Entity
	PaymentType PaymentMethod
	Description string
	Payer User
}

type StatusOfShipping struct{
	Entity
	Status string
	DateChangeStatus time.Time
	InProcess bool
}

type ShippingInfo struct {
	Entity
	ListStatus []StatusOfShipping
	AddressShipping Address
	Receiver User
}

type StatusOfBill struct {
	Entity
	Name string
	DateChangeStatus time.Time
	InProcess bool
}

type BillList struct {
	Entity
	Buyer User
	Seller Store
	Payment PaymentInfo
	ShippingInf ShippingInfo
	Status StatusOfBill
	BillItems []BillItem
	TotalPrice float64
}

type BillItem struct {
	Entity
	Pro Product
	NoInBill int
	TotalPrice float64
}

/*Entity to React of User with other entity*/

type Comment struct {
	Entity
	Title string
	Decription string
	UserComment User
	TypeEntity string
	EntityId uint
}

type Reaction struct {
	Entity
	Title string
	Emotion string
	UserReaction User
	TypeEntity string
	EntityId uint
}

type Rating struct {
	Entity
	Point int
	UserCreate User
	TypeRating TypeRatingCategory
	TypeEntity string
	EntityId uint
}

type TypeRatingCategory struct {
	Entity
	Title string
	UserCreate User
	TypeEntity string
	EntityId uint
}

/*Multi language*/

type Language struct {
	Entity
	Name string
}

type LanguageContent struct {
	Entity
	Content string
	Lang Language
}

type LanguageContentEntity struct {
	Entity
	ListContent []LanguageContent
	TypeEntity string
	EntityId uint
}
