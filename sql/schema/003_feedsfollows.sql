-- +goose Up
CREATE TABLE feed_follows (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid not null,
    feed_id uuid not null,
    constraint fk_userid foreign key (user_id)
    references users(id) on delete cascade,
    constraint fk_feedid foreign key (feed_id)
    references feeds(id) on delete cascade,
    constraint uq_userid_feedid unique (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;

