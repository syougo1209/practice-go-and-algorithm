
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
	followMap map[int]map[int]bool
	tweetMap  map[int][]*Tweet
	count     int
}
type Tweet struct {
	tweetId int
	userId  int
	count   int
	index   int
}

func Constructor() Twitter {
	return Twitter{followMap: map[int]map[int]bool{}, tweetMap: map[int][]*Tweet{}, count: 0}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.tweetMap[userId]; !ok {
		this.tweetMap[userId] = make([]*Tweet, 0)
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], &Tweet{count: this.count, tweetId: tweetId})
	this.count -= 1
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	if _, ok := this.followMap[userId]; !ok {
		this.followMap[userId] = map[int]bool{}
	}
	this.followMap[userId][userId] = true

	minHeap := TweetHeap{}
	for followUserId := range this.followMap[userId] {
		if _, ok := this.tweetMap[followUserId]; ok {
			lastIndex := len(this.tweetMap[followUserId]) - 1
			heap.Push(&minHeap, this.tweetMap[followUserId][lastIndex])
		}
	}
	for len(res) < 10 && len(minHeap) != 0 {
		tweet := heap.Pop(&minHeap).(*Tweet)
		res = append(res, tweet.tweetId)
		if tweet.index >= 1 {
			heap.Push(&minHeap, this.tweetMap[tweet.userId][tweet.index-1])
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.followMap[followerId]; !ok {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.followMap[followerId]; ok {
		delete(this.followMap[followerId], followeeId)
	}
}
