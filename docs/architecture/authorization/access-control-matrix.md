# National KYC System — Access Control Matrix

## 1. Purpose

This document defines the authorization and access-control matrix for the National KYC System.

The access-control model establishes which identities may perform which operations against protected system resources.

The model combines:

- Role-Based Access Control (RBAC)
- Attribute-Based Access Control (ABAC)
- Purpose-based access control
- Consent enforcement
- Organization boundaries
- Least privilege
- Separation of duties
- Deny-by-default enforcement
- Audit requirements

Authentication establishes who is making a request.

Authorization determines whether that identity is permitted to perform the requested operation.

---

## 2. Authorization Principles

The National KYC System follows these principles:

1. Access is denied by default.

2. Authentication must occur before authorization.

3. Authorization must be evaluated for every protected operation.

4. Possession of valid credentials does not automatically grant permission.

5. Permissions must be associated with an approved role and applicable attributes.

6. Institutional users may only access information required for their authorized purpose.

7. Citizen information must be protected from unauthorized cross-institution access.

8. Consent must be enforced where required.

9. Sensitive operations require stronger authorization controls.

10. Administrative privileges must be separated from ordinary operational privileges.

11. Authorization decisions must be auditable.

12. Permissions must be revocable.

13. Role changes must trigger authorization reevaluation.

14. Organization membership must be considered when evaluating access.

15. Emergency access must be explicitly controlled and audited.

---

## 3. Authorization Actors

The primary authorization subjects are:

### 3.1 Citizen

A citizen may:

- View permitted personal KYC information
- Submit KYC information
- Submit supporting documents
- Provide consent
- Review consent history
- Request permitted corrections
- View relevant KYC status
- Withdraw consent where legally permitted

A citizen must not access another citizen's KYC information.

---

### 3.2 Institutional User

An institutional user operates on behalf of an authorized institution.

Examples include:

- Bank KYC officer
- Insurance officer
- Telecom verification officer
- Government service officer

Institutional access must be limited by:

- Institution
- Role
- Purpose
- Resource
- Operation
- Citizen consent where required
- Applicable policy

---

### 3.3 National KYC Authority User

NKA users may perform authorized national KYC operations according to their assigned responsibilities.

NKA access must not automatically provide unrestricted access to all system capabilities.

Administrative and operational responsibilities should remain separated.

---

### 3.4 Oversight User

Oversight users may access information required for:

- Audit
- Compliance
- Investigation
- Oversight
- Regulatory review

Oversight access must respect applicable privacy and data-minimization requirements.

---

### 3.5 System Administrator

Administrators manage system infrastructure and configuration.

Administrative privileges should not automatically provide unrestricted access to citizen KYC content.

Technical administration and data access should remain separated wherever practical.

---

### 3.6 Service Identity

Services may access only the APIs, resources, and operations explicitly required for their function.

A service identity must not inherit human-user permissions.

---

## 4. Authorization Model

The authorization decision can be represented as:

```text
                Request
                   |
                   v
          Authentication
                   |
                   v
          Identity Context
                   |
                   v
        +--------------------+
        | Authorization      |
        | Policy Engine      |
        +--------------------+
          |    |    |    |
          |    |    |    |
          v    v    v    v
        Role Org Purpose Consent
          |    |    |    |
          +----+----+----+
                   |
                   v
             Resource Check
                   |
                   v
             Operation Check
                   |
                   v
             Policy Decision
              /          \
             /            \
          ALLOW          DENY
             |
             v
       Protected Action
             |
             v
             Audit