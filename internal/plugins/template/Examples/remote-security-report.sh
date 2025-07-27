#!/bin/bash
# remote-security-report.sh
# Usage: remote-security-report.sh cert host [report_name]

cert_path="$1"
host="$2"
report_name="${3:-report}"
temp_file="/tmp/security-report-${report_name}.txt"

# Copy the security report script to remote host
scp -i "$cert_path" /usr/local/bin/security-report.sh "${host}:~/security-report.sh" >&2

# Make it executable and run it on remote host
ssh -i "$cert_path" "$host" "chmod +x ~/security-report.sh && sudo ~/security-report.sh ${temp_file}" >&2

# Copy the report back
scp -i "$cert_path" "${host}:${temp_file}" "${temp_file}" >&2

# Cleanup remote files
ssh -i "$cert_path" "$host" "rm ~/security-report.sh ${temp_file}" >&2

# Output the local file path for fabric to read
echo "${temp_file}"

