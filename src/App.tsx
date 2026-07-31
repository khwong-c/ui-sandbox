import { ActionsPanel } from './components/ActionsPanel';
import { BottomBar } from './components/BottomBar';
import { DialogueLog } from './components/DialogueLog';
import { NpcHeader } from './components/NpcHeader';
import { QuestPanel } from './components/QuestPanel';
import { useDialogue } from './useDialogue';

function App() {
  const { log, quest, actions, input, setInput, logRef, handleSend, handleAction } = useDialogue();

  return (
    <div className="rpg-page relative grid h-screen w-screen grid-rows-[auto_minmax(260px,1fr)_auto] overflow-hidden font-body text-forest-600">
      <div className="rpg-texture-dots pointer-events-none absolute inset-0" />
      <div className="rpg-sunburst pointer-events-none absolute top-[-140px] left-1/2 h-[520px] w-[520px] -translate-x-1/2 animate-slowspin" />

      <NpcHeader />

      <div className="relative z-10 grid min-h-0 grid-cols-[264px_1fr_296px] gap-5 px-[30px]">
        <QuestPanel quest={quest} />
        <DialogueLog log={log} logRef={logRef} />
        <ActionsPanel actions={actions} onAction={handleAction} />
      </div>

      <BottomBar input={input} setInput={setInput} onSend={handleSend} />
    </div>
  );
}

export default App;
