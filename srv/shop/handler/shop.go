package handler

import (
	"context"
	"encoding/json"
	"github.com/prometheus/common/log"
	"go-micro-example/repository"
	shopProto "go-micro-example/srv/shop/proto/shop"
)

type ShopHandler struct {
	repo repository.IShopRepo
}

func NewShopHandler(repo repository.IShopRepo) ShopHandler {
	return ShopHandler{repo: repo}
}

func (s ShopHandler) Create(ctx context.Context, req *shopProto.ShopDetail, res *shopProto.Response) error {
	createReq := repository.ShopCreateReq{}

	createReq.MD5List = req.Md5List
	createReq.RunningPeriod = req.RunningPeriod
	createReq.AppUserId = req.AppUserId
	createReq.ShopName = req.ShopName
	createReq.Addr = req.Addr
	createReq.Lat = req.Lat
	createReq.Lon = req.Lon
	createReq.Phone = req.Phone
	createReq.IndustryId = req.IndustryId
	createReq.Industry = req.Industry
	createReq.Region = req.Region
	createReq.RegionId = req.RegionId

	shopId, err := s.repo.Create(createReq)
	if err != nil {
		log.Error("[shop] [create] repo.Create error.")
		log.Error(err)
		return err
	}

	log.Info("[shop] [create] create shop id " + string(shopId))
	res.Id = shopId

	return nil
}

func (s ShopHandler) GetList(ctx context.Context, req *shopProto.Request, res *shopProto.ShopListItem) error {
	shopList, count, err := s.repo.GetList(req.Page, req.PageSize, req.SearchKey)
	if err != nil {
		log.Error("[shop] [GetList] error.")
		log.Error(err)
		return err
	}

	for _, shop := range shopList {
		var runningPeriod []string
		var md5List []string
		json.Unmarshal(shop.RunningPeriod, &runningPeriod)
		json.Unmarshal(shop.MD5List, &md5List)

		res.Shops = append(res.Shops, &shopProto.ShopDetail{

			ShopName:      shop.ShopName,
			BsNumber:      shop.BsNumber,
			IsActive:      shop.IsActive,
			Addr:          shop.Addr,
			Lat:           shop.Lat,
			Lon:           shop.Lon,
			Phone:         shop.Phone,
			Industry:      shop.Industry,
			IndustryId:    shop.IndustryId,
			RunningPeriod: runningPeriod,
			Md5List:       md5List,
			AppUserId:     shop.AppUserId,
			RegionId:      shop.RegionId,
			Region:        shop.Region,
		})
	}

	res.Count = count

	return nil

}

func (s ShopHandler) GetById(context.Context, *shopProto.Request, *shopProto.ShopDetail) error {
	panic("implement me")
}

func (s ShopHandler) Disable(context.Context, *shopProto.Request, *shopProto.Response) error {
	panic("implement me")
}
