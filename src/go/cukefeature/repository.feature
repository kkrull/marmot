Feature: Repository
  In order to find code in one repository and port it to related repositories
  As a developer
  I want to keep track of all repositories I have worked on

  @LocalDir
  Scenario: Meta Repos have no repositories when initialized
    Given I have initialized a new meta repo
    When I list repositories in that meta repo
    Then that repository listing should be empty

  @LocalDir
  Scenario: A Meta Repo remembers local repositories
    Given I have initialized a new meta repo
    And I have registered local repositories
    When I list repositories in that meta repo
    Then that repository listing should include local repositories that were registered
