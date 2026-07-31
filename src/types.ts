export type SpeakerKind = 'npc' | 'companion' | 'player';

export interface LogEntry {
  speaker: string;
  kind: SpeakerKind;
  text: string;
}

export interface QuestItem {
  label: string;
  done: boolean;
}

export interface ActionDef {
  id: string;
  label: string;
  hint: string;
  reply: string;
}
