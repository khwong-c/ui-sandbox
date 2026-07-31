import { useEffect, useRef, useState } from 'react';
import { NPC, PLAYER, actionDefs, initialLog, initialQuest, npcReplies } from './data';
import type { ActionDef, LogEntry } from './types';

export function useDialogue() {
  const [log, setLog] = useState<LogEntry[]>(initialLog);
  const [quest] = useState(initialQuest);
  const [input, setInput] = useState('');
  const logRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const el = logRef.current;
    if (el) el.scrollTop = el.scrollHeight;
  }, [log]);

  function addEntry(speaker: string, kind: LogEntry['kind'], text: string) {
    setLog((prev) => [...prev, { speaker, kind, text }]);
  }

  function queueNpcReply() {
    setTimeout(() => {
      addEntry(NPC, 'npc', npcReplies[Math.floor(Math.random() * npcReplies.length)]);
    }, 650);
  }

  function handleSend() {
    const text = input.trim();
    if (!text) return;
    addEntry(PLAYER, 'player', text);
    setInput('');
    queueNpcReply();
  }

  function handleAction(action: ActionDef) {
    addEntry(PLAYER, 'player', action.reply);
    queueNpcReply();
  }

  return {
    log,
    quest,
    actions: actionDefs,
    input,
    setInput,
    logRef,
    handleSend,
    handleAction,
  };
}
