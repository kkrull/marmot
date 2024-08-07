Feature: Repository
  In order to find code in one repository and port it to related repositories
  As a developer
  I want to keep track of all repositories I have worked on

  @LocalDir
  Scenario: A Meta Repo has no repositories when initialized
    Given I have initialized a new meta repo
    When I list local repositories in that meta repo
    Then that repository listing should be empty
    When I list remote repositories in that meta repo
    Then that repository listing should be empty

  @LocalDir
  Scenario: A Meta Repo remembers local repositories
    Given Git repositories on the local filesystem
    And I have initialized a new meta repo
    And I have registered those local repositories with a meta repo
    When I list local repositories in that meta repo
    Then that repository listing should include those local repositories

  @LocalDir
  Scenario: A Meta Repo remembers remote repositories
    Given I have initialized a new meta repo
    And I have registered remote repositories
    When I list remote repositories in that meta repo
    Then that repository listing should include those remote repositories
