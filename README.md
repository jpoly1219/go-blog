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
  - Users who log in have their form data checked by comparing it to the database records, then receives a JWT access token if authenticated.
  - Users also receive refresh tokens that last for 7 days. This allows the users to stay logged in even if their access token has been invalidated after 15 minutes.
  - Protected endpoints require a valid access token to access.
  - Each token has a uuid, allowing one user to be logged into multiple devices at a time.

## Features to work on
- User authentication
  - If the user decides to log in from different devices, then the access token for one device has to be shared across other devices. This might be a better design than having to generate x amount of tokens for each device.
- Security
  - There is no way to invalidate the access tokens if a hacker successfully steals them. The hacker has full access for 15 minutes until the token is invalidated.
- The blog frontend
  - They say that the API is only as good as the frontend that consumes it...