"""JSON → TOON converter.

Spec-compliant implementation per TOON Specification v3.0.
See: https://github.com/toon-format/spec/blob/main/SPEC.md

Key spec sections implemented:
- §7.1: Escape sequences (\\ \" \n \r \t)
- §7.2: Quoting rules for string values
- §9.1: Primitive arrays (inline CSV)
- §9.3: Tabular arrays (uniform objects)
- §12: Indentation (2 spaces per level)
"""

from __future__ import annotations

import re
from typing import Any

__all__ = ["json_to_toon"]


# Characters requiring quoting per §7.2
_SPECIAL_CHARS = frozenset(':"\\\n\t\r[]{}')

# Regex for numeric-like strings per §7.2
_NUMERIC_PATTERN = re.compile(r"^-?\d+(?:\.\d+)?(?:e[+-]?\d+)?$", re.IGNORECASE)
_LEADING_ZERO_PATTERN = re.compile(r"^0\d+$")


def _needs_quoting(value: str, delimiter: str = ",") -> bool:
    """Check if string needs quoting per TOON Spec §7.2."""
    if not value:
        return True
    if value in ("true", "false", "null"):
        return True
    if value[0] == " " or value[-1] == " ":
        return True
    if value[0] == "-":
        return True
    if any(c in _SPECIAL_CHARS for c in value):
        return True
    if delimiter in value:
        return True
    if _NUMERIC_PATTERN.match(value) or _LEADING_ZERO_PATTERN.match(value):
        return True
    return False


def _escape(value: str) -> str:
    """Escape string per TOON Spec §7.1."""
    return (value
        .replace("\\", "\\\\")
        .replace('"', '\\"')
        .replace("\n", "\\n")
        .replace("\r", "\\r")
        .replace("\t", "\\t"))


def _quote(value: str, delimiter: str = ",") -> str:
    """Quote and escape string if needed."""
    return f'"{_escape(value)}"' if _needs_quoting(value, delimiter) else value


def _encode_value(value: Any) -> str:
    """Encode a primitive value to TOON."""
    if value is None:
        return "null"
    if isinstance(value, bool):
        return "true" if value else "false"
    if isinstance(value, (int, float)):
        return str(value)
    if isinstance(value, str):
        return _quote(value)
    return _quote(str(value))


def _is_primitive(value: Any) -> bool:
    """Check if value is a TOON primitive."""
    return isinstance(value, (str, int, float, bool, type(None)))


def _is_tabular(arr: list) -> bool:
    """Check if array qualifies for tabular format per §9.3."""
    if not arr or not isinstance(arr[0], dict):
        return False
    keys = frozenset(arr[0].keys())
    return all(
        isinstance(item, dict)
        and frozenset(item.keys()) == keys
        and all(_is_primitive(v) for v in item.values())
        for item in arr
    )


def _encode_list(arr: list, indent: int = 0) -> str:
    """Encode a list to TOON format."""
    if not arr:
        return "[0]:"

    # Primitive array → inline CSV (§9.1)
    if all(_is_primitive(item) for item in arr):
        values = ",".join(_encode_value(v) for v in arr)
        return f"[{len(arr)}]: {values}"

    # Tabular array → header + rows (§9.3)
    if _is_tabular(arr):
        keys = list(arr[0].keys())
        header = f"[{len(arr)}]{{{','.join(keys)}}}:"
        rows = [",".join(_encode_value(item[k]) for k in keys) for item in arr]
        return header + "\n" + "\n".join(rows)

    # Mixed array → list format (§9.4)
    pad = "  " * indent
    lines = [f"[{len(arr)}]:"]
    for item in arr:
        if isinstance(item, dict):
            nested = _encode_dict(item, indent + 1).split("\n")
            lines.append(f"{pad}  - {nested[0]}")
            lines.extend(f"{pad}    {ln}" for ln in nested[1:])
        else:
            lines.append(f"{pad}  - {_encode_value(item)}")
    return "\n".join(lines)


def _encode_dict(obj: dict, indent: int = 0) -> str:
    """Encode a dict to TOON format."""
    if not obj:
        return ""

    lines: list[str] = []

    for key, value in obj.items():
        if isinstance(value, dict):
            lines.append(f"{key}:")
            nested = _encode_dict(value, indent + 1)
            lines.extend(f"  {ln}" for ln in nested.split("\n") if ln)
        elif isinstance(value, list):
            encoded = _encode_list(value, indent + 1)
            if "\n" in encoded:
                parts = encoded.split("\n")
                lines.append(f"{key}{parts[0]}")
                lines.extend(f"  {ln}" for ln in parts[1:])
            else:
                lines.append(f"{key}{encoded}")
        else:
            lines.append(f"{key}: {_encode_value(value)}")

    return "\n".join(lines)


def json_to_toon(data: dict | list) -> str:
    """Convert JSON-compatible data to TOON format.

    Args:
        data: A dict or list to encode.

    Returns:
        TOON-formatted string.
    """
    if isinstance(data, dict):
        return _encode_dict(data)
    if isinstance(data, list):
        return _encode_list(data)
    return _encode_value(data)
