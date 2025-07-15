CREATE TABLE IF NOT EXISTS ping_results (
    id SERIAL PRIMARY KEY,
    host TEXT,
    packet_loss REAL,
    avg_latency REAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS http_results (
    id SERIAL PRIMARY KEY,
    url TEXT,
    status_code INT,
    duration_ms INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
