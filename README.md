# RPG Dialogue UI

A single-page RPG dialogue interface — Alphonse Mucha meets RPG Maker, dark
green window chrome on a creamy parchment backdrop. Built with React, Vite,
and Tailwind CSS v4.

Implements a design exported from Claude Design (`RPG Dialogue UI.dc.html` /
`rpg-dialogue.css`): an NPC portrait and nameplate up top, a quest tracker,
scrolling dialogue log, and an items/skills choice list in the main row, and
companion/player portraits flanking a free-text reply box at the bottom.
Sending a message or picking an action appends to the log and triggers a
delayed NPC reply, matching the original mockup's behavior.

## Develop

```bash
npm install
npm run dev
```

## Build

```bash
npm run build
```

## Structure

- `src/rpg-dialogue.css` — the Tailwind v4 theme (colors, fonts, window
  chrome, animations) extracted from the design.
- `src/components/` — `NpcHeader`, `QuestPanel`, `DialogueLog`,
  `ActionsPanel`, `BottomBar`.
- `src/useDialogue.ts` + `src/data.ts` — dialogue/quest/action state and the
  send/reply logic.
