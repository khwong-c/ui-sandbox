# CLAUDE.md

This file gives guidance to Claude Code (claude.ai/code) for work in this repository.

## Commands

```bash
pnpm dev          # start the dev server
pnpm build        # type-check with tsc -b, then bundle with Vite
pnpm lint         # run oxlint (not ESLint)
pnpm preview      # preview the production build
```

There is no test runner.

## Architecture

This application is a React 19 + TypeScript + Vite SPA. It shows a dungeon-themed RPG chat interface for negotiation.

**Start sequence**: `main.tsx` → `App.tsx` → `views/deal-chat/DealChat.tsx`

`App` shows a full-screen background image. It mounts `DealChat`, which is the only view.

`DealChat` contains these components:
- `Dialog` — a scrollable chat log panel. It accepts a `chatItems[]` array. The interface is not yet defined (TODO stub).
- `ChatItem` — a single chat bubble. It renders inside `Dialog`. The interface is not yet defined (stub).
- `Avatar` — a character portrait. It uses the `.rpg-face` frame. It takes a `src` prop.

**Image files**: Character images are in `src/assets/<character>/avatar-NNN.png`. The three-digit number shows the emotional state. Each character has more than one image for different expressions.

## Style System

Tailwind v4 uses a CSS-first approach. There is no `tailwind.config.js`. All custom tokens and component classes are in `src/dialogue.css`. The file `src/index.css` imports `dialogue.css`.

**Color tokens** — use them as Tailwind utilities or CSS variables:
- `cream-*` — parchment and gold tones for surfaces and text
- `forest-*` — dark green for windows, borders, and primary text

**Component classes** — use these classes instead of ad-hoc Tailwind when the pattern matches:

| Class | Purpose |
|---|---|
| `.rpg-window` | Dark-green gradient panel with a forest border and a cream inset shadow |
| `.rpg-face` | Avatar portrait frame with a border and hidden overflow |
| `.rpg-bubble` | Chat message. Combine with `.rpg-bubble-npc`, `.rpg-bubble-player`, or `.rpg-bubble-whisper`. |
| `.rpg-log` | Scrollable container with a custom scrollbar style |
| `.rpg-input` | Styled text field |
| `.rpg-button` | Parchment-colored button (Cinzel font, uppercase) |
| `.rpg-nameplate` | Character name banner |
| `.rpg-choice` | Hoverable choice row with a ▶ cursor animation |
| `.rpg-speaker` | Label above a bubble. Combine with `-player` or `-whisper` variants. |

**Typography**: Use `--font-display` (Cinzel, serif) for nameplates, buttons, and titles. Use `--font-body` (EB Garamond, serif) for dialogue text.

**Animations**: `slowspin`, `glowpulse`, and `bobarrow` are in `@theme`. Use them as `animate-slowspin`, `animate-glowpulse`, and `animate-bobarrow`.
