import React, { useState } from "react";

export default function UserInput({ onSend }) {
  const [text, setText] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    onSend(text);
    setText("");
  };

  return (
    <form className="user-input" onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Écris ton message..."
        value={text}
        onChange={(e) => setText(e.target.value)}
      />
      <button type="submit">Envoyer</button>
    </form>
  );
}
