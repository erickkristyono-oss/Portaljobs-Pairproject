CREATE TABLE applications ( 
id BIGSERIAL PRIMARY KEY, 
user_id BIGINT NOT NULL, 
job_id BIGINT NOT NULL, 
status VARCHAR(20) NOT NULL, 
applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ); 


CREATE INDEX idx_applications_user_id ON applications(user_id); 
CREATE INDEX idx_applications_job_id ON applications(job_id); 
CREATE INDEX idx_applications_status ON applications(status); 
CREATE UNIQUE INDEX idx_applications_user_job ON applications(user_id, job_id);

select * from applications a ;

SELECT id, user_id, job_id, status
FROM applications
ORDER BY id;


TRUNCATE TABLE applications RESTART IDENTITY CASCADE;