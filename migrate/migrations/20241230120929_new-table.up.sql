-- Create videos table
CREATE TABLE videos (
    id SERIAL PRIMARY KEY,
    video_id VARCHAR(255) UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    published_at TIMESTAMP WITH TIME ZONE,
    thumbnail_url VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index on published_at
CREATE INDEX idx_videos_published_at ON videos(published_at);