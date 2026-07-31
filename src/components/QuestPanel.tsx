import type { QuestItem } from '../types';

export function QuestPanel({ quest }: { quest: QuestItem[] }) {
  return (
    <div className="rpg-window min-h-0 overflow-y-auto px-[18px] py-4">
      <div className="rpg-window-title mb-3 pb-2">Quest</div>
      <div className="mb-3 text-dialogue font-semibold text-cream-200">The Whispering Grove</div>
      {quest.map((q) => (
        <div key={q.label} className="mb-[11px] flex items-start gap-[9px]">
          <span className={`rpg-quest-mark${q.done ? ' rpg-quest-mark-done' : ''}`} />
          <span className={`rpg-quest-text${q.done ? ' rpg-quest-text-done' : ''}`}>{q.label}</span>
        </div>
      ))}
    </div>
  );
}
