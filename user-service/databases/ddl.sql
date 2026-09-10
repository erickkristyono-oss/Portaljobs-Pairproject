-- =========================================
-- USERS
-- =========================================

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users
ADD CONSTRAINT users_role_check
CHECK (role IN ('jobseeker', 'company', 'admin'));

ALTER TABLE users
ADD CONSTRAINT users_status_check
CHECK (status IN ('active', 'suspended'));


-- =========================================
-- PROFILES
-- =========================================

CREATE TABLE profiles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(30) NOT NULL,
    address TEXT NOT NULL,
    faculty VARCHAR(100),
    major VARCHAR(100),
    education_level VARCHAR(100),
    started INT,
    graduated INT,
    cv_url TEXT,
	portfolio_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE profiles
ADD CONSTRAINT profiles_user_id_fk
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;


-- =========================================
-- SKILLS
-- =========================================

CREATE TABLE skills (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name_license VARCHAR(150) NOT NULL,
    level VARCHAR(20) NOT NULL,
  	skill_tag text,
    organization VARCHAR(150),
    grade VARCHAR(50),
    expired_date VARCHAR(50),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE skills
ADD CONSTRAINT skills_user_id_fk
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE skills
ADD CONSTRAINT skills_level_check
CHECK (level IN ('beginner', 'intermediate', 'expert'));


-- =========================================
-- PORTFOLIOS
-- =========================================

CREATE TABLE portfolios (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name_project VARCHAR(150) NOT NULL,
    organization VARCHAR(150) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE portfolios
ADD CONSTRAINT portfolios_user_id_fk
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;


-- =========================================
-- INDEX
-- =========================================

CREATE INDEX idx_skills_user_id
ON skills(user_id);

CREATE INDEX idx_portfolios_user_id
ON portfolios(user_id);


select * from users u ;
select * from profiles p ;
select * from skills s ;
select * from portfolios p ;

TRUNCATE TABLE users RESTART IDENTITY CASCADE;
TRUNCATE TABLE profiles RESTART IDENTITY CASCADE;
TRUNCATE TABLE skills RESTART IDENTITY CASCADE;
TRUNCATE TABLE portfolios RESTART IDENTITY CASCADE;


