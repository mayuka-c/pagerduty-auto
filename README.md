# PagerDuty Auto Acknowledge

## Pre-requisites
### Get email_id, pagerduty_user_id and api_token from PagerDuty
- email_id - Go to Pagerduty -> UserProfile -> get the email address
- pagerduty_user_id - Go to PagerDuty -> UserProfile -> Get the userID which is available in URL (<pagerduty.com>/users/<user_id>)
- api_token - Check online on how to generate apiToken in PagerDuty

## Usage
### Clone the repository
```bash
    git clone https://github.com/mayuka-c/pagerduty-auto.git
    cd pagerduty-auto/runner
```

### Run the program
```bash
    go run main.go -email <email_id> -userID <pagerduty_user_id> -apiToken <api_token>
```
