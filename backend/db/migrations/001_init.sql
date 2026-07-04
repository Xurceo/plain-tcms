-- ============ USERS ============
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email TEXT NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============ AUTH/ORG LAYER ============
CREATE TABLE organizations (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               name TEXT NOT NULL,
                               created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE organization_members (
                                      org_id UUID REFERENCES organizations(id) ON DELETE CASCADE,
                                      user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                                      role TEXT NOT NULL CHECK (role IN ('owner','admin','member','viewer')),
                                      PRIMARY KEY (org_id, user_id)
);

-- ============ PROJECTS ============
CREATE TABLE projects (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          org_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
                          name TEXT NOT NULL,
                          description TEXT,
                          created_by UUID REFERENCES users(id),
                          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============ TEST CASE ORGANIZATION ============
CREATE TABLE test_suites (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                             parent_id UUID REFERENCES test_suites(id) ON DELETE CASCADE,
                             name TEXT NOT NULL,
                             description TEXT,
                             created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE test_cases (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                            suite_id UUID REFERENCES test_suites(id) ON DELETE SET NULL,
                            title TEXT NOT NULL,
                            description TEXT,
                            preconditions TEXT,
                            steps JSONB NOT NULL DEFAULT '[]',
                            expected TEXT,
                            status TEXT NOT NULL DEFAULT 'draft' CHECK (status IN ('draft','active','deprecated')),
                            priority TEXT NOT NULL DEFAULT 'medium' CHECK (priority IN ('low','medium','high','critical')),
                            type TEXT DEFAULT 'functional' CHECK (type IN ('functional','smoke','regression','performance','security')),
                            tags TEXT[] DEFAULT '{}',
                            created_by UUID REFERENCES users(id),
                            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE test_case_history (
                                   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                   test_case_id UUID NOT NULL REFERENCES test_cases(id) ON DELETE CASCADE,
                                   snapshot JSONB NOT NULL,
                                   changed_by UUID REFERENCES users(id),
                                   changed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============ TEST PLANS ============
CREATE TABLE test_plans (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                            name TEXT NOT NULL,
                            description TEXT,
                            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============ TEST RUNS ============
CREATE TABLE test_runs (
                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                           plan_id UUID REFERENCES test_plans(id) ON DELETE SET NULL,
                           name TEXT NOT NULL,
                           status TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending','in_progress','completed')),
                           environment TEXT,
                           build_version TEXT,
                           created_by UUID REFERENCES users(id),
                           created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                           completed_at TIMESTAMPTZ
);

CREATE TABLE test_run_cases (
                                run_id UUID REFERENCES test_runs(id) ON DELETE CASCADE,
                                test_case_id UUID REFERENCES test_cases(id) ON DELETE CASCADE,
                                PRIMARY KEY (run_id, test_case_id)
);

-- ============ RESULTS ============
CREATE TABLE test_results (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              run_id UUID NOT NULL REFERENCES test_runs(id) ON DELETE CASCADE,
                              test_case_id UUID NOT NULL REFERENCES test_cases(id) ON DELETE CASCADE,
                              status TEXT NOT NULL DEFAULT 'untested' CHECK (status IN ('untested','passed','failed','blocked','skipped')),
                              comment TEXT,
                              executed_by UUID REFERENCES users(id),
                              duration_ms INT,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE result_attachments (
                                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                    result_id UUID NOT NULL REFERENCES test_results(id) ON DELETE CASCADE,
                                    file_url TEXT NOT NULL,
                                    file_type TEXT,
                                    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============ DEFECTS ============
CREATE TABLE defects (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         result_id UUID REFERENCES test_results(id) ON DELETE SET NULL,
                         external_link TEXT,
                         title TEXT NOT NULL,
                         severity TEXT CHECK (severity IN ('low','medium','high','critical')),
                         status TEXT DEFAULT 'open' CHECK (status IN ('open','in_progress','resolved','closed')),
                         created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============ INDEXES ============
CREATE INDEX ON test_cases(project_id);
CREATE INDEX ON test_cases(suite_id);
CREATE INDEX ON test_suites(project_id);
CREATE INDEX ON test_suites(parent_id);
CREATE INDEX ON test_runs(project_id);
CREATE INDEX ON test_results(run_id);
CREATE INDEX ON test_results(test_case_id);
CREATE INDEX ON organization_members(user_id);