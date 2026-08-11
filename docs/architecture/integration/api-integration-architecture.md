# National KYC System — API & Integration Architecture

## 1. Purpose

This document defines the API and integration architecture of the National KYC System.

The architecture establishes how citizens, authorized institutions, government systems, internal application services, verification services, and the Hyperledger Fabric network communicate with the National KYC platform.

The API and integration architecture establishes:

- External institution integration
- API Gateway responsibilities
- Authentication and identity validation
- Authorization and policy enforcement
- Institution onboarding
- KYC registration APIs
- KYC verification APIs
- KYC retrieval APIs
- Consent APIs
- Document verification APIs
- Biometric verification APIs
- Fabric Gateway integration
- Internal service-to-service communication
- Error handling
- Rate limiting
- Audit logging
- API security
- API versioning
- Request and response boundaries
- Integration monitoring

The architecture ensures that external institutions do not directly access internal services or Hyperledger Fabric peers.

All external integration must pass through controlled API boundaries where identity, authorization, consent, security, validation, auditing, and policy enforcement can be applied.

---

## 2. Integration Principles

The National KYC System follows these API and integration principles:

1. External institutions must not directly connect to Hyperledger Fabric peers.

2. All external API communication must use encrypted transport.

3. Every API request must be associated with an authenticated identity.

4. Authorization must be evaluated before protected operations are performed.

5. APIs must enforce least-privilege access.

6. Sensitive information must only be returned when the requesting institution has an approved purpose and sufficient authorization.

7. Consent must be validated before disclosure where consent is required.

8. API requests must be validated before reaching application services.

9. Rate limiting must protect the platform against abuse and excessive requests.

10. Security-sensitive operations must generate auditable records.

11. API responses must expose only the information required for the approved operation.

12. Internal services must communicate through authenticated service boundaries.

13. Fabric transactions must be submitted through the controlled Fabric Gateway integration.

14. API failures must not create inconsistent KYC records.

15. APIs must support controlled versioning to prevent breaking existing institutional integrations.

16. Integration activity must be monitored for security, reliability, and operational anomalies.

---

## 3. High-Level Integration Architecture

The high-level integration flow is:

```text
+-------------------------------------------------------------+
|                 External Institutions                       |
|                                                             |
| Banks | Government | Telecom | Insurance | Other Approved  |
| Institutions                                              |
+-----------------------------+-------------------------------+
                              |
                              | HTTPS / mTLS
                              v
+-------------------------------------------------------------+
|                       API Gateway                           |
|                                                             |
| TLS Termination                                             |
| Request Validation                                          |
| Authentication                                             |
| Authorization                                               |
| Rate Limiting                                               |
| Request Routing                                             |
| Security Policy                                             |
| API Versioning                                              |
| Audit Event Generation                                      |
+-----------------------------+
                              |
                              v
+-------------------------------------------------------------+
|                  Application Service Layer                  |
|                                                             |
| KYC Service                                                 |
| Identity Service                                            |
| Document Service                                            |
| Biometric Service                                           |
| Verification Service                                        |
| Consent Service                                             |
| KYC Record Service                                          |
| Institution Service                                         |
+----------------------+----------------------+---------------+
                       |                      |
                       v                      v
+--------------------------------+   +------------------------+
|      Protected Off-Chain       |   |    Fabric Gateway      |
|           Storage              |   |                        |
|                                |   | Transaction Submission |
| Documents                      |   | Query                   |
| Biometrics                     |   | Endorsement             |
| Verification Evidence          |   | Policy Enforcement      |
| Sensitive KYC Data             |   |                         |
+--------------------------------+   +-----------+------------+
                                                |
                                                v
                                      +----------------------+
                                      | Hyperledger Fabric    |
                                      |                      |
                                      | Peers                |
                                      | Orderers             |
                                      | Channels             |
                                      | Chaincode            |
                                      | Ledger               |
                                      +----------------------



                                ## 8. Institution Onboarding & API Credential Lifecycle

Before an institution can access the National KYC System, it must complete an approved institutional onboarding process.

The onboarding process establishes the institution's identity, authority, technical integration requirements, and permitted access scope.

The lifecycle is:

```text
Institution Application
        |
        v
Institution Verification
        |
        v
Administrative Approval
        |
        v
Institution Registration
        |
        v
Identity / Certificate Provisioning
        |
        v
Role & Permission Assignment
        |
        v
API Access Configuration
        |
        v
Integration Testing
        |
        v
Production Activation

## 16. Fabric Gateway Integration

Application services must not communicate directly with individual Fabric peers for normal application operations.

The preferred integration path is through the Fabric Gateway.

```text
Application Service
        |
        v
  Fabric Gateway
        |
        v
 Endorsement Process
        |
        v
 Ordering Service
        |
        v
    Ledger