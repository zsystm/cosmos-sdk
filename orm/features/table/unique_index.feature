Feature: unique indexes

  Scenario: get an entity that doesn't exists by primary key
    Given a key "foo"
    When I get an entity by the "name" key
    Then expect a "not found" error
    And expect grpc error code "NOT_FOUND"
#
  Scenario: get an entity that doesn't exists by unique index
    Given a key "bar"
    When I get an entity by the "unique" key
    Then expect a "not found" error
    And expect grpc error code "NOT_FOUND"
