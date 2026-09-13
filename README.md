# Vanilla Tools

This repository is a collection of useful tools for Vanilla OS.

## Tools

- `cur-gpu` - A minimal glxinfo re-implementation to get the current GPU/Driver and optionally the GL extensions.
- `lpkg` - A package manager locker/unlocker for Vanilla OS.
- `nrun` - A simple utility to add the environment variables needed to run a program with a NVIDIA GPU.
- `prime-switch` - Utility to switch PRIME profiles on a Debian + ABRoot based system.

## Use of Generative AI

Maintainers may use generative AI tools as assistants while working on vanilla-tools. Non-trivial assisted commits disclose the tool, model, and scope of the work.

AI tools may assist with code comments, documentation, repetitive code, and issue triage. Maintainers make project decisions and review every assisted change before it is merged.

Use these trailers for non-trivial assisted commits:

```plain
Assisted-by: <tool>:<model-version>
AI-Scope: <what the tool generated and the prompt or a short prompt summary>
```

Single-line completions, renames, and formatting changes do not need trailers.

Coding agents must also follow [AGENTS.md](AGENTS.md) before changing files,
creating commits, or opening pull requests.
