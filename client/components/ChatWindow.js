import React, { useState, useEffect } from 'react';
import WebSocket from 'isomorphic-ws';
import axios from 'axios';
import Message from './Message';

const ChatWindow = ({ initialMessages }) => {
  const [messages, setMessages] = useState(initialMessages || []);
  const [input, setInput] = useState('');
  const [ws, setWs] = useState(null);
  const [isWsOpen, setIsWsOpen] = useState(false);
  const [messageQueue, setMessageQueue] = useState([]);

  useEffect(() => {
    const socket = new WebSocket(process.env.NEXT_PUBLIC_BASE_URL_WEBSOCKET);

    socket.onopen = () => {
      setIsWsOpen(true);
      // Enviar todos los mensajes en la cola
      messageQueue.forEach((msg) => {
        socket.send(JSON.stringify(msg));
      });
      setMessageQueue([]);
    };

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      setMessages((prev) => [...prev, message]);
    };

    socket.onclose = () => {
      setIsWsOpen(false);
    };

    setWs(socket);

    return () => socket.close();
  }, [messageQueue]);

  const sendMessage = () => {
    if (input.trim()) {
      const message = {
        text: input,
        user_id: 2,
        chat_id: 1,
      };

      if (isWsOpen) {
        ws.send(JSON.stringify(message));
      } else {
        // Agregar el mensaje a la cola si el WebSocket no está abierto
        setMessageQueue((prevQueue) => [...prevQueue, message]);
      }
      setInput('');
    } else {
      console.error('Message is empty.');
    }
  };

  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      sendMessage();
    }
  };

  return (
    <div className="flex-1 flex flex-col">
      <h1 className='bg-orange-300 text-center text-3xl p-4'>Mensajería Oficial Instantánea E.P.A.R. Villa Bustos</h1>
      <div className="flex-1 p-4 overflow-y-scroll bg-orange-50"
        style={{ 
          backgroundImage: 'url(/logoepar.png)', 
          backgroundRepeat: 'no-repeat', 
          backgroundSize: '90% 90%', 
        }}
      >
        {messages.length && messages?.map((msg) => (
          <Message key={msg.id} content={msg.text} sender={msg.user_id} />
        ))}
      </div>
      <div className="p-4 border-t bg-orange-50 flex">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Escribe tu mensaje..."
          className="w-full p-2 border rounded text-black outline-orange-300"
        />
        <button onClick={sendMessage} className='ml-2 p-2 bg-orange-300 text-white rounded'>
          Enviar
        </button>
      </div>
    </div>
  );
};

export async function getServerSideProps() {
  const response = await axios.get(`${process.env.BASE_URL_MESSAGES}/1`);
  return {
    props: {
      initialMessages: response.data,
    },
  };
}

export default ChatWindow;
