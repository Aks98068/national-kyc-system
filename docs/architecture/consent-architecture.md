# National KYC System — Consent Management Architecture

## 1. Purpose

This document defines the consent management architecture of the National KYC System.

The consent architecture establishes how consent is created, captured, validated, used, withdrawn, renewed, audited, and enforced when protected KYC information is accessed or disclosed.

The architecture covers:

- Consent creation
- Consent capture
- Consent validation
- Consent scope
- Purpose limitation
- Institution-specific consent
- Consent expiry
- Consent withdrawal
- Consent renewal
- Consent versioning
- Consent enforcement
- Consent auditability
- Consent privacy
- Legal and emergency access
- Consent integration with KYC records
- Consent integration with Hyperledger Fabric
- Consent failure handling

---

## 2. Consent Principles

The National KYC System follows these principles:

1. Consent must be explicit where required by applicable law or policy.

2. Consent must be associated with an identifiable consenting party.

3. Consent must clearly identify the permitted purpose.

4. Consent must define the scope of information that may be accessed or disclosed.

5. Consent must identify the authorized institution or recipient where required.

6. Consent must have a defined lifecycle.

7. Consent must not grant access beyond the authorization policy.

8. Withdrawal of consent must prevent future access where legally applicable.

9. Historical consent events must remain auditable.

10. Consent records must be protected from unauthorized modification.

11. Consent must not be reused for an unrelated purpose.

12. Institutions must receive only the minimum information permitted by the approved consent and authorization policies.

13. Consent must be validated at the time of protected disclosure.

14. Consent failures must prevent unauthorized disclosure.

15. Consent information stored on blockchain must be minimized and must not unnecessarily expose personal information.

---

## 3. Consent Architecture

The logical architecture is:

```text
                 Citizen / Authorized Actor
                          |
                          v
                  Consent Interface
                          |
                          v
                  Consent Service
                          |
             +------------+------------+
             |                         |
             v                         v
      Consent Policy              Consent Store
        Evaluation                Protected Data
             |
             v
       Consent Record
             |
             v
      Authorization Layer
             |
             v
       KYC Data Request
             |
       +-----+------+
       |            |
       v            v
   Allowed        Denied
       |            |
       v            v
  KYC Data      Audit Event
   Access
       |
       v
   Audit Service
       |
       v
  Fabric Integrity