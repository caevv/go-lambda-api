Feature: CRUD

  Scenario: Create user
    Given I have a new client "John"
    When I ask to create a new user "John"
    Then the user "John" should have been created

  Scenario: Find user
    Given I have a client "John"
    When I search for "John"
    Then the user "John" should have been returned
