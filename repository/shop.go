package repository

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"go-micro-example/model"
	"go-micro-example/pkg/util"
)

type ShopCreateReq struct {
	AppUserId uint32

	ShopName      string
	Addr          string
	Lat           float32
	Lon           float32
	Phone         string
	RegionId      uint32
	Region        string
	IndustryId    uint32
	Industry      string
	RunningPeriod []string
	MD5List       []string
}

type IShopRepo interface {
	Create(shop ShopCreateReq) (shopId uint32, err error)
	GetList(page, pageSize int32, searchKey string) ([]model.ShopModel, int32, error)
	GetById(shopId uint32) (model.ShopModel, error)
	UpdateIsActive(shopId uint32, isActive bool) (bool, error)
}

type ShopRepo struct {
	db *gorm.DB
}

func NewShopRepo(db *gorm.DB) ShopRepo {
	return ShopRepo{db: db}
}

//Create 创建店铺
func (s ShopRepo) Create(shop ShopCreateReq) (shopId uint32, err error) {
	runningPeriodBytes, err := json.Marshal(shop.RunningPeriod)
	log.Info(string(runningPeriodBytes))
	if err != nil {
		log.Error(err)
		return 0, err
	}
	md5Bytes, err := json.Marshal(shop.MD5List)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	shopModel := model.ShopModel{
		ShopName:            shop.ShopName,
		RegionId:            shop.RegionId,
		Region:              shop.Region,
		Lon:                 shop.Lon,
		Lat:                 shop.Lat,
		Addr:                shop.Addr,
		Phone:               shop.Phone,
		RunningPeriod:       runningPeriodBytes,
		MD5List:             md5Bytes,
		UnitPrice:           0,
		UnitDuration:        0,
		Proportion:          0,
		ExpectIncome:        0,
		Heat:                0,
		IndustryId:          0,
		AppUserId:           shop.AppUserId,
		AppUserAccount:      "",
		DayAndNumberOfAdMap: []byte("[]"),
		NumberOfAd:          0,
		BsNumber:            0,
		IsActive:            true,
	}
	err = s.db.Create(&shopModel).Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return shopModel.ID, nil
}

//GetList 获取店铺列表
func (s ShopRepo) GetList(page, pageSize int32, searchKey string) ([]model.ShopModel, int32, error) {
	var shopList []model.ShopModel
	var count int32

	db := s.db.Model(&model.ShopModel{})
	if searchKey != "" {
		db = db.Where("contact(shop_name, app_user_account, phone, industry) like %?%", searchKey)
	}

	db.Count(&count)

	if page > 0 && pageSize > 0 {
		offset := util.GetOffset(page, pageSize)
		db = db.Offset(offset).Limit(pageSize)
	}

	err := db.Find(&shopList).Error
	if err != nil {
		return shopList, count, err
	}

	return shopList, count, nil
}

//GetById 通过店铺 ID 获取店铺信息
func (s ShopRepo) GetById(shopId uint32) (model.ShopModel, error) {

	var shop model.ShopModel
	sql := s.db
	err := sql.Where("id = ?", shopId).First(&shop).Error
	if err != nil {
		return shop, err
	}

	return shop, nil

}

//UpdateIsActive 更新店铺 禁用/启用
func (s ShopRepo) UpdateIsActive(shopId uint32, isActive bool) (bool, error) {
	sql := s.db
	err := sql.Model(&model.ShopModel{}).Updates(model.ShopModel{
		IsActive: isActive,
	}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
