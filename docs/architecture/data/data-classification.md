# National KYC System — Data Classification and Storage Architecture

## 1. Purpose

This document defines the classification, storage location, protection
requirements, and lifecycle considerations for information processed by
the National KYC System.

The system must avoid unnecessarily storing sensitive personal information
on the shared blockchain ledger.

---

## 2. Data Architecture Principles

The system follows these principles:

- Data minimization
- Purpose limitation
- Least-privilege access
- Separation of sensitive data from ledger data
- Encryption in transit
- Encryption at rest
- Controlled cross-organization sharing
- Integrity verification
- Complete auditability
- Explicit retention policies
- Secure deletion where legally and technically applicable

---

## 3. Data Classification

The system will initially classify information into:

### PUBLIC

Information that may safely be exposed publicly.

Examples:

- Public system information
- Public organization information
- Public service metadata

---

### INTERNAL

Information intended only for authorized system participants.

Examples:

- Non-sensitive operational metadata
- Internal configuration information
- Non-sensitive organization metadata

---

### CONFIDENTIAL

Information that requires controlled organizational access.

Examples:

- KYC workflow information
- Verification metadata
- Institution-specific operational information
- Internal compliance information

---

### HIGHLY SENSITIVE

Information requiring the strongest protection.

Examples:

- Personally identifiable information
- Identity document information
- Identity document images
- Biometric information
- Sensitive investigation information
- Authentication secrets
- Cryptographic private keys

---

## 4. Blockchain Data

The blockchain should primarily contain information required for:

- Transaction integrity
- Verification
- State management
- Authorization decisions
- Organizational accountability
- Auditability

Potential blockchain records include:

- KYC record identifier
- Record state
- Verification state
- Relevant timestamps
- Organization identifier
- Transaction identifier
- Integrity hashes
- Approval events
- Revocation events
- Status transitions
- Authorized audit metadata

Raw sensitive personal information should not be placed on the shared
ledger unless there is a documented legal, operational, and architectural
requirement.

---

## 5. Off-Chain KYC Data

Sensitive KYC information should be stored in protected off-chain
storage.

Potential information includes:

- Personal information
- Contact information
- Address information
- Identity document information
- Document images
- Supporting evidence

The off-chain system must provide:

- Encryption at rest
- Encryption in transit
- Access control
- Audit logging
- Backup
- Recovery
- Retention management
- Secure deletion mechanisms
- Access monitoring

---

## 6. Document Storage

Identity documents and supporting files should be stored outside the
blockchain ledger.

Document storage must provide:

- Encryption
- Access control
- Integrity protection
- Malware scanning
- File-type validation
- Size restrictions
- Access auditing
- Retention controls

The blockchain may contain a cryptographic integrity reference to the
document rather than the document itself.

---

## 7. Biometric Data

Biometric information requires a dedicated security architecture.

Biometric information must not be placed directly into the ordinary
shared ledger.

The biometric architecture must consider:

- Template protection
- Encryption
- Strict access control
- Purpose limitation
- Consent/legal requirements where applicable
- Retention
- Revocation
- Auditability
- Replay protection
- Presentation-attack detection where applicable

The detailed biometric architecture will be defined separately.

---

## 8. Private Data Collections

Hyperledger Fabric Private Data Collections may be used where information
must be shared only among explicitly authorized organizations.

Potential examples:

- Institution-specific KYC information
- Restricted verification information
- Sensitive inter-organization records

Private Data Collections must not be treated as a replacement for all
off-chain storage.

Their use must be justified according to:

- Data sensitivity
- Participants requiring access
- Retention requirements
- Performance
- Legal requirements
- Operational requirements

---

## 9. Cryptographic Integrity

For selected off-chain objects, the system may maintain a cryptographic
hash on the blockchain.

Example:

Off-chain document
        |
        v
Cryptographic hash
        |
        v
Blockchain record

When the document is later retrieved:

1. Retrieve authorized document.
2. Calculate its cryptographic hash.
3. Compare it with the blockchain integrity reference.
4. Detect unexpected modification.

---

## 10. Data Encryption

Sensitive data must be encrypted:

### In Transit

Use authenticated encrypted communication such as TLS.

### At Rest

Sensitive databases, object storage, backups, and other protected storage
must use encryption at rest.

### Application-Level Encryption

Particularly sensitive fields may require application-level encryption
in addition to storage-level encryption.

Encryption keys must be managed separately from encrypted data.

---

## 11. Key Management

Cryptographic keys must not be stored directly inside source code.

The production architecture should use a dedicated key-management
solution.

Potential technologies may include:

- HSM
- Enterprise key-management system
- Cloud KMS where appropriate
- Fabric-compatible HSM integration

The final key-management architecture will be documented separately.

---

## 12. Access Control

Access to sensitive information must require:

- Authenticated identity
- Active account
- Valid organization
- Appropriate role
- Required permission
- Resource authorization
- Purpose/attribute checks where applicable

Access must follow the principle of least privilege.

---

## 13. Audit

Access to sensitive KYC information must generate an audit event.

The audit event should contain appropriate metadata such as:

- Actor
- Organization
- Operation
- Resource identifier
- Timestamp
- Result
- Request identifier
- Authorization decision

Sensitive personal information should not unnecessarily appear in logs.

---

## 14. Data Lifecycle

KYC information must have a defined lifecycle:

1. Collection
2. Validation
3. Verification
4. Storage
5. Controlled access
6. Update
7. Review
8. Suspension
9. Revocation
10. Retention
11. Archival where required
12. Secure disposal where applicable

---

## 15. Data Minimization

The system must collect and retain only information required for an
approved purpose.

Duplicate storage of sensitive information should be avoided.

---

## 16. Backup

Backups must protect:

- Application databases
- Blockchain state where required
- Private data
- Object storage
- Configuration
- Audit records
- Cryptographic material according to key-management policy

Backups must be:

- Encrypted
- Access controlled
- Tested for restoration
- Protected against unauthorized modification

---

## 17. Disaster Recovery

The system must define recovery procedures for:

- Database failure
- Peer failure
- Orderer failure
- Certificate authority failure
- Storage failure
- Application failure
- Network failure
- Security incident
- Data corruption

Recovery objectives will be defined during infrastructure architecture.

---

## 18. Data Residency

The production deployment must define where sensitive KYC information
is physically and logically stored.

Data residency requirements must be considered before selecting
infrastructure providers.

---

## 19. Separation of Duties

No single infrastructure or application role should automatically have
unrestricted access to all KYC information.

Database administration, application administration, blockchain
administration, security administration, and audit responsibilities should
be separated where practical.

---

## 20. Initial Storage Model

The initial architecture is:

                    KYC DATA
                       |
        ┌──────────────┼──────────────┐
        |              |              |
        v              v              v
   Fabric Ledger   Private Data   Off-chain Storage
        |           Collections        |
        |              |               |
   Integrity       Restricted       Sensitive
   /state          sharing          information
        |
        v
      Audit

This architecture will be refined during detailed system design.