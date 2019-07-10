### Why this program ?

Since there's a limit in the number of files we can merge with the [gsutil compose](https://cloud.google.com/storage/docs/json_api/v1/objects/compose) command, I made this program to simplify merging more than 32 files at once.

### Use case

Merge extracted files from Bigquery to Google Cloud Storage

### Usage 

```go run gcs-merge-files.go gs://<source-bucket>/filename_*.csv <number_of_files> gs://<destination-bucket>/<destination_filename>```

The wildcard must be a 12 ciffers number

### What does it do

1. Clones the first file to create destination file
2. Composes all files with the newly created destination file
