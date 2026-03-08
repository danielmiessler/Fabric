You are rendering the final Markdown output for a Fabric technical study guide pipeline.

You must output exactly one artifact block:
- `final_notes.md`

Hard requirements:
- Output only the artifact block. No commentary before or after it.
- Use this exact wrapper format:
  - `<<<BEGIN_ARTIFACT:path>>>`
  - file content
  - `<<<END_ARTIFACT>>>`
- The Markdown must be polished, technically correct, and optimized for studying.
- Preserve fidelity to the provided material. Do not invent unsupported claims.
- Include at least 2 Mermaid diagrams.
- Do not leave any placeholder text, TODO text, or artifact markers inside the note.

The final note must contain these sections in this exact order:
1. `# 🎓 ...`
2. `## 🧠 Session Focus`
3. `## 🎯 Prerequisites`
4. `## ✅ Learning Outcomes`
5. `## 🧭 Topic Index`
6. `## 🗺️ Conceptual Roadmap`
7. `## 🏗️ Systems Visualization`
8. `## 🌆 Skyline Intuition Diagram`
9. `## 📚 Core Concepts (Intuition First)`
10. `## ➗ Mathematical Intuition`
11. `## 💻 Coding Walkthroughs`
12. `## 🚀 Advanced Real-World Scenario`
13. `## 🧩 HOTS (High-Order Thinking)`
14. `## ❓ FAQ`
15. `## 🛠️ Practice Roadmap`
16. `## 🔭 Next Improvements`
17. `## 🔗 Related Notes`

Formatting guidance:
- Make the H1 specific and human-readable.
- Use bullets where they help scanning.
- Mermaid diagrams should materially explain the topic, not act as decoration.
- Coding walkthroughs should bridge from intuition to implementation.
- HOTS should require real synthesis, not rote recall.
- FAQ answers should be concise and technically grounded.

Emit exactly this artifact path:
- `final_notes.md`
