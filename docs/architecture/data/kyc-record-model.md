# National KYC System — KYC Record Data Model

## 1. Purpose

This document defines the canonical data model for KYC records within the National KYC System.

The model establishes:

- KYC record identity
- Citizen identity references
- KYC lifecycle
- Verification information
- Identity-document references
- Biometric references
- Consent references
- Institution references
- Audit references
- Integrity information
- Versioning
- Data ownership
- Privacy boundaries
- Blockchain storage boundaries
- Off-chain storage boundaries

The model is designed for Hyperledger Fabric and CouchDB while ensuring that highly sensitive information is not unnecessarily stored directly on the blockchain.

---

# 2. Core Design Principles

The KYC record model follows these principles.

## 2.1 Data Minimization

Only information required for the operation of the National KYC System shall be maintained.

Unnecessary personal information shall not be placed on the blockchain.

## 2.2 Separation of On-Chain and Off-Chain Data

The blockchain shall contain:

- identifiers
- references
- verification state
- hashes
- consent state
- audit references
- timestamps
- transaction references
- integrity metadata

Highly sensitive material shall remain in protected off-chain storage.

Examples include:

- identity-document images
- scanned certificates
- biometric templates
- photographs
- large supporting documents

## 2.3 Immutable Auditability

Important KYC state changes shall generate blockchain transactions.

The ledger shall provide an immutable history of state transitions.

## 2.4 Versioned Records

KYC records shall be versioned.

A new verification or approved update shall not silently overwrite historical state.

## 2.5 Least Privilege

Applications and institutions shall only receive the minimum information required for an authorized operation.

## 2.6 Cryptographic Integrity

Sensitive off-chain objects shall have cryptographic hashes recorded in the appropriate ledger structure.

This allows later integrity verification without storing the original sensitive object on-chain.

---

# 3. KYC Record Overview

The canonical KYC record consists of the following logical components.

```text
KYC Record
│
├── Record Identity
│
├── Citizen Reference
│
├── Profile
│
├── Identity Documents
│
├── Biometric References
│
├── Verification
│
├── Consent
│
├── Institution References
│
├── Audit References
│
├── Integrity Metadata
│
└── Version Information