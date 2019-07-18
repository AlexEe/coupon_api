Feature: coupon api
  Coupon api must be able to list one, list all, update, add and delete coupons

  Scenario: List one
    Given there is a database
    And there is a webserver
    When I call list one endpoint
    Then webserver should call database
    And return one coupon

  Scenario: List all
    Given there is a database
    And there is a webserver
    When I call list one endpoint
    Then webserver should call database
    And return all coupons
    
