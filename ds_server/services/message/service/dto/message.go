package dto

import "ds_server/services/message/model"

type Banners struct {
	BannerType model.BannerType
	BannerUrl  []model.DsBannerUrl
}
