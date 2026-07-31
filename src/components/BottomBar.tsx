import type { ChangeEvent, KeyboardEvent } from 'react';

export function BottomBar({
  input,
  setInput,
  onSend,
}: {
  input: string;
  setInput: (value: string) => void;
  onSend: () => void;
}) {
  function handleKeyDown(e: KeyboardEvent<HTMLInputElement>) {
    if (e.key === 'Enter') onSend();
  }

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    setInput(e.target.value);
  }

  return (
    <div className="relative z-10 flex items-end gap-5 px-[30px] pt-4 pb-5">
      <div className="flex w-[132px] flex-none flex-col items-center gap-1.5">
        <div className="rpg-face rpg-face-companion flex h-[84px] w-[84px] items-center justify-center">
          <span className="text-center font-mono text-[8px] leading-[1.6] text-[rgba(240,227,189,0.6)]">
            FOX
            <br />
            SPIRIT
          </span>
        </div>
        <div className="rpg-nameplate px-2.5 py-[3px] text-center">
          <div className="text-[13px] tracking-[.04em]">SUZU</div>
        </div>
        <div className="rpg-role text-[10px]">Companion</div>
      </div>

      <div className="rpg-window flex flex-1 min-w-0 flex-col gap-2 px-4 py-3">
        <div className="flex items-center gap-2">
          <span className="rpg-cursor text-xs">▶</span>
          <span className="font-display text-xs tracking-[.12em] text-cream-500 uppercase">Kaida's reply</span>
        </div>
        <div className="flex gap-3">
          <input
            className="rpg-input flex-1"
            value={input}
            onChange={handleChange}
            onKeyDown={handleKeyDown}
            placeholder="Speak your words to the Warden..."
          />
          <button className="rpg-button" onClick={onSend}>
            Send
          </button>
        </div>
      </div>

      <div className="flex w-[132px] flex-none flex-col items-center gap-1.5">
        <div className="rpg-face flex h-[84px] w-[84px] items-center justify-center">
          <span className="text-center font-mono text-[8px] leading-[1.6] text-[rgba(240,227,189,0.6)]">
            PLAYER
            <br />
            FACE
          </span>
        </div>
        <div className="rpg-nameplate px-2.5 py-[3px] text-center">
          <div className="text-[13px] tracking-[.04em]">KAIDA</div>
        </div>
        <div className="rpg-role text-[10px]">Wanderer</div>
      </div>
    </div>
  );
}
