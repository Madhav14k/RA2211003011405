import React, { useEffect, useState } from 'react';
import { Container, Typography, Box, CircularProgress, Grid, Paper } from '@mui/material';
import { getTopUsers, getLatestPosts, getPopularPosts } from '../services/api';
import { UserPostCount, PostWithUser } from '../types';

const SimpleDashboard: React.FC = () => {
  const [topUsers, setTopUsers] = useState<UserPostCount[]>([]);
  const [latestPosts, setLatestPosts] = useState<PostWithUser[]>([]);
  const [popularPosts, setPopularPosts] = useState<PostWithUser[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        const [topUsersData, latestPostsData, popularPostsData] = await Promise.all([
          getTopUsers(),
          getLatestPosts(),
          getPopularPosts()
        ]);
        
        setTopUsers(topUsersData);
        setLatestPosts(latestPostsData);
        setPopularPosts(popularPostsData);
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="80vh">
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" gutterBottom sx={{ mb: 4 }}>
        Social Media Analytics Dashboard
      </Typography>

      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Top Users
            </Typography>
            <ul>
              {topUsers.map(user => (
                <li key={user.user.id}>
                  {user.user.name} - {user.postCount} posts
                </li>
              ))}
            </ul>
          </Paper>
        </Grid>

        <Grid item xs={12}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Latest Posts
            </Typography>
            <ul>
              {latestPosts.map(item => (
                <li key={item.post.id}>
                  <strong>{item.user.name}</strong>: {item.post.content}
                </li>
              ))}
            </ul>
          </Paper>
        </Grid>

        <Grid item xs={12}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Popular Posts
            </Typography>
            <ul>
              {popularPosts.map(item => (
                <li key={item.post.id}>
                  <strong>{item.user.name}</strong>: {item.post.content}
                  {item.commentCount && <span> ({item.commentCount} comments)</span>}
                </li>
              ))}
            </ul>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
};

export default SimpleDashboard; 