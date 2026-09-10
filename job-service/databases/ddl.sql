CREATE TABLE jobs (
    id BIGSERIAL PRIMARY KEY,
    company_id BIGINT NOT NULL,
    judul VARCHAR(255) NOT NULL,
    about_role TEXT NOT NULL,
    responsibilities TEXT NOT NULL,
    deskripsi TEXT NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    gaji BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE job_required_skills (
    id BIGSERIAL,
    job_id BIGINT NOT NULL,
    name_license VARCHAR(255) NOT NULL,
    skill_tag VARCHAR(255) NOT NULL,
    required BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_jobs_company_id
ON jobs(company_id);

CREATE INDEX idx_jobs_status
ON jobs(status);

CREATE INDEX idx_job_required_skills_job_id
ON job_required_skills(job_id);

/* =========================================== */

select * from jobs;
select * from job_required_skills;

SELECT id, company_id, judul
FROM jobs;

/* reset data */
TRUNCATE TABLE jobs RESTART IDENTITY CASCADE;
TRUNCATE TABLE job_required_skills RESTART IDENTITY CASCADE;