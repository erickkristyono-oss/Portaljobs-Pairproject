CREATE TABLE reports (
    id BIGSERIAL PRIMARY KEY,
    reporter_id BIGINT NOT NULL,
    target_type VARCHAR(20) NOT NULL,
    target_id BIGINT NOT NULL,
    reason TEXT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_reports_reporter_id
ON reports(reporter_id);

CREATE INDEX idx_reports_target
ON reports(target_type, target_id);

CREATE INDEX idx_reports_status
ON reports(status);


select * from reports;

TRUNCATE TABLE reports RESTART IDENTITY CASCADE;