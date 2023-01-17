-- +migrate Up
-- +migrate StaementBegin

CREATE TABLE category (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE type (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(256),    
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE movie (
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(256),
    description VARCHAR(256),
	image_url VARCHAR(256),
	aired VARCHAR(256),
    total_episode INT,
    status VARCHAR(100),
    category_id BIGINT,
    type_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE account (
    id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(256),
    password VARCHAR(256),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(256),
    image_url VARCHAR(256),
    account_id BIGINT,
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE watchlist (
    id SERIAL NOT NULL PRIMARY KEY,
    users_id BIGINT,
    movie_id BIGINT
);

ALTER TABLE movie
ADD CONSTRAINT fk_movie_category FOREIGN KEY (category_id) REFERENCES category (id);

ALTER TABLE movie
ADD CONSTRAINT fk_movie_type FOREIGN KEY (type_id) REFERENCES type (id);

ALTER TABLE users
ADD CONSTRAINT fk_users_account FOREIGN KEY (account_id) REFERENCES account (id);

ALTER TABLE watchlist
ADD CONSTRAINT fk_watchlist_users FOREIGN KEY (users_id) REFERENCES users (id);

ALTER TABLE watchlist
ADD CONSTRAINT fk_watchlist_movie FOREIGN KEY (movie_id) REFERENCES movie (id);

-- +migrate StatementEnd