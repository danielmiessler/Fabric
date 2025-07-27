# Learning questionnaire generation

This pattern generates questions to help a learner/student review the main concepts of the learning objectives provided.

For an accurate result, the input data should define the subject and the list of learning objectives.

Example prompt input:

```
# Optional to be defined here or in the context file
[Student Level: High school student]

Subject: Machine Learning

Learning Objectives:
* Define machine learning
* Define unsupervised learning
```

# Example run bash:

Copy the input query to the clipboard and execute the following command:

```bash
xclip -selection clipboard -o | fabric -sp create_quiz
```

## Meta

- **Author**: Marc Andreu (marc@itqualab.com)
- **Version Information**: Marc Andreu's main `create_quiz` version.
- **Published**: May 6, 2024
