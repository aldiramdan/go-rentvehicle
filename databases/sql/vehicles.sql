DROP TABLE IF EXISTS vehicle CASCADE;
CREATE TABLE vehicle (
    vehicle_id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    category_id INT NOT NULL,
    picture VARCHAR(255) NOT NULL,
    rating DECIMAL(10, 2) NOT NULL,
    total_rent INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category (category_id)
);