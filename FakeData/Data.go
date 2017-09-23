package FakeData

import (
	"github.com/truongtu268/OAuthServer/Model"
	"github.com/araddon/dateparse"
)

var Users = []Model.User{
	Model.User{
		Name:  "TruongTu",
		Email: "truongtu268@gmail.com",
		Profile: Model.Profile{
			Avatar:      "Tú đẹp trai",
			DateOfBirth: dateparse.MustParse("1992/8/26"),
			Country:     "VietNam",
		},
		SecurityInfo: Model.SecurityInfo{
			UserName: "truongtu268",
			Password: "truongtu268@",
		},
		BillAddress: []Model.Address{
			Model.Address{
				City:     "HoChiMinh",
				District: "TanPhu",
				Street:   "Nguyen Hong Dao",
				HouseNo:  "33",
			},
			Model.Address{
				City:     "HoChiMinh",
				District: "TanPhu",
				Street:   "Nguyen Hong Dao",
				HouseNo:  "33",
			},
		},
		ShippingAddress: []Model.Address{
			Model.Address{
				City:     "HoChiMinh",
				District: "TanPhu",
				Street:   "Nguyen Hong Dao",
				HouseNo:  "33",
			},
		},
	},
	Model.User{
		Name:  "TruongTuan",
		Email: "truongtuan268@gmail.com",
		Profile: Model.Profile{
			Avatar:      "Tuan đẹp trai",
			DateOfBirth: dateparse.MustParse("1987/12/23"),
			Country:     "VietNam",
		},
		SecurityInfo: Model.SecurityInfo{
			UserName: "truongtuan268",
			Password: "truongtuan268@",
		},
		BillAddress: []Model.Address{
			Model.Address{
				City:     "HoChiMinh",
				District: "TanPhu",
				Street:   "Nguyen Hong Dao",
				HouseNo:  "33",
			},
		},
		ShippingAddress: []Model.Address{
			Model.Address{
				City:     "HoChiMinh",
				District: "TanPhu",
				Street:   "Nguyen Hong Dao",
				HouseNo:  "33",
			},
		},
	},
}

var Stores = []Model.Store{
	Model.Store{
		Phone: "01228918306",
		AddressStore: Model.Address{
			City:     "HoChiMinh",
			District: "TanPhu",
			Street:   "Nguyen Hong Dao",
			HouseNo:  "33",
		},
		AboutStore: "Simple store to sell anything",
		WareHouseOfStore: Model.WareHouse{
			AddressWare: Model.Address{
				City:     "HoChiMinh",
				District: "TanPhu",
				Street:   "Nguyen Hong Dao",
				HouseNo:  "33",
			},
			Inventories:[]Model.Inventory{
				Model.Inventory{
					ProductIn:Model.Product{
						Name:"Bột giặt omo",
						Avatar:"Giặt trắng sạch",
						Information:"Sạch sành sanh",
						Category:Model.ProductCategory{
							Name:"Bột giặt",
							Description:"Giặt đồ",
						},
					},
					MinimumBalance:4,
					NoInStore:40,
					PriceIn:45.00,
					PriceOut:50.00,
				},
				Model.Inventory{
					ProductIn:Model.Product{
						Name:"Nước suối",
						Avatar:"Nước sạch",
						Information:"Nước mát",
						Category:Model.ProductCategory{
							Name:"Nước uống",
							Description:"Nước mát",
						},
					},
					MinimumBalance:4,
					NoInStore:40,
					PriceIn:5.00,
					PriceOut:10.00,
				},
				Model.Inventory{
					ProductIn:Model.Product{
						Name:"Nước tăng lực",
						Avatar:"Nước tăng lực",
						Information:"Nước tăng lực",
						Category:Model.ProductCategory{
							Name:"Nước tăng lực",
							Description:"Nước mát",
						},
					},
					MinimumBalance:4,
					NoInStore:40,
					PriceIn:5.00,
					PriceOut:10.00,
				},
				Model.Inventory{
					ProductIn:Model.Product{
						Name:"Bột giặt omo",
						Avatar:"Giặt trắng sạch",
						Information:"Sạch sành sanh",
						Category:Model.ProductCategory{
							Name:"Bột giặt",
							Description:"Giặt đồ",
						},
					},
					MinimumBalance:4,
					NoInStore:40,
					PriceIn:45.00,
					PriceOut:50.00,
				},
				Model.Inventory{
					ProductIn:Model.Product{
						Name:"Nước suối",
						Avatar:"Nước sạch",
						Information:"Nước mát",
						Category:Model.ProductCategory{
							Name:"Nước uống",
							Description:"Nước mát",
						},
					},
					MinimumBalance:4,
					NoInStore:40,
					PriceIn:5.00,
					PriceOut:10.00,
				},
				Model.Inventory{
					ProductIn:Model.Product{
						Name:"Nước suối",
						Avatar:"Nước sạch",
						Information:"Nước mát",
						Category:Model.ProductCategory{
							Name:"Nước uống",
							Description:"Nước mát",
						},
					},
					MinimumBalance:4,
					NoInStore:40,
					PriceIn:5.00,
					PriceOut:10.00,
				},
			},
		},
		Admin: Model.User{
			Name:  "thienkim",
			Email: "nguyenthienkim96@gmail.com",
			Profile: Model.Profile{
				Avatar:      "Thien kim đẹp gai",
				DateOfBirth: dateparse.MustParse("1996/08/26"),
				Country:     "VietNam",
			},
			SecurityInfo: Model.SecurityInfo{
				UserName: "thienkim268",
				Password: "thienkim268@",
			},
			BillAddress: []Model.Address{
				Model.Address{
					City:     "HoChiMinh",
					District: "District 7",
					Street:   "Tran Xuan Soan",
					HouseNo:  "793/35D",
				},
			},
			ShippingAddress: []Model.Address{
				Model.Address{
					City:     "HoChiMinh",
					District: "District 7",
					Street:   "Tran Xuan Soan",
					HouseNo:  "793/35D",
				},
			},
		},
	},
}

var PaymentMethod = []Model.PaymentMethod{
	Model.PaymentMethod{
		Name:        "COD",
		Description: "Nhận hàng trả tiền",
	},
	Model.PaymentMethod{
		Name:        "Credit card",
		Description: "Trả qua thẻ tín dụng",
	},
}
