import type { RefObject } from 'react';
import type { LogEntry } from '../types';

function bubbleClass(kind: LogEntry['kind']) {
  if (kind === 'player') return 'rpg-bubble rpg-bubble-player';
  if (kind === 'companion') return 'rpg-bubble rpg-bubble-whisper';
  return 'rpg-bubble rpg-bubble-npc';
}

function speakerClass(kind: LogEntry['kind']) {
  if (kind === 'player') return 'rpg-speaker rpg-speaker-player';
  if (kind === 'companion') return 'rpg-speaker rpg-speaker-whisper';
  return 'rpg-speaker';
}

export function DialogueLog({ log, logRef }: { log: LogEntry[]; logRef: RefObject<HTMLDivElement | null> }) {
  return (
    <div className="rpg-window flex min-h-0 flex-col overflow-hidden">
      <div className="flex items-center justify-between border-b-2 border-[rgba(232,217,166,0.35)] px-5 pt-[14px] pb-[10px] font-display text-[13px] font-bold tracking-[.14em] text-cream-500 uppercase">
        <span>Dialogue Log</span>
        <span className="text-[11px] tracking-[.08em] text-[rgba(245,234,208,0.55)]">scroll for history ▲</span>
      </div>
      <div ref={logRef} className="rpg-log flex min-h-0 flex-1 flex-col gap-3 px-5 py-4">
        {log.map((entry, i) => (
          <div key={i} className={`flex ${entry.kind === 'player' ? 'justify-end' : 'justify-start'}`}>
            <div className={bubbleClass(entry.kind)}>
              <div className={`${speakerClass(entry.kind)} mb-[3px]`}>{entry.speaker}</div>
              <div>{entry.text}</div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
