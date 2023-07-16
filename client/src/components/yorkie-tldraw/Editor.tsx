import { Tldraw, useFileSystem } from '@tldraw/tldraw';
import { useMultiplayerState } from '../../hooks/useMultiplayerState';
import CustomCursor from '../../CursorComponent';

export default function Editor() {
  const fileSystemEvents = useFileSystem();
  const { ...events } = useMultiplayerState(
    `tldraw-${new Date().toISOString().substring(0, 10).replace(/-/g, "")}`
  );
  const component = { Cursor: CustomCursor };

  return (
    <div className="tldraw">
      <Tldraw
        components={component}
        autofocus
        disableAssets={true}
        showPages={false}
        {...fileSystemEvents}
        {...events}
      />
    </div>
  );
}