import type { ActionDef } from '../types';

export function ActionsPanel({
  actions,
  onAction,
}: {
  actions: ActionDef[];
  onAction: (action: ActionDef) => void;
}) {
  return (
    <div className="rpg-window min-h-0 overflow-y-auto px-[14px] py-4">
      <div className="rpg-window-title mb-2.5 pb-2">Items &amp; Skills</div>

      {actions.map((act) => (
        <div key={act.id} className="rpg-choice mb-1" onClick={() => onAction(act)}>
          <span className="rpg-cursor mt-0.5 text-sm leading-[1.4]">▶</span>
          <div>
            <div className="rpg-choice-label">{act.label}</div>
            <div className="rpg-choice-hint mt-px">{act.hint}</div>
          </div>
        </div>
      ))}

      <div className="mt-3 border-t-2 border-[rgba(232,217,166,0.25)] pt-2.5 text-caption text-[rgba(245,234,208,0.55)] italic">
        Choose an item or skill, or speak freely below.
      </div>
    </div>
  );
}
