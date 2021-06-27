package constant

// Constant
var (
	PostStatusPending   = "pending"
	PostStatusTimer     = "timer"
	PostStatusCanceled  = "canceled"
	PostStatusRemoved   = "removed"
	PostStatusPublished = "published"
	PostStatusList      = []interface{}{
		PostStatusPending, PostStatusTimer, PostStatusCanceled, PostStatusRemoved, PostStatusPublished,
	}
)
