DROP TABLE IF EXISTS history;
CREATE TABLE history (
    history_id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    vehicle_id INT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    quantity INT NOT NULL,
    payment_code VARCHAR(255) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    prepayment INT NOT NULL,
    is_booked BOOLEAN DEFAULT FALSE,
    return_status VARCHAR(255) NOT NULL,
    rating DECIMAL(10, 2) NOT NULL,
    transaction_date TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (user_id),
    CONSTRAINT fk_vehicle FOREIGN KEY (vehicle_id) REFERENCES vehicle (vehicle_id)
);