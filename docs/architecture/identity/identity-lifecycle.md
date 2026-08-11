# National KYC System — Identity Lifecycle Architecture

## 1. Purpose

This document defines the lifecycle of identities within the National KYC System.

The identity lifecycle establishes how identities are:

- Registered
- Provisioned
- Verified
- Activated
- Authenticated
- Authorized
- Updated
- Suspended
- Revoked
- Archived

The lifecycle applies to:

- Citizens
- National KYC Authority users
- Oversight users
- Authorized institutional users
- System administrators
- Service identities
- Fabric identities

Identity lifecycle management must maintain a controlled relationship between a real-world actor, their digital identity, organizational membership, authentication credentials, authorization attributes, and audit history.

---

## 2. Identity Lifecycle Principles

The system follows these principles:

1. Every system identity must have a unique identity identifier.

2. Identity registration must be subject to appropriate verification.

3. Authentication credentials must never be treated as proof of authorization.

4. Authentication and authorization must remain separate security functions.

5. Organizational membership must be explicitly associated with institutional identities.

6. An identity must not receive permissions before required verification and approval are completed.

7. Identity privileges must follow the principle of least privilege.

8. Privileges must be reviewed when an identity changes role, organization, or status.

9. Suspended identities must not perform protected operations.

10. Revoked identities must not regain access without a controlled reactivation process.

11. Identity lifecycle events must be auditable.

12. Identity credentials must be rotatable and revocable.

13. Fabric identities must follow the identity lifecycle of their owning organization.

14. Deactivation of an identity must not destroy historical audit records.

15. Identity lifecycle changes must not silently modify historical KYC transactions.

---

## 3. Identity Categories

The system recognizes several identity categories.

### 3.1 Citizen Identity

Represents an individual whose KYC information is maintained by the National KYC System.

A citizen identity may contain references to:

- National identity identifier
- KYC record
- Identity documents
- Verification status
- Consent records
- Biometric references
- Lifecycle status

---

### 3.2 Institutional Identity

Represents an authorized user belonging to a participating institution.

Examples include:

- Bank officers
- Insurance officers
- Telecom verification officers
- Government officers

Institutional identities must be associated with:

- Institution
- Organizational role
- Permissions
- Purpose restrictions
- Identity status

---

### 3.3 Administrative Identity

Represents authorized personnel responsible for operating and administering the platform.

Administrative access must be strongly restricted and continuously audited.

---

### 3.4 Oversight Identity

Represents authorized personnel belonging to the independent oversight organization.

Oversight identities must have an independently managed organizational identity boundary.

---

### 3.5 Service Identity

Represents a software service communicating with another protected service.

Service identities must not be treated as human identities.

They must have:

- Unique service identity
- Credential or certificate
- Defined permissions
- Service ownership
- Lifecycle status
- Audit association

---

### 3.6 Fabric Identity

Represents an organizational blockchain identity used by Fabric peers, clients, administrators, and other authorized components.

Fabric identities are governed through organizational MSP and certificate infrastructure.

---

## 4. Identity Lifecycle States

The primary lifecycle states are:

```text
                    +-------------+
                    |   REGISTER  |
                    +------+------+
                           |
                           v
                    +-------------+
                    |   PENDING   |
                    | VERIFICATION|
                    +------+------+
                           |
                    Verification
                       successful
                           |
                           v
                    +-------------+
                    |   ACTIVE    |
                    +------+------+
                           |
             +-------------+-------------+
             |                           |
             v                           v
      +-------------+             +-------------+
      |  SUSPENDED  |             |   UPDATE    |
      +------+------+             +------+------+
             |                           |
             | Reactivation              |
             +------------+--------------+
                          |
                          v
                    +-------------+
                    |   ACTIVE    |
                    +------+------+
                           |
                           v
                    +-------------+
                    |   REVOKED   |
                    +------+------+
                           |
                           v
                    +-------------+
                    |  ARCHIVED   |
                    +-------------+