# National KYC System — Organization Architecture

## 1. Purpose

This document defines the participating organizations, their identities,
responsibilities, and trust boundaries within the National KYC blockchain
network.

## 2. Network Participants

### 2.1 National KYC Authority

Organization ID:

NKA

Responsibilities:

- National KYC governance
- KYC policy management
- Participant onboarding
- Identity governance
- Regulatory oversight
- KYC verification coordination
- Audit oversight
- Certificate and organizational authorization management

Blockchain role:

- Network administrator
- Governance authority
- KYC policy authority

---

### 2.2 Financial Institutions

Organization ID:

FINANCIAL_ORG

Responsibilities:

- Customer KYC submission
- Customer identity verification
- KYC updates
- KYC status queries
- Compliance operations
- Suspicious or anomalous KYC reporting

Blockchain role:

- Authorized KYC participant
- Transaction submitter
- KYC verification participant

---

### 2.3 Government / Regulatory Organizations

Organization ID:

GOVERNMENT_ORG

Responsibilities:

- Regulatory oversight
- Authorized KYC verification
- Compliance monitoring
- Audit activities
- Investigation support

Blockchain role:

- Authorized verifier
- Oversight participant
- Auditor

---

### 2.4 Orderer Organization

Organization ID:

ORDERER_ORG

Responsibilities:

- Transaction ordering
- Block creation
- Consensus participation
- Network availability

Blockchain role:

- Ordering service infrastructure

---

## 3. Trust Boundaries

The network contains separate trust domains:

1. National KYC Authority
2. Financial Institutions
3. Government / Regulatory Organizations
4. Ordering Infrastructure

Each organization must have independently managed:

- Membership Service Provider (MSP)
- Certificate Authority
- Organizational identities
- Peer identities
- Administrative identities
- TLS identities
- Access-control policies

## 4. Identity Principles

The system must:

- Never trust an identity solely because it knows an identifier.
- Authenticate every participant.
- Authorize operations according to organizational role.
- Use cryptographic identities for blockchain transactions.
- Separate administrative identities from application identities.
- Protect private keys.
- Support certificate revocation.
- Maintain auditable identity lifecycle events.

## 5. Initial Development Network

Development network:

- Org1
- Org2
- OrdererOrg
- CouchDB
- Fabric CA

The development network is not considered production-ready.

## 6. Production Architecture

The production architecture must use separate organizational
identities and infrastructure for each participating institution.

Production deployment must not reuse development certificates,
private keys, identities, or genesis/channel artifacts.