package config

import (
	"LemmyPiefedApi/controller"
	"LemmyPiefedApi/router"
)

func newRoute(path string, method router.HttpMethod, controller router.ControllerMethod) *router.Route {
	return router.NewRoute("/api/v3"+path, method, controller)
}

func init() {
	userController := controller.NewUserController(piefed)
	siteController := controller.NewSiteController(piefed, activityPub, simulateLemmy)
	postController := controller.NewPostController(piefed)
	commentController := controller.NewCommentController(piefed)
	communityController := controller.NewCommunityController(piefed)
	searchController := controller.NewSearchController(piefed)

	// implemented
	AppRouter.AddRoute(newRoute("/user/login", router.HttpMethodPost, userController.Login))
	AppRouter.AddRoute(newRoute("/user/unread_count", router.HttpMethodGet, userController.GetUnreadCount))
	AppRouter.AddRoute(newRoute("/site", router.HttpMethodGet, siteController.Site))
	AppRouter.AddRoute(newRoute("/post/list", router.HttpMethodGet, postController.GetPosts))
	AppRouter.AddRoute(newRoute("/post", router.HttpMethodGet, postController.GetPost))
	AppRouter.AddRoute(newRoute("/post/like", router.HttpMethodPost, postController.LikePost))
	AppRouter.AddRoute(newRoute("/comment/list", router.HttpMethodGet, commentController.GetComments))
	AppRouter.AddRoute(newRoute("/comment", router.HttpMethodGet, commentController.GetComment))
	AppRouter.AddRoute(newRoute("/comment", router.HttpMethodPost, commentController.CreateComment))
	AppRouter.AddRoute(newRoute("/comment/like", router.HttpMethodPost, commentController.LikeComment))
	AppRouter.AddRoute(newRoute("/community", router.HttpMethodGet, communityController.GetCommunity))
	AppRouter.AddRoute(newRoute("/community/list", router.HttpMethodGet, communityController.GetCommunities))
	AppRouter.AddRoute(newRoute("/search", router.HttpMethodGet, searchController.Search))

	// impossible to implement, error pages only
	AppRouter.AddRoute(newRoute("/user/register", router.HttpMethodPost, userController.Register))
	AppRouter.AddRoute(newRoute("/user/report_count", router.HttpMethodGet, userController.GetReportCount))
	AppRouter.AddRoute(newRoute("/post/mark_as_read", router.HttpMethodPost, postController.MarkPostAsRead))
}
