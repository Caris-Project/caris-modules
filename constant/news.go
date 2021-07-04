package constant

// Constant
var (
	NewsStatusPending   = "pending"
	NewsStatusTimer     = "timer"
	NewsStatusCanceled  = "canceled"
	NewsStatusRemoved   = "removed"
	NewsStatusPublished = "published"
	NewsStatusList      = []interface{}{
		NewsStatusPending, NewsStatusTimer, NewsStatusCanceled, NewsStatusRemoved, NewsStatusPublished,
	}

	RedisKeyNewsRecommended = "news_recommended"
)
