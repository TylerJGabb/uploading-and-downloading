# Running Locally

## Impersonation (doesn't work IDK why)

```sh
# https://cloud.google.com/docs/authentication/use-service-account-impersonation
gcloud auth application-default login --impersonate-service-account signed-url-agent@uploading-and-downloading-3988.iam.gserviceaccount.com
```

## Environment Variables

I've added a util function in [utils.go](utils/utils.go) to build the storage client with credentials file
if running locally. To configure your local env to work with this you'll need to set the

```sh
USE_CREDS_FILE=true
```

Additionally, you'll need to generate a key for the GSA and save it as $REPO_ROOT/creds.json

# Using the client library

https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest
https://pkg.go.dev/cloud.google.com/go/storage#hdr-Credential_requirements_for_signing

# Implementing one-time-upload use case

You can achieve this by allowing the service account to read and create objects only. This way the link - once used - can't be used to upload again to the same file since it needs to delete it first which the service account doesn't have permission to.
