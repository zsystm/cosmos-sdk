Feature: saving entities

  Scenario: can't insert duplicate primary key
    Given an existing entity
    When I insert an entity with the same primary key
    Then expect a "already exists" error
    And expect grpc error code "ALREADY_EXISTS"

  Scenario: can't update entity that doesn't exist
    When I update an entity that doesn't exist
    Then expect a "not found" error
    And expect grpc error code "NOT_FOUND"

  Scenario: can't violate unique constraint on insert
    Given an existing entity
    When I insert an entity with the same unique key
    Then expect a "constraint violation" error
    And expect grpc error code "FAILED_PRECONDITION"

  Scenario: can't violate unique constraint on update
    Given an existing entity
    When I update another entity to have the same unique key
    Then expect a "constraint violation" error
    And expect grpc error code "FAILED_PRECONDITION"
