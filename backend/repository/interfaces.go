/*
 * // TCMS - Test Case Management System
 * // Copyright (C) 2026 Pavlo Shnal
 * //
 * // This program is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Affero General Public License as published
 * // by the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // This program is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Affero General Public License for more details.
 * //
 * // You should have received a copy of the GNU Affero General Public License
 * // along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

package repository

import "github.com/xurceo/plain-tcms/entities"

type OrganizationRepository interface {
	GetAllOrganizations() ([]entities.Organization, error)
	GetOrganizationByID(id string) (entities.Organization, error)
	CreateOrganization(req entities.CreateOrganizationRequest) (entities.Organization, error)
	DeleteOrganization(id string) error
	GetMembers(orgID string) ([]entities.OrganizationMember, error)
	AddMember(orgID string, req entities.CreateOrganizationMemberRequest) (entities.OrganizationMember, error)
	RemoveMember(orgID string, userID string) error
}

type ProjectRepository interface {
	GetAllProjects() ([]entities.Project, error)
	GetProjectByID(id string) (entities.Project, error)
	GetProjectsByOrgID(orgID string) ([]entities.Project, error)
	CreateProject(req entities.CreateProjectRequest) (entities.Project, error)
	UpdateProject(id string, req entities.CreateProjectRequest) (entities.Project, error)
	DeleteProject(id string) error
}

type TestCaseRepository interface {
	GetTestCases(projectID string) ([]entities.TestCase, error)
	GetTestCaseByID(id string) (entities.TestCase, error)
	CreateTestCase(projectID string, req entities.CreateTestCaseRequest) (entities.TestCase, error)
	UpdateTestCase(id string, req entities.CreateTestCaseRequest) (entities.TestCase, error)
	DeleteTestCase(id string) error
	GetTestCaseHistory(testCaseID string) ([]entities.TestCaseHistory, error)
}

type TestSuiteRepository interface {
	GetTestSuitesByProject(projectID string) ([]entities.TestSuite, error)
	GetTestSuiteByID(id string) (entities.TestSuite, error)
	CreateTestSuite(projectID string, req entities.CreateTestSuiteRequest) (entities.TestSuite, error)
	UpdateTestSuite(id string, req entities.CreateTestSuiteRequest) (entities.TestSuite, error)
	DeleteTestSuite(id string) error
}

type TestPlanRepository interface {
	GetTestPlansByProject(projectID string) ([]entities.TestPlan, error)
	GetTestPlanByID(id string) (entities.TestPlan, error)
	CreateTestPlan(projectID string, req entities.CreateTestPlanRequest) (entities.TestPlan, error)
	UpdateTestPlan(id string, req entities.CreateTestPlanRequest) (entities.TestPlan, error)
	DeleteTestPlan(id string) error
}

type TestRunRepository interface {
	GetTestRunsByProject(projectID string) ([]entities.TestRun, error)
	GetTestRunByID(id string) (entities.TestRun, error)
	CreateTestRun(projectID string, req entities.CreateTestRunRequest) (entities.TestRun, error)
	UpdateTestRun(id string, req entities.CreateTestRunRequest) (entities.TestRun, error)
	DeleteTestRun(id string) error
	GetCasesByRun(runID string) ([]entities.TestCase, error)
	AddCaseToRun(runID string, testCaseID string) error
	RemoveCaseFromRun(runID string, testCaseID string) error
	GetResultsByRun(runID string) ([]entities.TestResult, error)
	AddResultToRun(runID string, req entities.CreateTestResultRequest) (entities.TestResult, error)
}

type TestResultRepository interface {
	GetTestResultByID(id string) (entities.TestResult, error)
	UpdateTestResult(id string, req entities.CreateTestResultRequest) (entities.TestResult, error)
	GetAttachmentsByResult(resultID string) ([]entities.ResultAttachment, error)
	AddAttachmentToResult(resultID string, req entities.CreateResultAttachmentRequest) (entities.ResultAttachment, error)
}

type DefectRepository interface {
	GetDefectByID(id string) (entities.Defect, error)
	UpdateDefect(id string, req entities.CreateDefectRequest) (entities.Defect, error)
	DeleteDefect(id string) error
}

type UserRepository interface {
	Register(email, password string) (entities.User, error)
	Login(email, password string) (entities.User, error)
}
