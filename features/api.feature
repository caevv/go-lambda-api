Feature: CRUD

  Scenario: Create user
    Given I have a new client "John"
    When I ask to create a new user "John"
    Then the user "John" should have been created
