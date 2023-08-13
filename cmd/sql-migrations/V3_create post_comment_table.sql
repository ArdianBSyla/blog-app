-- Creates post_comment table to save comments in a post
CREATE TABLE post_comment (
    id SERIAL PRIMARY KEY,
    post_id INT REFERENCES blog_post(id) ON DELETE CASCADE,
    author TEXT NOT NULL,
    content TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN DEFAULT FALSE
);