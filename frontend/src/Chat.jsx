import React, { useState, useEffect, useRef } from "react";
import Message from "./Message";
import UserInput from "./UserInput";

export default function Chat() {
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);
  const messagesEndRef = useRef(null);

  useEffect(() => {
    ws.current = new WebSocket("wss://gotalk-ykyu.onrender.com/ws");

    ws.current.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      setMessages((prev) => [...prev, msg]);
    };

    ws.current.onopen = () => console.log("WebSocket connected");
    ws.current.onclose = () => console.log("WebSocket disconnected");

    return () => ws.current.close();
  }, []);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const sendMessage = (text) => {
    if (!text) return;
    const msg = { user: "Moi", text, timestamp: new Date().toISOString() };
    ws.current.send(JSON.stringify(msg));
    setMessages((prev) => [...prev, msg]);
  };

  return (
    <div className="chat-container">
      <div className="messages">
        {messages.map((msg, idx) => (
          <Message key={idx} message={msg} />
        ))}
        <div ref={messagesEndRef} />
      </div>
      <UserInput onSend={sendMessage} />
    </div>
  );
}
