class Twitter:

    def __init__(self):
        self.users: Dict[userId, User] = dict()
        self.tweetTime = 0
        # list[[userId, tweet, tweetTime]]
    
    def postTweet(self, userId: int, tweetId: int) -> None:
        user = self.users.get(userId)
        if not user:
            newUser = User(userId)
            self.users[userId] = newUser
            user = newUser
         
        user.tweets.append((self.tweetTime, tweetId))
        self.tweetTime += 1
        

    def getNewsFeed(self, userId: int) -> List[int]:
        user = self.users.get(userId)
        if not user:
            newUser = User(userId)
            self.users[userId] = newUser
            user = newUser

        tweets = user.tweets.copy()

        for followingId in user.followings:
            curr = self.users[followingId]
            if len(curr.tweets) < 1:
                continue
            curr_tweets = curr.tweets.copy() 
            tweets.extend(curr_tweets)
        heapq.heapify_max(tweets)
        newsFeed = []
        while len(tweets) != 0 and len(newsFeed) < 10:
            newsFeed.append(heapq.heappop_max(tweets))  
        res = [item[1] for item in newsFeed]
        return res

    def follow(self, followerId: int, followeeId: int) -> None:
        follower, followee = self.users.get(followerId), self.users.get(followeeId)
        if not follower:
            self.users[followerId] = User(followerId)
            follower = self.users[followerId]
        if not followee:
            self.users[followeeId] = User(followeeId) 
            followee = self.users[followeeId]
    
        follower.followings.add(followeeId)
        followee.followers.add(followerId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        follower, followee = self.users.get(followerId), self.users.get(followeeId)
        if not follower:
            self.users[followerId] = User(followerId)
            follower = self.users[followerId]
        if not followee:
            self.users[followeeId] = User(followeeId) 
            followee = self.users[followeeId]       
            
        follower, followee = self.users[followerId], self.users[followeeId]
        follower.followings.discard(followeeId)
        followee.followers.discard(followerId)

class User:
    def __init__(self, userId):
        self.userId = userId
        self.tweets = []
        self.followings: set(int) = set()
        self.followers: set(int) = set()