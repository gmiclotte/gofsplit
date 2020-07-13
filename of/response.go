package of

type Response struct {
	ResponseType                                                                string
	Id                                                                          int
	PostedAt                                                                    string
	PostedAtPrecise                                                             string
	ExpiredAt                                                                   string
	Author                                                                      Author
	Text, RawText                                                               string
	LockedText, IsFavorite, CanReport, CanDelete, CanComment, CanEdit, IsPinned bool
	CommentsCount, FavoritesCount, MediaCount                                   int
	MediaType                                                                   string
	IsMediaReady                                                                bool
	Voting                                                                      interface{}
	IsOpened, CanToggleFavorite                                                 bool
	StreamId                                                                    string
	Price                                                                       string
	Preview                                                                     interface{}
	HasVoting, IsAddedToBookmarks, IsArchived                                   bool
	MentionedUsers, LinkedUsers, LinkedPosts                                    interface{}
	TipsAmount                                                                  string
	CanViewMedia                                                                bool
	Media                                                                       []MediumJSON
}

type MediumJSON struct {
	Id               int
	Type             string
	ConvertedToVideo bool
	CanView          bool
	Thumb            string
	Preview          string
	Full             string
	SquarePreview    string
	HasError         bool
	Info             struct {
		Source MediumInfo
	}
	Source       MediumInfo
	VideoSources map[string]string
}

type MediumInfo struct {
	Source   string
	Width    int
	Height   int
	Size     int
	Duration int
}

type Author struct {
	View         string
	Avatar       string
	AvatarThumbs map[string]string
	Header       string
	HeaderSize   struct {
		Width, Height int
	}
	HeaderThumbs                                                 map[string]string
	Id                                                           int
	Name                                                         string
	Username                                                     string
	CanLookStory, CanCommentStory, HasNotViewedStory, IsVerified bool
	CanPayInternal, HasStream, HasStories, TipsEnabled           bool
	TipsMin, TipsMax                                             int
	Bookmarked, CanBeBookmarked, CanEarn, CanAddSubscriber       bool
	SubscribePrice                                               int
	SubscriptionBundles                                          []SubscriptionBundle
	Unprofitable                                                 bool
	ListsStates                                                  []ListState
	IsMuted, SubscribedBy, SubscribedByExpire                    bool
	SubscribedByExpireDate                                       string
	SubscribedByAutoprolong, SubscribedIsExpiredNow              bool
	CurrentSubscribePrice                                        float64
	SubscribedOn                                                 bool
	SubscribedOnExpiredNow, SubscribedOnDuration                 string
	ShowPostsInFeed                                              bool
	CanTrialSend                                                 bool
}

type ListState struct {
	Id         int
	Type       string
	Name       string
	HasUser    bool
	CanAddUser bool
}

type SubscriptionBundle struct {
	Id, Discount, Duration, Price int
	CanBuy                        bool
}
