                         NATIONAL KYC SYSTEM
                                  │
             ┌────────────────────┼────────────────────┐
             │                    │                    │
             ▼                    ▼                    ▼
       Citizen Portal      Institution Client    NKA/Oversight
             │                    │                    │
             └────────────────────┼────────────────────┘
                                  │
                                  ▼
                            API Gateway
                                  │
              ┌───────────────────┼───────────────────┐
              │                   │                   │
              ▼                   ▼                   ▼
        Authentication       Authorization       Protection
        OAuth/OIDC           RBAC + ABAC         Rate Limit
              │                   │                   │
              └───────────────────┼───────────────────┘
                                  │
                                  ▼
                         Application Services
                                  │
       ┌──────────────┬───────────┼──────────┬──────────────┐
       │              │           │          │              │
       ▼              ▼           ▼          ▼              ▼
    Identity         KYC       Consent   Verification    Institution
       │              │           │          │              │
       └──────────────┴───────────┼──────────┴──────────────┘
                                  │
                    ┌─────────────┴─────────────┐
                    ▼                           ▼
              PostgreSQL                 Fabric Gateway
                    │                           │
                    ▼                           ▼
             Protected State             Hyperledger Fabric
                    │
                    ▼
             Object Storage
                    │
                    ▼
             Encrypted Documents

                    + Audit / Monitoring / SIEM




                    national-kyc-system/
│
├── cmd/
│   ├── api-gateway/
│   ├── identity-service/
│   ├── kyc-service/
│   ├── institution-service/
│   ├── verification-service/
│   ├── consent-service/
│   ├── biometric-service/
│   ├── otp-service/
│   ├── document-service/
│   ├── notification-service/
│   ├── audit-service/
│   └── anomaly-service/
│
├── internal/
│   ├── auth/
│   ├── authorization/
│   ├── blockchain/
│   ├── cache/
│   ├── config/
│   ├── database/
│   ├── errors/
│   ├── http/
│   ├── logger/
│   ├── observability/
│   ├── security/
│   ├── storage/
│   └── validation/
│
├── services/
│   ├── identity/
│   ├── kyc/
│   ├── institution/
│   ├── verification/
│   ├── consent/
│   ├── biometric/
│   ├── otp/
│   ├── document/
│   ├── notification/
│   ├── audit/
│   └── anomaly/
│
├── applications/
│   ├── citizen-portal/
│   ├── institution-client/
│   ├── nka-console/
│   └── oversight-console/
│
├── blockchain/
│   ├── chaincode/
│   ├── channel/
│   ├── collections/
│   ├── network/
│   ├── organizations/
│   └── scripts/
│
├── infrastructure/
│   ├── backup/
│   ├── couchdb/
│   ├── docker/
│   ├── monitoring/
│   ├── nginx/
│   └── storage/
│
├── security/
│   ├── certificates/
│   ├── policies/
│   ├── secrets/
│   └── threat-model/
│
├── tests/
│   ├── e2e/
│   ├── integration/
│   ├── performance/
│   ├── security/
│   └── unit/
│
├── docs/
│
├── .env.example
├── .gitignore
├── .dockerignore
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── README.md
└── SECURITY.md