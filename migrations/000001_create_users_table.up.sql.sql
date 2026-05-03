-- ===============================
-- USERS TABLE
-- ===============================
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       full_name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       job_role TEXT
);

-- ===============================
-- PROJECTS TABLE
-- ===============================
CREATE TABLE projects (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          status TEXT,
                          start_date DATE,
                          target_completion DATE NOT NULL
);

-- ===============================
-- TASKS TABLE
-- ===============================
CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       description TEXT,
                       task_type TEXT,
                       project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                       assigned_to INT REFERENCES users(id) ON DELETE SET NULL,
                       status TEXT,
                       priority TEXT,
                       percentage_done INT DEFAULT 0,
                       start_date DATE,
                       completion_date DATE,
                       due_date DATE NOT NULL
);

-- ===============================
-- TASK COMMENTS TABLE
-- ===============================
CREATE TABLE task_comments (
                               id SERIAL PRIMARY KEY,
                               user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                               task_id INT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                               created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               comment TEXT NOT NULL
);

-- ===============================
-- ATTACHMENTS TABLE
-- ===============================
CREATE TABLE attachments (
                             id SERIAL PRIMARY KEY,
                             file_name TEXT NOT NULL,
                             uploaded_by INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                             task_id INT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE
);

-- ===============================
-- TIME ENTRIES TABLE
-- ===============================
CREATE TABLE time_entries (
                              id SERIAL PRIMARY KEY,
                              user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                              task_id INT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                              date DATE NOT NULL,
                              start_time TIME NOT NULL,
                              end_time TIME NOT NULL,
                              activities TEXT NOT NULL
);

-- ===============================
-- NOTIFICATIONS TABLE
-- ===============================
CREATE TABLE notifications (
                               id SERIAL PRIMARY KEY,
                               user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                               message TEXT NOT NULL
);

-- ===============================
-- PROJECT MEMBERS TABLE
-- ===============================
CREATE TABLE project_members (
                                 id SERIAL PRIMARY KEY,
                                 user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                 project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE
);

-- ===============================
-- ROLES TABLE
-- ===============================
CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,
                       role TEXT NOT NULL
);

-- ===============================
-- MEMBER ROLES TABLE
-- ===============================
CREATE TABLE member_roles (
                              id SERIAL PRIMARY KEY,
                              project_member_id INT NOT NULL REFERENCES project_members(id) ON DELETE CASCADE,
                              role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE
);