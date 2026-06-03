CREATE TABLE
  IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    post_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    content text NOT NULL,
    created_at timestamp(0)
    WITH
      TIME ZONE NOT NULL DEFAULT NOW ()
  );
