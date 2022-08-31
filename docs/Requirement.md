# Scenario
_This part helps imagine the business in a more general way. Not only concern the task's mission, but also concern the business scenario behind it so that the output could be more product-oriented and could be used in the released product._

**Brief description:** Our company needs an app to automatically send marketing email to a list of customers using a predefined template. Because each customer has their own info, the mail should adapt that info into the template. My duty is implementing the core execution of the application except the email sending part.
# User requirement
_Users here could possibly be a programmer who implements the API to automatically send emails to customers._
## Functional requirements
- The app receives predefined arguments and outputs the prepared emails to a specific file.
- The inputs are defined clearly:
    - path_to_email_template: JSON file
    - path_to_customer: CSV file
    - path_to_output_emails: directory path
    - path_to_errors: CSV file
## Non-functional requirements:
- Must use unit tests
- Flexible enough to integrate a module that send real emails
- Docker build and run are preferred