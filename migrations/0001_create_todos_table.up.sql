CREATE TABLE IF NOT EXISTS todos (
                                               id SERIAL PRIMARY KEY,
                                               description TEXT NOT NULL,
                                               status TEXT NOT NULL,
                                               created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                               updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
