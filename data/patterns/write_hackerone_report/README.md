# `write_hackerone_report` Pattern

## Description

The `write_hackerone_report` pattern is designed to assist a bug bounty hunter with writing a bug bounty report for the HackerOne platform. It knows the structure that is normally in place on HackerOne, and is instructed on how to extrapolate from requests, responses, and comments, what the report should be about and how to create steps to reproduce for that vulnerability. 

**This is version 0.1**. Please improve this prompt.

## Functionality

- Reviews the requests provided
- Reviews the responses provided
- Reviews the comments provided
- Generates a report which can be copy-pasted into HackerOne and adjusted for details.

### Use cases

1. This can be helpful for dynamic report generation for automation
2. This can be helpful when integrated with a Caido or Burp plugin to rapidly generate reports
3. This can be helpful when generating reports from the command-line

## Usage

This pattern is intended to be used with the `bbReportFormatter` tool which can be found here: https://github.com/rhynorater/bbReportFormatter

This utility automatically helps with the format that this pattern ingests which looks like this:

Request 1:
```
GET /...
```
Response 1:
```
HTTP/1.1 200 found...
```
Comment 1:
```
This request is vulnerable to blah blah blah
```

So, you'll add requests/responses to the report by using `cat req | bbReportFormatter`.
You'll add comments to the report using `echo "This request is vulnerable to blah blah blah" | bbReportFormatter`.

Then, when you run `bbReportFromatter --print-report` it will output the above, `write_hackerone_report` format.

So, in the end, this usage will be `bbReportFormatter --print-report | fabric -sp write_hackerone_report`.


## Meta

- **Author**: Justin Gardner (@Rhynorater)
- **Version Information**: 0.1
- **Published**: Jul 3, 2024

