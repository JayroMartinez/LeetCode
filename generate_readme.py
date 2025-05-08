import os
import re

README_HEADER = """# LeetCode Solutions in Python

This repository contains my personal solutions to problems from [LeetCode](https://leetcode.com/), written in Python 3.

---

## Problems

| #   | Title | Link | File |
|-----|-------|------|------|
"""

def extract_info(filename):
    match = re.match(r"(\d+)_([a-zA-Z0-9_]+)\.py", filename)
    if match:
        number = int(match.group(1))
        raw_title = match.group(2).replace('_', ' ')
        url_title = match.group(2).lower().replace('_', '-')
        link = f"https://leetcode.com/problems/{url_title}/"
        return (number, raw_title.title(), link, filename)
    return None

def main():
    files = [f for f in os.listdir('.') if f.endswith(('.py', '.go'))]
    problems = []

    for f in files:
        info = extract_info(f)
        if info:
            problems.append(info)

    problems.sort(key=lambda x: x[0])

    with open("README.md", "w") as f:
        f.write(README_HEADER)
        for number, title, link, filename in problems:
            f.write(f"| {number} | {title} | [🔗]({link}) | `{filename}` |\n")

if __name__ == "__main__":
    main()
