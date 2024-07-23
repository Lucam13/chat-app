import React from 'react';

// Message component

const Message = ({ text, user_id }) => {
    const isUser = user_id === 'user';
    return (
      <div className={`flex ${isUser ? 'justify-end' : 'justify-start'} mb-2`}>
        <div
          className={`p-2 rounded ${
            isUser ? 'bg-orange-400 text-white' : 'bg-green-500 text-white'
          }`}
        >
          {text}
        </div>
      </div>
    );
  };
  
  export default Message;

  