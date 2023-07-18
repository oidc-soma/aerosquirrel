import {
  Tldraw,
  useFileSystem,
  TDUserStatus,
  TDAsset,
  TDBinding,
  TDShape,
  TDUser,
  TldrawApp,
} from "@tldraw/tldraw";
import { createShapes } from "@tldraw/tldraw/dist/state/commands";
import { useMultiplayerState } from '../../hooks/useMultiplayerState';
import CustomCursor from '../../CursorComponent';
import { useState, useCallback } from 'react';

export default function Editor() {
   const parentFunction = (x: any) => {
     console.log(x);
    //createShapes(x, 
   };
  const fileSystemEvents = useFileSystem();
  const { ...events } = useMultiplayerState(
    // `tldraw-${new Date().toISOString().substring(0, 10).replace(/-/g, "")}`
    'tldraw-Saver', {parentFunction}
  );
  
 

  const component = { 
    Cursor: CustomCursor };

  return (
    <>
    
      <button style={{ zIndex: 10000}}>테스트용 버튼</button>
      <div className="tldraw" style={{position: 'absolute', top: '200px', width: '100%', height: '600px'}}>
        <Tldraw
          components={component}
          autofocus
          disableAssets={false}
          showPages={false}
          {...fileSystemEvents}
          {...events}
        />
      </div>
    </>
  );
}