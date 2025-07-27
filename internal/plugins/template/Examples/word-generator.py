#!/usr/bin/env python3
import sys
import json
import random

# A small set of words for demonstration!
WORD_LIST = [
    "apple", "banana", "cherry", "date", "elderberry",
    "fig", "grape", "honeydew", "kiwi", "lemon",
    "mango", "nectarine", "orange", "papaya", "quince",
    "raspberry", "strawberry", "tangerine", "ugli", "watermelon"
]

def generate_words(count):
    try:
        count = int(count)
        if count < 1:
            return json.dumps({"error": "Count must be positive"})
        
        # Generate random words
        words = random.sample(WORD_LIST, min(count, len(WORD_LIST)))
        
        # Return JSON formatted result
        return json.dumps({
            "words": words,
            "count": len(words)
        })
    except ValueError:
        return json.dumps({"error": "Invalid count parameter"})

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({"error": "Exactly one argument required"}))
        sys.exit(1)
    
    print(generate_words(sys.argv[1]))
