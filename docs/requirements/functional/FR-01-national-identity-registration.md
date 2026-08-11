# FR-01 — National Identity Registration

## 1. Requirement Identification

| Field | Description |
|---|---|
| Requirement ID | FR-01 |
| Requirement Name | National Identity Registration |
| Requirement Type | Functional |
| Priority | Critical |
| Primary Actor | Citizen |
| Supporting Actors | National KYC Authority, Authorized Registration Officer |
| Related Components | Identity Service, KYC Service, Authentication Service, Authorization Service, Consent Service, Audit Service, Off-Chain Storage, Hyperledger Fabric |
| Security Classification | Highly Sensitive |

---

## 2. Purpose

The National KYC System shall provide a controlled mechanism for registering a citizen within the national KYC ecosystem.

The registration process shall establish a trusted digital identity reference that can subsequently be used for KYC creation, verification, authorized information retrieval, institutional onboarding, consent management, and other approved national KYC operations.

The system shall ensure that a citizen cannot be registered multiple times using the same authoritative identity information unless the operation is explicitly handled as a correction, recovery, or identity-resolution process.

---

## 3. Scope

FR-01 covers:

- Citizen registration initiation
- Identity information collection
- Identity-document submission
- Identity validation
- Duplicate identity detection
- Identity uniqueness validation
- Citizen identity creation
- Initial KYC record association
- Secure storage of sensitive information
- Creation of the initial trusted identity state
- Audit recording
- Registration status management
- Registration rejection and failure handling

FR-01 does not define the detailed biometric verification algorithm, document verification algorithm, or institutional KYC sharing process. Those capabilities are defined by separate functional requirements.

---

## 4. Actors

### 4.1 Citizen

The citizen provides the information and evidence required for registration.

### 4.2 National KYC Authority

The National KYC Authority manages the national identity and KYC infrastructure and performs authorized administrative operations.

### 4.3 Authorized Registration Officer

An authorized officer may assist citizens with registration where an assisted-registration process is permitted.

### 4.4 Identity Service

The Identity Service creates and manages the citizen's system identity reference.

### 4.5 Verification Service

The Verification Service validates submitted identity information against approved verification sources.

### 4.6 Audit Service

The Audit Service records security-sensitive and state-changing registration events.

---

## 5. Preconditions

Before registration begins:

1. The registration service must be operational.

2. The requesting actor must have an authenticated session where authentication is required.

3. The registration channel must use an approved secure communication mechanism.

4. The registration actor must have permission to initiate the registration operation.

5. Required authoritative identity information must be available.

6. The system must have access to the required verification services.

7. Required protected storage services must be available.

8. The system must be able to create an auditable registration transaction.

---

## 6. Registration Inputs

The registration process may require:

- National identity number or approved identity reference
- Full legal name
- Date of birth
- Address information
- Contact information
- Identity-document information
- Identity-document evidence
- Required biometric evidence
- Citizenship information where applicable
- Registration location
- Registration channel
- Consent information where required

The exact mandatory fields shall be determined by national identity policy and data-governance rules.

---

## 7. Registration Workflow

The standard registration workflow shall operate as follows:

```text
Citizen / Registration Officer
            |
            v
     Registration Request
            |
            v
      Authentication
            |
            v
       Authorization
            |
            v
    Input Validation
            |
            v
 Identity / Document Validation
            |
            v
   Duplicate Detection
            |
            v
   Identity Verification
            |
            v
   Consent Validation
            |
            v
     Identity Creation
            |
            +--------------------+
            |                    |
            v                    v
     Protected Storage      Fabric Transaction
            |                    |
            +---------+----------+
                      |
                      v
                Audit Record
                      |
                      v
           Registration Completed
