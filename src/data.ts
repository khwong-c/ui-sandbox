import type { ActionDef, LogEntry, QuestItem } from './types';

export const NPC = 'Mossgrin';
export const COMPANION = 'Suzu';
export const PLAYER = 'Kaida';

function kindFor(speaker: string): LogEntry['kind'] {
  return speaker === PLAYER ? 'player' : speaker === COMPANION ? 'companion' : 'npc';
}

export const initialLog: LogEntry[] = [
  { speaker: NPC, text: 'Halt, traveler. The shrine road is not for wandering feet.' },
  { speaker: COMPANION, text: '(whispers) He is testing you. Choose your words carefully.' },
  { speaker: PLAYER, text: 'I mean no disrespect. I seek passage to the Whispering Grove.' },
  { speaker: NPC, text: 'The Grove does not open for the empty-handed.' },
  { speaker: COMPANION, text: '(whispers) Show him the token, quickly.' },
  { speaker: PLAYER, text: 'I carry an offering — a sprig of silver moss.' },
  { speaker: NPC, text: '...Silver moss. It has been many moons since one remembered the old rite.' },
  { speaker: NPC, text: 'Perhaps you are not so foolish as you appear.' },
  { speaker: COMPANION, text: '(whispers) He is softening. Keep going.' },
  { speaker: NPC, text: 'Speak your purpose, and I shall judge if the Grove wills to hear it.' },
].map((e) => ({ ...e, kind: kindFor(e.speaker) }));

export const initialQuest: QuestItem[] = [
  { label: 'Find the shrine keeper', done: true },
  { label: 'Present a worthy offering', done: false },
  { label: 'Gain passage to the Grove', done: false },
];

export const actionDefs: ActionDef[] = [
  { id: 'offer', label: 'Offer Silver Moss', hint: 'Item ×1 — forest token', reply: 'I offer this silver moss, gathered beneath the old oak.' },
  { id: 'rite', label: 'Recite the Old Rite', hint: 'Skill — Lore 4', reply: 'I recite the old rite, the words my grandmother taught me.' },
  { id: 'ward', label: 'Show Spirit Ward', hint: 'Item ×1 — protective charm', reply: 'I raise the spirit ward so he may see I mean no harm.' },
  { id: 'persuade', label: 'Persuade', hint: 'Skill — Charm 3', reply: 'Warden, I ask only for safe passage. Nothing more.' },
];

export const npcReplies = [
  'Mossgrin studies you a moment longer, lantern light flickering in his eyes.',
  "The old goblin's posture eases, just slightly.",
  'Hmph. Few remember such things anymore.',
  'Very well. The Grove may yet hear your plea.',
];
