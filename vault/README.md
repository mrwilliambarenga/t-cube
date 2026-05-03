# vault
The API

# Decision Log

## (2026-05-03)

### Monetary Calculations

**Decision:** Use `shopspring/decimal` for all monetary values

**Rationale:** Floating-point types (`float64`) introduce rounding errors due to binary representation, which is unacceptable for financial calculations. `shopspring/decimal` provides precise base-10 arithmetic, ensuring deterministic and accurate results.

### Time Handling

**Decision:** Store all timestamps internally in UTC

**Rationale:** Trading activity may originate from multiple time zones, and tax calculations depend on consistent ordering and correct date boundaries. Using UTC avoids ambiguity and ensures deterministic behaviour across environments.

