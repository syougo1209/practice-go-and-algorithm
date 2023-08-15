type Tweet struct {
	count      int
	tweetId    int
	followeeId int
	index      int
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Twitter struct {
	cnt       int
	followmap map[int]map[int]bool
	tweetMap  map[int][]*Tweet
}

func Constructor() Twitter {
	return Twitter{cnt: 0, followMap: map[int]map[int]bool{}, tweetMap: map[int][]*Tweet{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := tweetMap[userId]; !ok {
		tweetMap[userId] = []*Tweet{}
	}
	index := len(tweetMap[userId])
	tweetMap[userId] = append(tweetMap[userId], &Tweet{count: this.cnt, tweetId: tweetId, followeeId: userId, index: index})
	this.cnt--
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)
	if _, ok := this.followMap[userId]; !ok {
		this.followMap[userId] = map[int]bool{}
	}
	this.followMap[userId][userId] = true

	minHeap := TweetHeap{}
	for followeeId := range this.followMap[userId] {
		if _, ok := this.TweetMap[followeeId]; ok {
			heap.Push(&minHeap, this.TweetMap[followeeId][len(this.TweetMap[followeeId])-1])
		}
	}
	for len(heap) > 0 && len(res) < 10 {
		tweet := heap.Pop(&minHeap).(*Tweet)
		if tweet.index > 0 {
			heap.Push(this.tweetMap[tweet.followeeId][tweet.index-1])
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := followmap[followerId]; !ok {
		followmap[followerId] = map[int]bool{}
	}
	followmap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(followmap[followerId], followeeId)
}
