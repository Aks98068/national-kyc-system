# Fabric Network Topology

## 1. Purpose

This document defines the Hyperledger Fabric network topology
for the National Centralized KYC System.

## 2. Organizations

The initial blockchain network contains:

1. National KYC Authority (NKA)
2. Independent Oversight Body

External financial institutions and other participating
institutions interact through the secured API Gateway and are
not direct Fabric network members.

## 3. Development Network

The initial development network will use:

- NKA organization
- Two NKA peers
- Independent Oversight organization
- One Oversight peer
- Development ordering service
- Fabric CA for organizational identities
- KYC channel
- CouchDB world state

The development topology is intentionally smaller than the
production target because it runs on a local development VM.

## 4. Production Target

The production architecture is intended to support:

- Multiple NKA peers
- Independent Oversight peer
- Five-node Raft ordering service
- Multi-site deployment
- Independent organizational administration
- Certificate rotation
- HSM-backed key protection
- Disaster recovery
- High availability

## 5. Organizations

### NKA Organization

Responsibilities:

- Maintain NKA peers
- Operate NKA administrative identities
- Perform authorized KYC verification
- Endorse authorized high-trust transactions
- Maintain NKA organizational identity infrastructure

### Independent Oversight Organization

Responsibilities:

- Maintain an independently administered peer
- Maintain independent organizational identities
- Participate in oversight and audit functions
- Maintain an independent ledger copy
- Prevent dependence on NKA infrastructure for oversight access

## 6. Ordering Service

The production target uses a five-node Raft ordering service.

The local development environment will use a reduced topology
to minimize resource requirements.

## 7. Channel

The system uses one shared KYC channel.

Conceptually:

    kycchannel
       |
       +-- NKA
       |
       +-- Oversight

## 8. Private Data Collections

Institution-specific information will use Private Data
Collections rather than separate channels for every institution.

Conceptually:

    kycchannel
       |
       +-- Institution A PDC
       +-- Institution B PDC
       +-- Institution C PDC

The exact collection membership and data classification rules
will be defined separately.

## 9. Identity

Fabric identities are managed through organizational
certificate authorities and MSP configuration.

Each organization maintains its own identity boundary.

## 10. Application Access

External institutions do not directly connect to Fabric peers.

The intended flow is:

    Institution
        |
        | HTTPS / mTLS
        v
    API Gateway
        |
        +-- Authentication
        +-- Authorization
        +-- Rate Limiting
        +-- Anomaly Detection
        +-- Consent Enforcement
        +-- Audit
        |
        v
    Fabric Gateway
        |
        v
    Hyperledger Fabric

## 11. Security Principles

- Organization separation
- Least privilege
- Strong identity
- Certificate-based authentication
- Multi-party endorsement
- Private data isolation
- Encryption in transit
- Encryption at rest
- Independent oversight
- Immutable auditability

## 12. Development vs Production

The local development network is not considered a
production deployment.

Production infrastructure will require additional controls
including HSMs, multi-site deployment, offline root CA,
operational certificate rotation, backup/recovery,
monitoring, and formal operational controls.