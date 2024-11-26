-- USERS TABLE --
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT
);
-- POSTS TABLE --
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
-- categories TABLE --
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    categorie TEXT NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(id)
);
 
-- COMMENTS TABLE --
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    user_id INTEGER,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);
-- Likes TABLE --
CREATE TABLE IF NOT EXISTS likes_dislikes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_or_comment_id INTEGER,
    user_id INTEGER,
    is_like BOOLEAN,
    is_comment BOOLEAN,
    FOREIGN KEY(post_or_comment_id) REFERENCES posts(id),
    FOREIGN KEY(post_or_comment_id) REFERENCES comments(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Session TABLE--
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);