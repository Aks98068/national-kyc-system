# National KYC System — System Architecture

## 1. Purpose

This document defines the high-level system architecture of the National KYC System.

The architecture establishes the major technical components, trust boundaries, communication paths, security boundaries, data boundaries, blockchain integration, off-chain storage, identity management, authorization, verification services, and operational infrastructure required to implement the National KYC platform.

The system architecture integrates:

- National KYC identity services
- Institution identity and membership
- Authentication
- Authorization
- KYC processing
- Identity verification
- Document verification
- Biometric verification
- Consent management
- KYC record management
- Off-chain protected storage
- Hyperledger Fabric
- Audit services
- Monitoring
- Security services
- Administrative services
- External institutional integrations

---

## 2. Architectural Principles

The National KYC System follows these principles:

1. Sensitive citizen information must be protected throughout its lifecycle.

2. Blockchain must not be used as unrestricted storage for sensitive personal information.

3. Hyperledger Fabric provides trusted state, provenance, integrity, and auditable transaction history.

4. Sensitive documents and other large protected data remain in controlled off-chain storage.

5. Every system actor must have an authenticated identity.

6. Authorization must be evaluated before protected operations are performed.

7. Access must follow least-privilege principles.

8. Institutional access must be limited according to role, attributes, purpose, and applicable policy.

9. Consent must be enforced wherever legally or institutionally required.

10. All security-sensitive and state-changing operations must be auditable.

11. Services must communicate through authenticated and encrypted channels.

12. System components must be independently scalable where appropriate.

13. Failure of one subsystem must not result in an inconsistent KYC state.

14. Previous KYC versions must remain traceable through controlled lifecycle transitions.

15. Blockchain and off-chain data must be cryptographically connected where integrity verification is required.

---

## 3. High-Level Architecture

The logical architecture is:

```text
+-------------------------------------------------------------+
|                    External Actors                          |
|                                                             |
| Citizens | Banks | Government | Telecom | Insurance |      |
| Other Authorized Institutions | Administrators              |
+-----------------------------+-------------------------------+
                              |
                              v
+-------------------------------------------------------------+
|                    Client Applications                      |
|                                                             |
| Citizen Portal | Institution Portal | Admin Portal |       |
| Verification Applications                                   |
+-----------------------------+-------------------------------+
                              |
                              v
+-------------------------------------------------------------+
|                       API Gateway                           |
|                                                             |
| TLS Termination | Routing | Rate Limiting | Request Policy |
+-----------------------------+-------------------------------+
                              |
                              v
+-------------------------------------------------------------+
|                 Authentication & Identity                   |
|                                                             |
| Identity Provider | OAuth/OIDC | Institution Identity      |
| Certificate Identity | Session Management                   |
+-----------------------------+-------------------------------+
                              |
                              v
+-------------------------------------------------------------+
|              Authorization & Policy Layer                   |
|                                                             |
| RBAC | ABAC | Least Privilege | Purpose | Consent          |
+-----------------------------+-------------------------------+
                              |
                              v
+-------------------------------------------------------------+
|                    Application Services                     |
|                                                             |
| KYC Service                                                  |
| Identity Service                                             |
| Document Service                                             |
| Biometric Service                                            |
| Verification Service                                         |
| Consent Service                                              |
| KYC Record Service                                           |
| Institution Service                                          |
| Audit Service                                                |
| Notification Service                                         |
+----------------------+----------------------+---------------+
                       |                      |
                       v                      v
+--------------------------------+   +------------------------+
|       Protected Off-Chain      |   |    Hyperledger Fabric  |
|             Data               |   |                        |
|                                |   | Peers                  |
| Encrypted KYC Data             |   | Orderer                |
| Documents                      |   | Channels               |
| Protected Biometrics           |   | Chaincode              |
| Verification Evidence          |   | Ledger                 |
| Supporting Files               |   | MSP / Certificates     |
+--------------------------------+   +------------------------+
                       |                      |
                       +----------+-----------+
                                  |
                                  v
+-------------------------------------------------------------+
|                 Audit & Security Operations                 |
|                                                             |
| Audit Storage | Monitoring | Alerting | SIEM | Investigation|
+-------------------------------------------------------------+