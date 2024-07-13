Feature: Repository
  In order to find code in one repository and port it to related repositories
  As a developer
  I want to keep track of all repositories I have worked on

  @LocalDir
  Scenario: Meta Repos have no repositories when initialized
    Given I have initialized a new meta repo
    When I list repositories in that meta repo
    Then that repository listing should be empty
