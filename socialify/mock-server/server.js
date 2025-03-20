const express = require('express');
const cors = require('cors');
const app = express();
const port = 8080;

app.use(cors());
app.use(express.json());

// Mock data
const users = {
  "1": "John Doe",
  "2": "Jane Doe",
  "3": "Alice Smith",
  "4": "Bob Johnson",
  "5": "Charlie Brown",
  "6": "Diana White",
  "7": "Edward Davis",
  "8": "Fiona Miller",
  "9": "George Wilson",
  "10": "Helen Moore",
};

const posts = [
  { id: 246, userid: "1", content: "Post about ant" },
  { id: 161, userid: "1", content: "Post about elephant" },
  { id: 150, userid: "1", content: "Post about ocean" },
  { id: 370, userid: "1", content: "Post about monkey" },
  { id: 344, userid: "1", content: "Post about ocean" },
  { id: 952, userid: "1", content: "Post about zebra" },
  { id: 647, userid: "1", content: "Post about igloo" },
  { id: 421, userid: "1", content: "Post about house" },
  { id: 890, userid: "1", content: "Post about bat" },
  { id: 461, userid: "1", content: "Post about umbrella" },
  { id: 247, userid: "2", content: "Post about flowers" },
  { id: 162, userid: "2", content: "Post about gardens" },
  { id: 151, userid: "2", content: "Post about rivers" },
  { id: 371, userid: "2", content: "Post about mountains" },
  { id: 345, userid: "3", content: "Post about hiking" },
  { id: 953, userid: "3", content: "Post about camping" },
  { id: 648, userid: "4", content: "Post about cooking" },
  { id: 422, userid: "4", content: "Post about baking" },
  { id: 891, userid: "5", content: "Post about music" },
  { id: 462, userid: "5", content: "Post about art" },
];

const comments = {
  150: [
    { id: 3893, postid: 150, content: "Old comment" },
    { id: 4791, postid: 150, content: "Boring comment" },
    { id: 4792, postid: 150, content: "Interesting comment" },
  ],
  161: [
    { id: 3894, postid: 161, content: "Nice post" },
    { id: 4793, postid: 161, content: "Great observation" },
  ],
  246: [
    { id: 3895, postid: 246, content: "I agree" },
  ],
  370: [
    { id: 3896, postid: 370, content: "Funny post" },
    { id: 4794, postid: 370, content: "LOL" },
    { id: 4795, postid: 370, content: "ROFL" },
  ],
};

// Routes
app.get('/api/users', (req, res) => {
  res.json({ users });
});

app.get('/api/users/top', (req, res) => {
  const userPostCounts = Object.keys(users).map(id => {
    const userPosts = posts.filter(post => post.userid === id);
    return {
      user: {
        id,
        name: users[id]
      },
      postCount: userPosts.length
    };
  }).sort((a, b) => b.postCount - a.postCount).slice(0, 5);

  res.json({ topUsers: userPostCounts });
});

app.get('/api/users/:userId/posts', (req, res) => {
  const userId = req.params.userId;
  const userPosts = posts.filter(post => post.userid === userId);
  res.json({ posts: userPosts });
});

app.get('/api/posts/latest', (req, res) => {
  const latestPosts = [...posts]
    .sort((a, b) => b.id - a.id)
    .slice(0, 5)
    .map(post => ({
      post,
      user: {
        id: post.userid,
        name: users[post.userid]
      }
    }));

  res.json({ latestPosts });
});

app.get('/api/posts/popular', (req, res) => {
  const postCommentCounts = posts.map(post => {
    const postComments = comments[post.id] || [];
    return {
      post,
      user: {
        id: post.userid,
        name: users[post.userid]
      },
      commentCount: postComments.length
    };
  }).sort((a, b) => b.commentCount - a.commentCount);

  let maxCount = 0;
  if (postCommentCounts.length > 0) {
    maxCount = postCommentCounts[0].commentCount;
  }

  const popularPosts = postCommentCounts.filter(p => p.commentCount === maxCount);

  res.json({ popularPosts });
});

app.get('/api/posts/:postId/comments', (req, res) => {
  const postId = parseInt(req.params.postId);
  const postComments = comments[postId] || [];
  res.json({ comments: postComments });
});

app.listen(port, () => {
  console.log(`Mock server running at http://localhost:${port}`);
}); 