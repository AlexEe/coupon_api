Feature: list all
    In order to list all coupons
    As an API user
    I need to be able to access database and make GET requests

    Scenario: should get coupons
      Given there are coupons
      When I send "GET" request to "/coupons"
      Then the response code should be 200
      And all coupons should be listed