# FR-02 — Institution Registration and Onboarding

## 1. Requirement Identification

| Field | Description |
|---|---|
| Requirement ID | FR-02 |
| Requirement Name | Institution Registration and Onboarding |
| Requirement Type | Functional |
| Priority | Critical |
| Primary Actor | National KYC Authority |
| Supporting Actors | Institution Administrator, Authorized Institutional Representative |
| Related Components | Institution Service, Identity Service, Authentication Service, Authorization Service, API Gateway, Fabric Network, Audit Service |
| Security Classification | Highly Sensitive |

---

## 2. Purpose

The National KYC System shall provide a controlled process through which eligible institutions can be registered, verified, approved, and onboarded into the national KYC ecosystem.

Institution onboarding shall establish a trusted institutional identity and define the organization's permitted access to national KYC services.

External institutions shall interact with the National KYC System through approved APIs and shall not receive unrestricted direct access to the Hyperledger Fabric network.

---

## 3. Scope

FR-02 covers:

- Institution registration
- Institutional identity creation
- Institutional representative registration
- Organization verification
- Institutional classification
- Institutional approval
- Certificate and credential provisioning
- Initial role assignment
- Permission assignment
- API access provisioning
- Institution lifecycle management
- Audit recording
- Suspension and revocation

---

## 4. Institution Types

The system may support approved institutional categories including:

- Banks
- Financial institutions
- Government agencies
- Insurance organizations
- Telecommunications organizations
- Other legally authorized institutions

The final categories shall be determined by national governance policy.

---

## 5. Preconditions

Before onboarding:

1. The institution must be legally eligible to participate.

2. The institution must submit required registration information.

3. The submitting representative must be authenticated.

4. The representative must be authorized to act on behalf of the institution.

5. Required institutional verification services must be available.

6. The National KYC Authority must have an authorized onboarding function.

---

## 6. Institution Registration Information

The system may collect:

- Legal institution name
- Institution registration number
- Institution type
- Regulatory authority
- Registered address
- Contact information
- Authorized representatives
- Compliance contacts
- Technical contacts
- Required certificates
- Approved service endpoints
- Intended KYC use cases
- Applicable regulatory classification

Sensitive information shall be handled according to the established data-classification policy.

---

## 7. High-Level Onboarding Flow

```text
Institution
     |
     v
Registration Request
     |
     v
Representative Authentication
     |
     v
Authorization Validation
     |
     v
Institution Verification
     |
     v
Regulatory / Eligibility Check
     |
     v
Duplicate Institution Detection
     |
     v
Administrative Approval
     |
     v
Institution Identity Creation
     |
     v
Role & Permission Assignment
     |
     v
Credential / Certificate Provisioning
     |
     v
API Access Activation
     |
     v
Audit Record
     |
     v
Institution ACTIVE
