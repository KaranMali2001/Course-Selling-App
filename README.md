# Course Purchasing Website Backend with Golang

This is a backend-only project for a course purchasing website implemented in Golang. The website allows admins to create courses and users to purchase them. Both admins and users can view available courses without authentication.

## Features

- **Admin Features:**
  - Create new courses.
  - View list of all courses.

- **User Features:**
  - Browse available courses.
  - Purchase courses.
  - View purchased courses history.

- **General Features:**
  - Authentication using JWT.
  - Data stored in MongoDB.
  - API endpoints for admin and user actions.

## Technologies Used

- Golang
- Echo framework
- MongoDB

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your/repository.git
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables:

   - `JWT_SECRET`: Secret key for JWT authentication.
   - `DB_URL`: MongoDB connection string.

4. Run the application:

   ```bash
   go run main.go
   ```

5. The application should now be running on `localhost:8080`.

## API Endpoints

### Admin Routes

- `POST /admin/signup`: Admin sign up.
- `POST /admin/createcourse`: Create a new course (Admin only).
- `GET /admin/course`: Get all courses (Admin only).

### User Routes

- `GET /user/course`: Get all courses.
- `GET /user/course/:id`: Purchase a course by ID.
- `POST /user/usersignup`: User sign up.
- `GET /user/PurchasedCourse`: Get purchased courses (User only).

## Authentication

- Authentication is required for creating courses and purchasing courses.
- Include a JWT token in the `Authorization` header with the value `Bearer <your_token>`.
