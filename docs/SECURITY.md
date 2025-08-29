# Security Policy

## Supported Versions

We aim to provide security updates for the latest version of Fabric.

We recommend always using the latest version of Fabric for security fixes and improvements.

## Reporting Security Vulnerabilities

**Please DO NOT report security vulnerabilities through public GitHub issues.**

### Preferred Reporting Method

Send security reports directly to: **<kayvan@sylvan.com>** and CC to the project maintainer at **<daniel@danielmiessler.com>**

### What to Include

Please provide the following information:

1. **Vulnerability Type**: What kind of security issue (e.g., injection, authentication bypass, etc.)
2. **Affected Components**: Which parts of Fabric are affected
3. **Impact Assessment**: What could an attacker accomplish
4. **Reproduction Steps**: Clear steps to reproduce the vulnerability
5. **Proposed Fix**: If you have suggestions for remediation
6. **Disclosure Timeline**: Your preferred timeline for public disclosure

### Example Report Format

```text
Subject: [SECURITY] Brief description of vulnerability

Vulnerability Type: SQL Injection
Affected Component: Pattern database queries
Impact: Potential data exposure
Severity: High

Reproduction Steps:
1. Navigate to...
2. Submit payload: ...
3. Observe...

Evidence:
[Screenshots, logs, or proof of concept]

Suggested Fix:
Use parameterized queries instead of string concatenation...
```

## Security Considerations

### API Keys and Secrets

- Never commit API keys to the repository
- Store secrets in environment variables or secure configuration
- Use the built-in setup process for key management
- Regularly rotate API keys

### Input Validation

- All user inputs are validated before processing
- Special attention to pattern definitions and user content
- URL validation for web scraping features

### AI Provider Integration

- Secure communication with AI providers (HTTPS/TLS)
- Token handling follows provider best practices
- No sensitive data logged or cached unencrypted

### Network Security

- Web server endpoints properly authenticated when required
- CORS policies appropriately configured
- Rate limiting implemented where necessary

## Vulnerability Response Process

1. **Report Received**: We'll acknowledge receipt within 24 hours
2. **Initial Assessment**: We'll evaluate severity and impact within 72 hours
3. **Investigation**: We'll investigate and develop fixes
4. **Fix Development**: We'll create and test patches
5. **Coordinated Disclosure**: We'll work with reporter on disclosure timeline
6. **Release**: We'll release patched version with security advisory

### Timeline Expectations

- **Critical**: 1-7 days
- **High**: 7-30 days
- **Medium**: 30-90 days
- **Low**: Next scheduled release

## Bug Bounty

We don't currently offer a formal bug bounty program, but we deeply appreciate security research and will:

- Acknowledge contributors in release notes
- Provide credit in security advisories
- Consider swag or small rewards for significant findings

## Security Best Practices for Users

### Installation

- Download Fabric only from official sources
- Verify checksums when available
- Keep installations up to date

### Configuration

- Use strong, unique API keys
- Don't share configuration files containing secrets
- Set appropriate file permissions on config directories

### Usage

- Be cautious with patterns that process sensitive data
- Review AI provider terms for data handling
- Consider using local models for sensitive content

## Known Security Limitations

### AI Provider Dependencies

Fabric relies on external AI providers. Security depends partly on:

- Provider security practices
- Data transmission security
- Provider data handling policies

### Pattern Execution

Custom patterns could potentially:

- Process sensitive inputs inappropriately
- Generate outputs containing sensitive information
- Be used for adversarial prompt injection

**Recommendation**: Review patterns carefully, especially those from untrusted sources.

## Security Updates

Security updates are distributed through:

- GitHub Releases with security tags
- Security advisories on GitHub
- Project documentation updates

Subscribe to the repository to receive notifications about security updates.

## Contact

For non-security issues, please use GitHub issues.
For security concerns, email: **<kayvan@sylvan.com>** and CC to **<daniel@danielmiessler.com>**

---

*We take security seriously and appreciate the security research community's help in keeping Fabric secure.*
