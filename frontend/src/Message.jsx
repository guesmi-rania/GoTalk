import React from "react";

export default function Message({ message }) {
  const { user, text, timestamp } = message;
  const time = new Date(timestamp).toLocaleTimeString();
  return (
    <div className={`message ${user === "Moi" ? "my-message" : "other-message"}`}>
      <div className="message-header">
        <strong>{user}</strong> <span>{time}</span>
      </div>
      <p>{text}</p>
    </div>
  );
}
