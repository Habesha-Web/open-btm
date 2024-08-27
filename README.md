# GraphQL Schema for Blue Test Management

This GraphQL schema defines types, queries, and mutations for managing requirements, tests, testsets, testtestsets, and issues in a test management system.

## Types

### Requirement
- `id`: ID of the requirement (Int!)
- `name`: Name of the requirement (String!)
- `description`: Description of the requirement (String!)
- `tests`: List of tests associated with the requirement ([Test!]!)

### Test
- `id`: ID of the test (Int!)
- `name`: Name of the test (String!)
- `steps`: Steps to execute the test (String!)
- `expected_result`: Expected result of the test (String!)
- `requirement_id`: ID of the associated requirement (String!)

### Testset
- `id`: ID of the testset (Int!)
- `name`: Name of the testset (String!)
- `description`: Description of the testset (String!)
- `tests`: List of tests included in the testset ([Test!]!)

### TestTestset
- `id`: ID of the testtestset (Int!)
- `test_id`: ID of the associated test (String!)
- `testset_id`: ID of the associated testset (String!)
- `run_status`: Status of the test run (RunStatus!)
- `run`: Run information (String!)
- `severity`: Severity of the issue (Severity!)
- `issues`: List of issues related to the testtestset ([Issue!]!)

### Issue
- `id`: ID of the issue (Int!)
- `issue_name`: Name of the issue (String!)
- `issue_status`: Status of the issue (String!)
- `issue_description`: Description of the issue (String!)

## Enums

### Severity
- `Critical`
- `High`
- `Medium`
- `Low`

### RunStatus
- `NA` (Not Available)
- `NR` (Not Required)
- `Blocked`
- `Passed`
- `Failed`
- `InProgress`

## Queries

- `requirements(page: Int!, size: Int!): [Requirement!]!` - Retrieve a paginated list of requirements.
- `requirement(id: Int!): Requirement!` - Retrieve a specific requirement by ID.
- `tests(page: Int!, size: Int!): [Test!]!` - Retrieve a paginated list of tests.
- `test(id: Int!): Test!` - Retrieve a specific test by ID.
- `testsets(page: Int!, size: Int!): [Testset!]!` - Retrieve a paginated list of testsets.
- `testset(id: Int!): Testset!` - Retrieve a specific testset by ID.
- `testsettests(test_id: Int!, testset_id: Int!, page: Int!, size: Int!): [Test!]!` - List names of tests in a specific testset.
- `testtestsets(page: Int!, size: Int!): [TestTestset!]!` - Retrieve a paginated list of testtestsets.
- `testtestset(id: Int!): TestTestset!` - Retrieve a specific testtestset by ID.
- `testtestsetissues(issue_id: Int!, test_testset_id: Int!, page: Int!, size: Int!): [Issue!]!` - List issues in a specific testtestset.
- `issues(page: Int!, size: Int!): [Issue!]!` - Retrieve a paginated list of issues.
- `issue(id: Int!): Issue!` - Retrieve a specific issue by ID.

## Mutations

- `createrequirement(input: CreateRequirementInput!): Requirement!` - Create a new requirement.
- `updaterequirement(input: UpdateRequirementInput!): Requirement!` - Update an existing requirement.
- `deleterequirement(id: Int!): Boolean!` - Delete a requirement by ID.
- `createtest(input: CreateTestInput!): Test!` - Create a new test.
- `updatetest(input: UpdateTestInput!): Test!` - Update an existing test.
- `deletetest(id: Int!): Boolean!` - Delete a test by ID.
- `createtestset(input: CreateTestsetInput!): Testset!` - Create a new testset.
- `updatetestset(input: UpdateTestsetInput!): Testset!` - Update an existing testset.
- `deletetestset(id: Int!): Boolean!` - Delete a testset by ID.
- `createtesttestset(test_id: Int!, testset_id: Int!): Test!` - Associate a test with a testset.
- `deletetesttestset(test_id: Int!, testset_id: Int!): Test!` - Remove a test from a testset.
- `updatetesttestset(input: UpdateTestTestsetInput!): TestTestset!` - Update an existing testtestset.
- `createissuetesttestset(issue_id: Int!, testtestset_id: Int!): Issue!` - Associate an issue with a testtestset.
- `deleteissuetesttestset(issue_id: Int!, testtestset_id: Int!): Issue!` - Remove an issue from a testtestset.
- `createissue(input: CreateIssueInput!): Issue!` - Create a new issue.
- `updateissue(input: UpdateIssueInput!): Issue!` - Update an existing issue.
- `deleteissue(id: Int!): Boolean!` - Delete an issue by ID.
