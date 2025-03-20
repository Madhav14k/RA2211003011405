# Social Media Analytics

This project consists of a social media analytics platform with a Go backend and React frontend.

## Project Structure

```
└── socialify/
    ├── backend/         # Go backend
    │   ├── handlers/    # API handlers
    │   ├── models/      # Data models
    │   └── utils/       # Utility functions
    └── frontend/        # React frontend
        ├── public/      # Static assets
        └── src/         # Source code
            ├── components/  # React components
            ├── services/    # API services
            └── types/       # TypeScript types
```

## Backend

The backend is built with Go using the Gin framework. It provides APIs for:

- Authentication with the test server
- Getting the top users by post count
- Getting the latest posts
- Getting the most popular posts (with the most comments)

## Frontend

The frontend is built with React and TypeScript. It includes:

- A Top Users page showing users with the most posts
- A Trending Posts page showing posts with the most comments
- A Feed page showing the latest posts

## Setup and Running

### Backend

1. Navigate to the backend directory:
   ```
   cd backend
   ```

2. Set the following environment variables:
   ```
   CLIENT_ID=your_client_id
   CLIENT_SECRET=your_client_secret
   COMPANY_NAME=your_company_name
   OWNER_NAME=your_name
   OWNER_EMAIL=your_email
   ROLL_NO=your_roll_no
   ```

3. Run the backend:
   ```
   go run main.go
   ```

### Frontend

1. Navigate to the frontend directory:
   ```
   cd frontend
   ```

2. Install dependencies:
   ```
   npm install
   ```

3. Run the frontend:
   ```
   npm start
   ```

4. Open [http://localhost:3000](http://localhost:3000) in your browser.

## API Integration

The frontend communicates with the backend through RESTful APIs:

- GET /api/users/top - Get top users by post count
- GET /api/posts/latest - Get latest posts
- GET /api/posts/popular - Get posts with most comments 