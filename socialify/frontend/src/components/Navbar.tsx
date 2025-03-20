import React from 'react';
import { AppBar, Toolbar, Typography, Button, Box } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';

const Navbar: React.FC = () => {
  return (
    <AppBar position="static" sx={{ mb: 3 }}>
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          Socialify
        </Typography>
        <Box>
          <Button 
            color="inherit" 
            component={RouterLink} 
            to="/"
            sx={{ mr: 1 }}
          >
            Dashboard
          </Button>
          <Button 
            color="inherit" 
            component={RouterLink} 
            to="/top-users"
            sx={{ mr: 1 }}
          >
            Top Users
          </Button>
          <Button 
            color="inherit" 
            component={RouterLink} 
            to="/trending"
            sx={{ mr: 1 }}
          >
            Trending
          </Button>
          <Button 
            color="inherit" 
            component={RouterLink} 
            to="/feed"
          >
            Feed
          </Button>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar; 