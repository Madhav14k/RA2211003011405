import React, { useEffect, useState } from 'react';
import { Typography, Box, CircularProgress, Grid, Paper, Avatar } from '@mui/material';
import { Person } from '@mui/icons-material';
import { getLatestPosts } from '../services/api';
import { PostWithUser } from '../types';

const Feed: React.FC = () => {
  const [posts, setPosts] = useState<PostWithUser[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchLatestPosts = async () => {
      setLoading(true);
      try {
        const data = await getLatestPosts();
        setPosts(data);
      } catch (error) {
        console.error('Error fetching latest posts:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchLatestPosts();
    
    // Set up polling for real-time updates
    const interval = setInterval(fetchLatestPosts, 10000);
    return () => clearInterval(interval);
  }, []);

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Latest Posts Feed
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
                  <Typography variant="body1">
                    {postData.post.content}
                  </Typography>
                </Paper>
              </Grid>
            ))
          ) : (
            <Grid item xs={12}>
              <Paper elevation={3} sx={{ p: 2 }}>
                <Typography>No posts found</Typography>
              </Paper>
            </Grid>
          )}
        </Grid>
      )}
    </Box>
  );
};

export default Feed; 