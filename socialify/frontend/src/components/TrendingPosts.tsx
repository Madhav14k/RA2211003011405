import React, { useEffect, useState } from 'react';
import { Typography, Card, CardContent, Box, CircularProgress, Grid, Paper, Avatar } from '@mui/material';
import { Comment as CommentIcon, Person } from '@mui/icons-material';
import { getPopularPosts } from '../services/api';
import { PostWithUser } from '../types';

const TrendingPosts: React.FC = () => {
  const [posts, setPosts] = useState<PostWithUser[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTrendingPosts = async () => {
      setLoading(true);
      try {
        const data = await getPopularPosts();
        setPosts(data);
      } catch (error) {
        console.error('Error fetching trending posts:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchTrendingPosts();
  }, []);

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Trending Posts (Most Comments)
      </Typography>
      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <Grid container spacing={3}>
          {posts.length > 0 ? (
            posts.map((postData) => (
              <Grid item xs={12} key={postData.post.id}>
                <Paper elevation={3} sx={{ p: 2 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <Avatar sx={{ mr: 2 }}>
                      <Person />
                    </Avatar>
                    <Typography variant="subtitle1">
                      {postData.user.name}
                    </Typography>
                  </Box>
                  <Typography variant="body1" paragraph>
                    {postData.post.content}
                  </Typography>
                  <Box sx={{ display: 'flex', alignItems: 'center' }}>
                    <CommentIcon sx={{ mr: 1, fontSize: 18 }} />
                    <Typography variant="body2">
                      {postData.commentCount} comments
                    </Typography>
                  </Box>
                </Paper>
              </Grid>
            ))
          ) : (
            <Grid item xs={12}>
              <Card>
                <CardContent>
                  <Typography>No trending posts found</Typography>
                </CardContent>
              </Card>
            </Grid>
          )}
        </Grid>
      )}
    </Box>
  );
};

export default TrendingPosts; 