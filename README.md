# go-blog

A blogging platform written in Go. Developed for practicing web development with Go.

## Features currently implemented
- REST API
  - `GET` retrieves posts.
  - `POST` writes new posts.
  - `PUT` updates posts.
  - `DELETE` deletes posts.
- MySQL database to store all posts and user info.
  - `users` table stores the user account info.
  - `posts` table stores the post data.
- User authentication
  - Users who sign up have their account data recorded to the MySQL database.
  - Users who log in have their form data checked by comparing it to the database records, then receives a JWT access  
  token if authenticated.

## Features to work on
- User authentication
  - Users need to get refresh tokens along with the access token in order to improve user experience. It would be bad UX  
  if the users had to log in every time their access token gets invalidated after 15 minutes.
  - There are currently no functions to check if the user is validated and can have access to protected endpoints.
  - If the user decides to log in from different devices, then the access token for one device has to be shared across  
  other devices. Or better yet just make more JWTs with different ID's for different devices?
- Security
  - There is no way to invalidate the access tokens if a hacker successfully steals them. The hacker has full access for  
  15 minutes until the token is invalidated.
- The blog frontend
  - They say that the API is only as good as the frontend that consumes it...