# DateTime Plugin Tests

Simple test file for validating datetime plugin functionality.

## Basic Time Operations

```
Current Time: {{plugin:datetime:now}}
Time Only: {{plugin:datetime:time}}
Unix Timestamp: {{plugin:datetime:unix}}
Hour Start: {{plugin:datetime:startofhour}}
Hour End: {{plugin:datetime:endofhour}}
```

## Date Operations

```
Today: {{plugin:datetime:today}}
Full Date: {{plugin:datetime:full}}
Current Month: {{plugin:datetime:month}}
Current Year: {{plugin:datetime:year}}
```

## Period Operations

```
Week Start: {{plugin:datetime:startofweek}}
Week End: {{plugin:datetime:endofweek}}
Month Start: {{plugin:datetime:startofmonth}}
Month End: {{plugin:datetime:endofmonth}}
```

## Relative Time/Date

```
2 Hours Ahead: {{plugin:datetime:rel:2h}}
1 Day Ago: {{plugin:datetime:rel:-1d}}
Next Week: {{plugin:datetime:rel:1w}}
Last Month: {{plugin:datetime:rel:-1m}}
Next Year: {{plugin:datetime:rel:1y}}
```