# National KYC System — Data Flow Architecture

## 1. Purpose

This document defines how information flows through the National KYC System from the point at which a citizen or authorized institution initiates a KYC operation until the operation is validated, recorded, audited, and completed.

The data-flow architecture establishes:

- KYC registration flow
- KYC verification flow
- Identity verification flow
- Document verification flow
- Biometric verification flow
- Consent validation flow
- Authorization flow
- KYC record creation flow
- KYC record update flow
- KYC retrieval flow
- Institution-to-institution KYC sharing flow
- Blockchain transaction flow
- Off-chain data flow
- Audit flow
- Integrity verification flow
- Error and rejection flow
- Data ownership boundaries
- Privacy boundaries
- Security boundaries

The data-flow architecture separates sensitive personal information from blockchain transaction metadata while maintaining cryptographic integrity and auditable state transitions.

---

## 2. Data Flow Principles

The National KYC System follows these principles:

1. Sensitive personal information must not be stored directly on the blockchain unless explicitly approved by the national data governance policy.

2. Large documents, biometric templates, photographs, and other sensitive payloads are stored in protected off-chain storage.

3. The blockchain stores the minimum information required to establish trusted state, integrity, provenance, authorization, and auditability.

4. Every important KYC state transition must be authenticated and authorized.

5. Every access to protected KYC information must be auditable.

6. Consent must be validated before information is disclosed to an institution where consent is legally required.

7. Institutions must only receive the minimum information required for the approved purpose.

8. Hashes and integrity metadata must be used to detect unauthorized modification of off-chain information.

9. Communication between system components must use authenticated and encrypted channels.

10. Failed operations must not create partially valid KYC states.

---

## 3. High-Level Data Flow

The high-level flow is:

```text
Citizen / Authorized Institution
            |
            v
      Client Application
            |
            v
       API Gateway
            |
            v
    Authentication Layer
            |
            v
     Authorization Layer
            |
            v
       KYC Service
       /    |     \
      /     |      \
     v      v       v
Identity  Document  Biometric
Service   Service   Service
     \      |       /
      \     |      /
       v    v     v
      Verification Engine
            |
            v
       Consent Service
            |
            v
      KYC Record Service
        /           \
       v             v
Off-chain Storage   Fabric Network
       |                |
       |                v
       |           Endorsement
       |                |
       |                v
       |             Ordering
       |                |
       |                v
       |              Ledger
       |
       v
Encrypted Sensitive Data

            |
            v
       Audit Service
            |
            v
     Audit / Monitoring