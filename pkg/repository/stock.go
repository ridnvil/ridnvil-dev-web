package repository

import (
	"ridnvil-dev/pkg/database"
)

type StockUom struct {
	Iduom            int     `json:"iduom" gorm:"type:int(11);primaryKey;autoIncrement;not null"`
	Itemid           string  `json:"itemid" gorm:"type:varchar(30);not null"`
	Uom              string  `json:"uom" gorm:"type:varchar(30);not null"`
	Rate             float64 `json:"rate" gorm:"type:double;default:0"`
	Cost             float64 `json:"cost" gorm:"type:double;default:0"`
	Price            float64 `json:"price" gorm:"type:double;default:0"`
	Marketprice      float64 `json:"marketprice" gorm:"type:double;default:0"`
	Percentageprofit float64 `json:"percentageprofit" gorm:"type:double;default:0"`
	Profit           float64 `json:"profit" gorm:"type:double;default:0"`
	Minprice         float64 `json:"minprice" gorm:"type:double;default:0"`
	Parent           string  `json:"parent" gorm:"type:varchar(4);default:null"`
}

func (*StockUom) TableName() string {
	return "stockuom"
}

type StockItem struct {
	Itemid       string  `json:"itemid" gorm:"type:varchar(30);not null"`
	Description  string  `json:"description" gorm:"type:varchar(150)"`
	Description2 string  `json:"description2" gorm:"type:text"`
	Stockid      string  `json:"stockid" gorm:"type:varchar(20)"`
	Uomid        int     `json:"uomid" gorm:"type:int(11);"`
	Shelf        string  `json:"shelf" gorm:"type:varchar(40)"`
	Reorderlevel int     `json:"reorderlevel" gorm:"default:0"`
	Reorderqty   int     `json:"reorderqty" gorm:"default:1"`
	Leadtime     int     `json:"leadtime" gorm:"default:0"`
	Outputtaxid  string  `json:"outputtaxid" gorm:"type:varchar(20)"`
	Inputtaxid   string  `json:"inputtaxid" gorm:"type:varchar(20)"`
	Remark1      string  `json:"remark1" gorm:"type:varchar(150)"`
	Remark2      string  `json:"remark2" gorm:"type:varchar(150)"`
	Barcode      string  `json:"barcode" gorm:"type:varchar(150)"`
	Balanceqty   int     `json:"balanceqty" gorm:"default:0"`
	Serial       bool    `json:"serial" gorm:"type:tinyint(1);default:0"`
	Control      bool    `json:"control" gorm:"type:tinyint(1);default:1"`
	Active       bool    `json:"active" gorm:"type:tinyint(1);default:1"`
	Safetystock  float64 `json:"safetystock" gorm:"type:double;default:0"`
}

func (*StockItem) TableName() string {
	return "stockitem"
}

type StockItemResponse struct {
	Itemid       string     `json:"itemid"`
	Description  string     `json:"description"`
	Description2 string     `json:"description2"`
	Shelf        string     `json:"shelf"`
	Reorderlevel int        `json:"reorderlevel"`
	Reorderqty   int        `json:"reorderqty"`
	Leadtime     int        `json:"leadtime"`
	Outputtaxid  string     `json:"outputtaxid"`
	Inputtaxid   string     `json:"inputtaxid"`
	Remark1      string     `json:"remark1"`
	Remark2      string     `json:"remark2"`
	Barcode      string     `json:"barcode"`
	Balanceqty   int        `json:"balanceqty"`
	Serial       bool       `json:"serial"`
	Control      bool       `json:"control"`
	Active       bool       `json:"active"`
	Safetystock  float64    `json:"safetystock"`
	StockUomData []StockUom `json:"stock_uom_data"`
}

func GetListStock() ([]StockItemResponse, error) {
	db, err := database.OpenConnectionSQLite()
	if err != nil {
		return []StockItemResponse{}, err
	}
	var stocks []StockItem

	if errget := db.Model(&StockItem{}).Find(&stocks).Error; errget != nil {
		return []StockItemResponse{}, errget
	}

	var stockResponse []StockItemResponse
	for _, stock := range stocks {
		var stockuom []StockUom
		db.Model(&StockUom{}).Where("itemid = ?", stock.Itemid).Find(&stockuom)
		stockitemresponse := ConvertToStockItemResponse(stock, stockuom)
		stockResponse = append(stockResponse, stockitemresponse)
	}

	return stockResponse, nil
}

func ConvertToStockItemResponse(stockItem StockItem, stockuom []StockUom) StockItemResponse {
	return StockItemResponse{
		Itemid:       stockItem.Itemid,
		Description:  stockItem.Description,
		Description2: stockItem.Description2,
		Shelf:        stockItem.Shelf,
		Reorderlevel: stockItem.Reorderlevel,
		Reorderqty:   stockItem.Reorderqty,
		Leadtime:     stockItem.Leadtime,
		Outputtaxid:  stockItem.Outputtaxid,
		Inputtaxid:   stockItem.Inputtaxid,
		Remark1:      stockItem.Remark1,
		Remark2:      stockItem.Remark2,
		Barcode:      stockItem.Barcode,
		Balanceqty:   stockItem.Balanceqty,
		Serial:       stockItem.Serial,
		Control:      stockItem.Control,
		Active:       stockItem.Active,
		Safetystock:  stockItem.Safetystock,
		StockUomData: stockuom,
	}
}
