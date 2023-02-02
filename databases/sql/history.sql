DROP TABLE IF EXISTS history CASCADE;
CREATE TABLE history (
    history_id BIGSERIAL PRIMARY KEY NOT NULL,
    reservation_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    CONSTRAINT fk_reservation FOREIGN KEY (reservation_id) REFERENCES reservation (reservation_id)
);