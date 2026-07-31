export function NpcHeader() {
  return (
    <div className="relative z-10 flex items-end gap-4 px-[30px] pt-[18px] pb-3">
      <div className="rpg-window relative flex h-[108px] w-[108px] flex-none items-center justify-center overflow-hidden">
        <span className="text-center font-mono text-[9px] leading-[1.6] tracking-[.06em] text-[rgba(240,227,189,0.6)]">
          GOBLIN
          <br />
          FACE
          <br />
          GRAPHIC
        </span>
      </div>

      <div className="flex flex-col gap-2 pb-1">
        <div className="inline-flex items-center gap-2.5 self-start rounded-[6px] border-[3px] border-forest-600 bg-gradient-to-b from-forest-500 to-forest-700 px-4 py-1.5 shadow-[0_3px_0_rgba(31,59,42,0.25),inset_0_0_0_2px_var(--color-cream-500)]">
          <span className="font-display text-[19px] font-bold tracking-[.05em] text-cream-200">MOSSGRIN</span>
          <span className="rpg-lozenge" />
          <span className="text-xs uppercase tracking-[.1em] text-[rgba(245,234,208,0.7)]">Lantern Warden</span>
        </div>

        <div className="rpg-pill inline-flex items-center gap-[7px] self-start px-2.5 py-0.5">
          <span className="h-1.5 w-1.5 animate-glowpulse rounded-full bg-forest-600" />
          <span>Mood — Wary</span>
        </div>
      </div>
    </div>
  );
}
