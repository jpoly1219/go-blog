# go-blog

A blogging platform written in Go. Developed for practicing web development with Go. Frontend is written in Svelte.
The project is meant to be kept bare-bones and minimal. I am trying not to use frameworks for the backend. Just using some libraries to help implement certain parts such as the router. Currently the backend uses these libraries:  
- The Go Standard Library
- gorilla/mux
- go-sql-driver/mysql
- google/uuid
- dgrijalva/jwt-go
- joho/godotenv

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
  - Each token has a UUID, allowing one user to be logged into multiple devices at a time.
- Where to store the JWT
  - Access token is stored in memory as a Svelte store, while the refresh token is stored in a HttpOnly cookie.
  - Backend handles everything CORS, and this allows cookies to be sent over cross origin (same domain, but different ports).
- Frontend
  - A basic blogging platform frontend built with Svelte.
  - I thought Svelte fits the minimalist theme quite well, and works pretty well with my Go backend. Svelte also makes my frontend insanely fast.
  - I may or may not switch over to SvelteKit, but I want to stick with Svelte for the time being.
    - Frontend now starts a timer that counts until the access token expiration time, and when it does, it sends a request to the backend for a refresh token.
  - Clicking on a post title links to a page dedicated for that single post.
  - Clicking on the username in the navbar after login redirects to the user profile page.
  - Each post component shows the title, author and the content of the post.
  - Users can write their own posts.
  - Posts are shortened in the homepage if they exceed around 100 words or so.
- Security
  - Passwords are hashed using bcrypt.

## Features to work on
- Frontend
  - Pagination
    - It's a bad idea to load every post from the database, so there needs to be a way to load x amount of posts at a time.
    - Have an infinitely scrolling page that loads posts as the user scrolls down.
    - MySQL LIMIT, GROUP BY keywords can be handy.
    - Have a store called `page`, and whenever the user scrolls to the bottom of the screen, increment this value by 1.
    - This value will be sent back to the backend to load the next batch of posts.
      - Each batch consists of 10 posts.
      - SELECT * FROM posts ORDER BY id DESC LIMIT 0, 10;
      - SELECT * FROM posts ORDER BY id DESC LIMIT 10, 10; ...
      - If the range becomes out of bound (eg LIMIT 10, 10 but there are only two elements left), MySQL just stops after selecting what's left.

- User authentication
  - If the user decides to log in from different devices, then the access token for one device has to be shared across other devices. This might be a better design than having to generate x amount of tokens for each device.
- Security
  - There is no way to invalidate the access tokens if a hacker successfully steals them. The hacker has full access for 15 minutes until the token is invalidated.