import React, { useEffect, useState } from 'react';
import { Typography, Card, CardContent, List, ListItem, ListItemText, Avatar, ListItemAvatar, Box, CircularProgress } from '@mui/material';
import { Person } from '@mui/icons-material';
import { getTopUsers } from '../services/api';
import { UserPostCount } from '../types';

const TopUsers: React.FC = () => {
  const [users, setUsers] = useState<UserPostCount[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTopUsers = async () => {
      setLoading(true);
      try {
        const data = await getTopUsers();
        setUsers(data);
      } catch (error) {
        console.error('Error fetching top users:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchTopUsers();
  }, []);

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Top Users by Post Count
      </Typography>
      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <Card>
          <CardContent>
            <List>
              {users.length > 0 ? (
                users.map((item, index) => (
                  <ListItem key={item.user.id} divider={index < users.length - 1}>
                    <ListItemAvatar>
                      <Avatar>
                        <Person />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText
                      primary={item.user.name}
                      secondary={`${item.postCount} posts`}
                    />
                  </ListItem>
                ))
              ) : (
                <ListItem>
                  <ListItemText primary="No users found" />
                </ListItem>
              )}
            </List>
          </CardContent>
        </Card>
      )}
    </Box>
  );
};

export default TopUsers; 