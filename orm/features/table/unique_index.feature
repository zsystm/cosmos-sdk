Feature: unique indexes

#  Scenario: get an entity that doesn't exists by primary key
#    Given a primary key for a non-existing entity
#    When I get the entity by primary key
#    Then expect a "not found" error
#    And expect grpc error code "NOT_FOUND"
#
#  Scenario: get an entity that doesn't exists by unique index
#    Given a unique key for a non-existing entity
#    When I get the entity by unique key
#    Then expect a "not found" error
#    And expect grpc error code "NOT_FOUND"
