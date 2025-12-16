# TypesAggregationType

## Example Usage

```typescript
import { TypesAggregationType } from "@flexprice/sdk/models";

let value: TypesAggregationType = "SUM";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"COUNT" | "SUM" | "AVG" | "COUNT_UNIQUE" | "LATEST" | "SUM_WITH_MULTIPLIER" | "MAX" | "WEIGHTED_SUM" | Unrecognized<string>
```