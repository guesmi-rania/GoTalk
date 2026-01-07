import React from "react";
import Chat from "./Chat";

export default function App() {
  return (
    <div className="app">
      <header className="app-header">
        <h1>GoTalk Chat 💬</h1>
      </header>
      <main>
        <Chat />
      </main>
      <footer className="app-footer">
        <p>&copy; 2026 GoTalk</p>
      </footer>
    </div>
  );
}
