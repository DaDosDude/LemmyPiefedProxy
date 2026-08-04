package controller

import (
	lemmyModel "LemmyBeProxy/dto/model/lemmy"
	piefedModel "LemmyBeProxy/dto/model/piefed"
	"LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/helper/converter"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service"
	"LemmyBeProxy/service/piefed"
	goHttp "net/http"
)

type SiteController struct {
	piefed        *piefed.Piefed
	activityPub   *service.ActivityPub
	simulateLemmy bool
}

func NewSiteController(
	piefed *piefed.Piefed,
	activityPub *service.ActivityPub,
	simulateLemmy bool,
) *SiteController {
	return &SiteController{
		piefed:        piefed,
		simulateLemmy: simulateLemmy,
		activityPub:   activityPub,
	}
}

func (receiver *SiteController) Site(request *http.Request) (*http.Response, error) {
	resp, err := receiver.piefed.Site(request.Headers)
	if err != nil {
		return nil, err
	}

	apActor, err := receiver.activityPub.FetchActor(resp.Site.ActorId)
	if err != nil {
		return nil, err
	}

	var version string
	if receiver.simulateLemmy {
		version = "0.19.11"
	} else {
		version = resp.Version
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmy.GetSiteResponse{
			Admins:       helper.MapSlice(resp.Admins, converter.ConvertPersonView),
			AllLanguages: helper.MapSlice(resp.Site.AllLanguages, converter.ConvertLanguageView),
			BlockedUrls:  []lemmyModel.LocalSiteUrlBlocklist{},
			CustomEmojis: []lemmyModel.CustomEmojiView{},
			DiscussionLanguages: helper.MapSlice(resp.Site.AllLanguages, func(in piefedModel.LanguageView) uint {
				return in.Id
			}),
			MyUser:   converter.ConvertMyUserInfo(resp.MyUser, apActor),
			SiteView: converter.ConvertSiteToView(&resp.Site, apActor),
			Taglines: []lemmyModel.Tagline{},
			Version:  version,
		},
	}, nil
}
