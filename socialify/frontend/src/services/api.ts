import axios from 'axios';
import { UserPostCount, PostWithUser } from '../types';

const API_URL = 'http://localhost:8081/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getTopUsers = async (): Promise<UserPostCount[]> => {
  try {
    const response = await api.get('/users/top');
    return response.data.topUsers || [];
  } catch (error) {
    console.error('Error fetching top users:', error);
    return [];
  }
};

export const getLatestPosts = async (): Promise<PostWithUser[]> => {
  try {
    const response = await api.get('/posts/latest');
    return response.data.latestPosts || [];
  } catch (error) {
    console.error('Error fetching latest posts:', error);
    return [];
  }
};

export const getPopularPosts = async (): Promise<PostWithUser[]> => {
  try {
    const response = await api.get('/posts/popular');
    return response.data.popularPosts || [];
  } catch (error) {
    console.error('Error fetching popular posts:', error);
    return [];
  }
}; 