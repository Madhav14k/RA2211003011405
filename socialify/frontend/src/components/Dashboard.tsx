import React, { useEffect, useState } from 'react';
import { 
  Container, 
  Grid, 
  Paper, 
  Typography, 
  Box, 
  CircularProgress,
  Card,
  CardContent,
  List,
  ListItem,
  ListItemText,
  ListItemAvatar,
  Avatar,
  Divider
} from '@mui/material';
import { Person, Comment, Visibility, TrendingUp } from '@mui/icons-material';
import { getTopUsers, getLatestPosts, getPopularPosts } from '../services/api';
import { UserPostCount, PostWithUser } from '../types';

const Dashboard: React.FC = () => {
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
        {/* Summary Cards */}
        <Grid item xs={12} md={4}>
          <Paper
            sx={{
              p: 2,
              display: 'flex',
              flexDirection: 'column',
              height: 140,
              bgcolor: '#e3f2fd',
            }}
          >
            <Typography component="h2" variant="h6" color="primary" gutterBottom>
              Top User
            </Typography>
            {topUsers.length > 0 ? (
              <Box sx={{ display: 'flex', alignItems: 'center' }}>
                <Avatar sx={{ bgcolor: 'primary.main', mr: 2 }}>
                  <Person />
                </Avatar>
                <Box>
                  <Typography variant="h5">{topUsers[0].user.name}</Typography>
                  <Typography variant="body2">{topUsers[0].postCount} posts</Typography>
                </Box>
              </Box>
            ) : (
              <Typography>No data available</Typography>
            )}
          </Paper>
        </Grid>
        
        <Grid item xs={12} md={4}>
          <Paper
            sx={{
              p: 2,
              display: 'flex',
              flexDirection: 'column',
              height: 140,
              bgcolor: '#fff8e1',
            }}
          >
            <Typography component="h2" variant="h6" color="secondary" gutterBottom>
              Latest Activity
            </Typography>
            {latestPosts.length > 0 ? (
              <Box sx={{ display: 'flex', alignItems: 'center' }}>
                <Avatar sx={{ bgcolor: 'secondary.main', mr: 2 }}>
                  <Visibility />
                </Avatar>
                <Box>
                  <Typography variant="h5">{latestPosts[0].user.name}</Typography>
                  <Typography variant="body2" noWrap>{latestPosts[0].post.content}</Typography>
                </Box>
              </Box>
            ) : (
              <Typography>No data available</Typography>
            )}
          </Paper>
        </Grid>
        
        <Grid item xs={12} md={4}>
          <Paper
            sx={{
              p: 2,
              display: 'flex',
              flexDirection: 'column',
              height: 140,
              bgcolor: '#e8f5e9',
            }}
          >
            <Typography component="h2" variant="h6" color="success.main" gutterBottom>
              Most Popular Post
            </Typography>
            {popularPosts.length > 0 ? (
              <Box sx={{ display: 'flex', alignItems: 'center' }}>
                <Avatar sx={{ bgcolor: 'success.main', mr: 2 }}>
                  <TrendingUp />
                </Avatar>
                <Box>
                  <Typography variant="h5">{popularPosts[0].user.name}</Typography>
                  <Typography variant="body2" noWrap>{popularPosts[0].post.content}</Typography>
                </Box>
              </Box>
            ) : (
              <Typography>No data available</Typography>
            )}
          </Paper>
        </Grid>

        {/* Top Users List */}
        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Top Users by Post Count
            </Typography>
            <List>
              {topUsers.map((user, index) => (
                <React.Fragment key={user.user.id}>
                  <ListItem>
                    <ListItemAvatar>
                      <Avatar>
                        <Person />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText 
                      primary={user.user.name} 
                      secondary={`${user.postCount} posts`} 
                    />
                  </ListItem>
                  {index < topUsers.length - 1 && <Divider variant="inset" component="li" />}
                </React.Fragment>
              ))}
            </List>
          </Paper>
        </Grid>

        {/* Latest Posts */}
        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Latest Posts
            </Typography>
            <List>
              {latestPosts.map((item, index) => (
                <React.Fragment key={item.post.id}>
                  <ListItem alignItems="flex-start">
                    <ListItemAvatar>
                      <Avatar>
                        <Person />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText
                      primary={item.user.name}
                      secondary={
                        <React.Fragment>
                          <Typography
                            component="span"
                            variant="body2"
                            color="text.primary"
                          >
                            {item.post.content}
                          </Typography>
                        </React.Fragment>
                      }
                    />
                  </ListItem>
                  {index < latestPosts.length - 1 && <Divider variant="inset" component="li" />}
                </React.Fragment>
              ))}
            </List>
          </Paper>
        </Grid>

        {/* Popular Posts */}
        <Grid item xs={12}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Most Popular Posts
            </Typography>
            <Grid container spacing={2}>
              {popularPosts.map((item) => (
                <Grid item xs={12} sm={6} md={4} key={item.post.id}>
                  <Card>
                    <CardContent>
                      <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                        <Avatar sx={{ mr: 1 }}>
                          <Person />
                        </Avatar>
                        <Typography variant="subtitle1">{item.user.name}</Typography>
                      </Box>
                      <Typography variant="body1" component="div" sx={{ mb: 2 }}>
                        {item.post.content}
                      </Typography>
                      <Box sx={{ display: 'flex', alignItems: 'center' }}>
                        <Comment fontSize="small" color="action" />
                        <Typography variant="body2" color="text.secondary" sx={{ ml: 1 }}>
                          {item.commentCount || 0} comments
                        </Typography>
                      </Box>
                    </CardContent>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
};

export default Dashboard; 