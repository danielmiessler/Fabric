#!/usr/bin/env python3

"""Extracts pattern information from the ~/.config/fabric/patterns directory,
creates JSON files for pattern extracts and descriptions, and updates web static files.
"""
import os
import json
import shutil


def load_existing_file(filepath):
    """Load existing JSON file or return default structure"""
    if os.path.exists(filepath):
        try:
            with open(filepath, "r", encoding="utf-8") as f:
                return json.load(f)
        except json.JSONDecodeError:
            print(
                f"Warning: Malformed JSON in {filepath}. Starting with an empty list."
            )
            return {"patterns": []}
    return {"patterns": []}


def get_pattern_extract(pattern_path):
    """Extract first 500 words from pattern's system.md file"""
    system_md_path = os.path.join(pattern_path, "system.md")
    with open(system_md_path, "r", encoding="utf-8") as f:
        content = " ".join(f.read().split()[:500])
    return content


def extract_pattern_info():
    """Extract pattern information from the patterns directory"""
    script_dir = os.path.dirname(os.path.abspath(__file__))
    patterns_dir = os.path.expanduser("~/.config/fabric/patterns")
    print(f"\nScanning patterns directory: {patterns_dir}")

    extracts_path = os.path.join(script_dir, "pattern_extracts.json")
    descriptions_path = os.path.join(script_dir, "pattern_descriptions.json")

    existing_extracts = load_existing_file(extracts_path)
    existing_descriptions = load_existing_file(descriptions_path)

    existing_extract_names = {p["patternName"] for p in existing_extracts["patterns"]}
    existing_description_names = {
        p["patternName"] for p in existing_descriptions["patterns"]
    }
    print(f"Found existing patterns: {len(existing_extract_names)}")

    new_extracts = []
    new_descriptions = []

    for dirname in sorted(os.listdir(patterns_dir)):
        pattern_path = os.path.join(patterns_dir, dirname)
        system_md_path = os.path.join(pattern_path, "system.md")

        if os.path.isdir(pattern_path) and os.path.exists(system_md_path):
            if dirname not in existing_extract_names:
                print(f"Processing new pattern: {dirname}")

            try:
                if dirname not in existing_extract_names:
                    print(f"Creating new extract for: {dirname}")
                    pattern_extract = get_pattern_extract(
                        pattern_path
                    )  # Pass directory path
                    new_extracts.append(
                        {"patternName": dirname, "pattern_extract": pattern_extract}
                    )

                if dirname not in existing_description_names:
                    print(f"Creating new description for: {dirname}")
                    new_descriptions.append(
                        {
                            "patternName": dirname,
                            "description": "[Description pending]",
                            "tags": [],
                        }
                    )

            except OSError as e:
                print(f"Error processing {dirname}: {str(e)}")
        else:
            print(f"Invalid pattern directory or missing system.md: {dirname}")

    print("\nProcessing summary:")
    print(f"New extracts created: {len(new_extracts)}")
    print(f"New descriptions added: {len(new_descriptions)}")

    existing_extracts["patterns"].extend(new_extracts)
    existing_descriptions["patterns"].extend(new_descriptions)

    return existing_extracts, existing_descriptions, len(new_descriptions)


def update_web_static(descriptions_path):
    """Copy pattern descriptions to web static directory"""
    script_dir = os.path.dirname(os.path.abspath(__file__))
    static_dir = os.path.join(script_dir, "..", "..", "web", "static", "data")
    os.makedirs(static_dir, exist_ok=True)
    static_path = os.path.join(static_dir, "pattern_descriptions.json")
    shutil.copy2(descriptions_path, static_path)


def save_pattern_files():
    """Save both pattern files and sync to web"""
    script_dir = os.path.dirname(os.path.abspath(__file__))
    extracts_path = os.path.join(script_dir, "pattern_extracts.json")
    descriptions_path = os.path.join(script_dir, "pattern_descriptions.json")

    pattern_extracts, pattern_descriptions, new_count = extract_pattern_info()

    # Save files
    with open(extracts_path, "w", encoding="utf-8") as f:
        json.dump(pattern_extracts, f, indent=2, ensure_ascii=False)

    with open(descriptions_path, "w", encoding="utf-8") as f:
        json.dump(pattern_descriptions, f, indent=2, ensure_ascii=False)

    # Update web static
    update_web_static(descriptions_path)

    print("\nProcessing complete:")
    print(f"Total patterns: {len(pattern_descriptions['patterns'])}")
    print(f"New patterns added: {new_count}")


if __name__ == "__main__":
    save_pattern_files()
