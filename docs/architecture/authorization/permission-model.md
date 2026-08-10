# National KYC System — Identity and Permission Model

## 1. Purpose

This document defines the identity, authentication, authorization,
organization, role, and permission model for the National KYC System.

The objective is to ensure that every operation against KYC information
is performed by an authenticated and authorized identity.

---

## 2. Security Principles

The system follows these principles:

- Least privilege
- Need-to-know access
- Default deny
- Separation of duties
- Strong authentication
- Organization isolation
- Role-based authorization
- Attribute-based authorization
- Cryptographic identity
- Complete auditability
- Explicit authorization for sensitive operations

---

## 3. Organizational Identity

Every participating organization must have a unique Fabric identity.

Examples:

- NKA
- Financial Institution
- Government / Regulatory Organization
- Orderer Organization

An organization must not use another organization's credentials.

---

## 4. User Identity

Each human or application identity must be uniquely identifiable.

A user identity must contain sufficient attributes to determine:

- Organization
- Role
- Identity status
- Authorization scope
- Identity type

Shared user accounts are prohibited.

---

## 5. Identity Types

The system will distinguish between:

### 5.1 Human Users

Examples:

- NKA administrator
- Bank KYC officer
- Government auditor

### 5.2 Service Identities

Examples:

- API service
- KYC service
- Audit service
- Notification service
- Blockchain gateway

Service identities must not be treated as ordinary human accounts.

---

## 6. Roles

Initial roles:

### 6.1 NKA Administrator

Responsibilities:

- Manage participating organizations
- Manage KYC policies
- Manage authorized access
- Manage network governance
- Review audit information

### 6.2 Institution Administrator

Responsibilities:

- Manage users belonging to the institution
- Assign permitted institutional roles
- Manage institution-level configuration

### 6.3 KYC Officer

Responsibilities:

- Submit KYC information
- Verify KYC information
- Request KYC updates
- View authorized customer information

### 6.4 Government / Regulatory Officer

Responsibilities:

- Perform authorized verification
- Review compliance information
- Access permitted audit information
- Support authorized investigations

### 6.5 Auditor

Responsibilities:

- Review audit events
- Review authorized transactions
- Verify historical activity

Auditors must not automatically receive permission to modify KYC records.

### 6.6 System Administrator

Responsibilities:

- Operate infrastructure
- Maintain services
- Monitor system health

Infrastructure administration must not automatically grant access to
customer KYC information.

---

## 7. Role Separation

The following responsibilities must remain separated where practical:

- Infrastructure administration
- KYC administration
- User administration
- Financial institution operations
- Audit operations
- Governance operations

A system administrator must not automatically become a KYC administrator.

---

## 8. Permission Categories

Permissions will be divided into categories.

### Identity Management

- Create identity
- Disable identity
- Revoke identity
- Reactivate identity
- View identity status

### KYC Management

- Create KYC record
- Read KYC record
- Update KYC record
- Submit KYC verification
- Approve KYC verification
- Reject KYC verification
- Suspend KYC status
- Revoke KYC status

### Audit

- View audit events
- Search audit events
- Export authorized audit information

### Organization Management

- Register organization
- Approve organization
- Suspend organization
- Revoke organization

### Policy Management

- Create policy
- Update policy
- Activate policy
- Retire policy

---

## 9. Default Deny

Every request must be denied unless an explicit authorization rule
allows the requested operation.

Authorization must never depend only on:

- Username
- User ID
- Organization name
- Client-side role information

---

## 10. RBAC

Role-Based Access Control will determine the base permissions assigned
to an identity.

Example:

KYC Officer

may have:

- KYC_CREATE
- KYC_READ
- KYC_UPDATE
- KYC_SUBMIT_VERIFICATION

but must not automatically have:

- ORGANIZATION_APPROVE
- POLICY_UPDATE
- IDENTITY_REVOKE

---

## 11. ABAC

Attribute-Based Access Control will provide additional authorization
constraints.

Potential attributes include:

- Organization
- Role
- Identity status
- KYC record classification
- Request purpose
- Transaction type
- Regulatory authority
- Time constraints
- Geographic jurisdiction
- Approval status

A request must satisfy both the applicable role permissions and
attribute-based restrictions.

---

## 12. Organization Isolation

A financial institution must not automatically access all information
belonging to another institution.

Cross-organization access requires an explicit authorization rule.

---

## 13. Sensitive Operations

The following operations require stronger authorization controls:

- KYC approval
- KYC rejection
- KYC revocation
- Identity revocation
- Organization approval
- Policy modification
- Access to highly sensitive information
- Bulk data access
- Administrative changes

These operations may require:

- Additional approval
- Multi-party authorization
- Step-up authentication
- Enhanced audit logging

---

## 14. Audit Requirements

Every security-sensitive operation must generate an audit event.

The audit event should record information such as:

- Actor identity
- Organization
- Role
- Operation
- Resource
- Timestamp
- Result
- Request identifier
- Authorization decision

Sensitive personal information must not unnecessarily appear in audit logs.

---

## 15. Service-to-Service Authorization

Internal services must authenticate each other.

A service must not trust another service merely because it is located
inside the internal network.

Service communication must use authenticated and encrypted channels.

---

## 16. Fabric Identity

Blockchain transactions must use cryptographic identities managed through
Hyperledger Fabric MSP infrastructure.

Fabric identity and application identity must be mapped carefully.

An application user must never receive unrestricted access to Fabric
administrative credentials.

---

## 17. Private Keys

Private keys must:

- Never be committed to Git
- Never be stored in source code
- Never be placed in public repositories
- Have restricted filesystem permissions
- Be rotated according to policy
- Be protected using appropriate key-management mechanisms

Production environments should evaluate Hardware Security Module (HSM)
support for high-value cryptographic keys.

---

## 18. Certificate Lifecycle

The identity lifecycle must support:

1. Registration
2. Enrollment
3. Activation
4. Renewal
5. Suspension
6. Revocation
7. Replacement
8. Expiration

Revoked identities must not be permitted to perform authorized
operations.

---

## 19. Authorization Decision

A simplified authorization decision is:

Authenticated
    +
Valid identity
    +
Active organization
    +
Required role
    +
Required permission
    +
Required attributes
    +
Valid resource scope
    +
Valid operation
    =
ALLOW

Otherwise:

DENY

---

## 20. Security Objective

The authorization architecture must ensure that compromise of a single
application account does not automatically provide unrestricted access
to the national KYC system.

Compromise of one organization must also not automatically compromise
the identities or authorization boundaries of other organizations.