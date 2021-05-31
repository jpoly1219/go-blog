# go-blog

A blogging web application for practicing web development with Go.

## Features
- The home page lists all blog posts.
- Each post can be viewed on its own page when clicked on.
- Every user has their own profile page.
- The platform has sign up and log in capabilities.
  - User authentication with JWT
  - Users who aren't logged in can view posts, but cannot write new posts.
  - Only users who are logged in can view their profile pages.
- GET retrieves posts.
- POST writes new posts.
- PUT updates posts.
- DELETE deletes posts.