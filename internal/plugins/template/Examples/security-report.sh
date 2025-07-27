#!/bin/bash

# security-report.sh - Enhanced system security information collection
# Usage: security-report.sh [output_file]

output_file=${1:-/tmp/security-report.txt}

{
    echo "=== System Security Report ==="
    echo "Generated: $(date)"
    echo "Hostname: $(hostname)"
    echo "Kernel: $(uname -r)"
    echo

    echo "=== System Updates ==="
    echo "Last update: $(stat -c %y /var/cache/apt/pkgcache.bin | cut -d' ' -f1)"
    echo "Pending updates:"
    apt list --upgradable 2>/dev/null
    
    echo -e "\n=== Security Updates ==="
    echo "Pending security updates:"
    apt list --upgradable 2>/dev/null | grep -i security

    echo -e "\n=== User Accounts ==="
    echo "Users with login shells:"
    grep -v '/nologin\|/false' /etc/passwd
    echo -e "\nUsers who can login:"
    awk -F: '$2!="*" && $2!="!" {print $1}' /etc/shadow
    echo -e "\nUsers with empty passwords:"
    awk -F: '$2=="" {print $1}' /etc/shadow
    echo -e "\nUsers with UID 0:"
    awk -F: '$3==0 {print $1}' /etc/passwd

    echo -e "\n=== Sudo Configuration ==="
    echo "Users/groups with sudo privileges:"
    grep -h '^[^#]' /etc/sudoers.d/* /etc/sudoers 2>/dev/null
    echo -e "\nUsers with passwordless sudo:"
    grep -h NOPASSWD /etc/sudoers.d/* /etc/sudoers 2>/dev/null

    echo -e "\n=== SSH Configuration ==="
    if [ -f /etc/ssh/sshd_config ]; then
        echo "Key SSH settings:"
        grep -E '^(PermitRootLogin|PasswordAuthentication|Port|Protocol|X11Forwarding|MaxAuthTries|PermitEmptyPasswords)' /etc/ssh/sshd_config
    fi
    
    echo -e "\n=== SSH Keys ==="
    echo "Authorized keys found:"
    find /home -name "authorized_keys" -ls 2>/dev/null

    echo -e "\n=== Firewall Status ==="
    echo "UFW Status:"
    ufw status verbose
    echo -e "\nIPTables Rules:"
    iptables -L -n

    echo -e "\n=== Network Services ==="
    echo "Listening services (port - process):"
    netstat -tlpn 2>/dev/null | grep LISTEN

    echo -e "\n=== Recent Authentication Failures ==="
    echo "Last 5 failed SSH attempts:"
    grep "Failed password" /var/log/auth.log | tail -5

    echo -e "\n=== File Permissions ==="
    echo "World-writable files in /etc:"
    find /etc -type f -perm -002 -ls 2>/dev/null
    echo -e "\nWorld-writable directories in /etc:"
    find /etc -type d -perm -002 -ls 2>/dev/null

    echo -e "\n=== System Resource Usage ==="
    echo "Disk Usage:"
    df -h
    echo -e "\nMemory Usage:"
    free -h
    echo -e "\nTop 5 CPU-using processes:"
    ps aux --sort=-%cpu | head -6

    echo -e "\n=== System Timers ==="
    echo "Active timers (potential scheduled tasks):"
    systemctl list-timers --all

    echo -e "\n=== Important Service Status ==="
    for service in ssh ufw apparmor fail2ban clamav-freshclam; do
        echo "Status of $service:"
        systemctl status $service --no-pager 2>/dev/null
    done

    echo -e "\n=== Fail2Ban Logs ==="
    echo "Recent Fail2Ban activity (fail2ban.log):"
    if [ -f /var/log/fail2ban.log ]; then
        echo "=== Current log (fail2ban.log) ==="
        cat /var/log/fail2ban.log
    else
        echo "fail2ban.log not found"
    fi

    if [ -f /var/log/fail2ban.log.1 ]; then
        echo -e "\n=== Previous log (fail2ban.log.1) ==="
        cat /var/log/fail2ban.log.1
    else
        echo -e "\nfail2ban.log.1 not found"
    fi

    echo -e "\n=== Fail2Ban Status ==="
    echo "Currently banned IPs:"
    sudo fail2ban-client status


} > "$output_file"

# Output the file path for fabric to read
echo "$output_file"

