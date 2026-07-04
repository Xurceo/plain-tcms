-- Users
INSERT INTO users (email, password_hash) VALUES
                                             ('admin@tcms.com', 'hash1'),
                                             ('user@tcms.com', 'hash2');

-- Organizations
INSERT INTO organizations (name) VALUES
    ('My Org');

-- Members
INSERT INTO organization_members (org_id, user_id, role) VALUES
                                                             ((SELECT id FROM organizations WHERE name = 'My Org'), (SELECT id FROM users WHERE email = 'admin@tcms.com'), 'owner'),
                                                             ((SELECT id FROM organizations WHERE name = 'My Org'), (SELECT id FROM users WHERE email = 'user@tcms.com'), 'member');

-- Projects
INSERT INTO projects (org_id, name, description, created_by) VALUES
                                                                 ((SELECT id FROM organizations WHERE name = 'My Org'), 'Web App', 'Main web application', (SELECT id FROM users WHERE email = 'admin@tcms.com')),
                                                                 ((SELECT id FROM organizations WHERE name = 'My Org'), 'Mobile App', 'iOS/Android app', (SELECT id FROM users WHERE email = 'admin@tcms.com'));

-- Test Suites
INSERT INTO test_suites (project_id, name) VALUES
                                               ((SELECT id FROM projects WHERE name = 'Web App'), 'Auth'),
                                               ((SELECT id FROM projects WHERE name = 'Web App'), 'Dashboard'),
                                               ((SELECT id FROM projects WHERE name = 'Mobile App'), 'Onboarding');

-- Test Cases
INSERT INTO test_cases (project_id, suite_id, title, status, priority, type) VALUES
                                                                                 ((SELECT id FROM projects WHERE name = 'Web App'), (SELECT id FROM test_suites WHERE name = 'Auth'), 'Login with valid credentials', 'active', 'critical', 'functional'),
                                                                                 ((SELECT id FROM projects WHERE name = 'Web App'), (SELECT id FROM test_suites WHERE name = 'Auth'), 'Login with invalid password', 'active', 'high', 'functional'),
                                                                                 ((SELECT id FROM projects WHERE name = 'Web App'), (SELECT id FROM test_suites WHERE name = 'Dashboard'), 'Dashboard loads in under 2s', 'active', 'medium', 'performance'),
                                                                                 ((SELECT id FROM projects WHERE name = 'Mobile App'), (SELECT id FROM test_suites WHERE name = 'Onboarding'), 'First launch shows welcome screen', 'draft', 'medium', 'functional');

-- Test Plans
INSERT INTO test_plans (project_id, name) VALUES
    ((SELECT id FROM projects WHERE name = 'Web App'), 'Sprint 1 Plan');

-- Test Runs
INSERT INTO test_runs (project_id, name, status, environment, build_version, created_by) VALUES
    ((SELECT id FROM projects WHERE name = 'Web App'), 'Sprint 1 Run', 'in_progress', 'staging', 'v1.0.0', (SELECT id FROM users WHERE email = 'admin@tcms.com'));

-- Test Results
INSERT INTO test_results (run_id, test_case_id, status, executed_by) VALUES
                                                                         ((SELECT id FROM test_runs WHERE name = 'Sprint 1 Run'), (SELECT id FROM test_cases WHERE title = 'Login with valid credentials'), 'passed', (SELECT id FROM users WHERE email = 'user@tcms.com')),
                                                                         ((SELECT id FROM test_runs WHERE name = 'Sprint 1 Run'), (SELECT id FROM test_cases WHERE title = 'Login with invalid password'), 'failed', (SELECT id FROM users WHERE email = 'user@tcms.com'));