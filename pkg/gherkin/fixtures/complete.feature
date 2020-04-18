# Some feature comment
@featureTag1 @featureTag2
Feature: feature title
  The description

  # Some background comment
  Background: background title
  The description.

    Given given value
    And and value
    When when value
    Then then value
    But but value

  @ScenarioTag1 @Scenario_Tag2
  Scenario: test tags

    # step with tags
    @stepTag
    Given tags

  Scenario: multiline arguments.

    Given json:
    """
    {
      "key": "value"
    }
    """
    And empty json:
    """
    """
    And table:
      | header |
      | value  |
    And empty table row:
      |  |
