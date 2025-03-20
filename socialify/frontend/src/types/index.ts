export interface User {
  id: string;
  name: string;
}

export interface Post {
  id: number;
  userid: string;
  content: string;
}

export interface Comment {
  id: number;
  postid: number;
  content: string;
}

export interface UserPostCount {
  user: User;
  postCount: number;
}

export interface PostWithUser {
  post: Post;
  user: User;
  commentCount?: number;
}

export interface ApiResponse<T> {
  data: T;
  error?: string;
} 