# Impersonating a GSA so you can run this go-code locally

```sh
# https://cloud.google.com/docs/authentication/use-service-account-impersonation
gcloud auth application-default login --impersonate-service-account signed-url-agent@uploading-and-downloading-3988.iam.gserviceaccount.com
```

# Using the client library

https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest
https://pkg.go.dev/cloud.google.com/go/storage#hdr-Credential_requirements_for_signing

# Implementing one-time-use

You can achieve this by allowing the service account to read and create objects only. This way the link - once used - can't be used to upload again to the same file since it needs to delete it first which the service account doesn't have permission to.
