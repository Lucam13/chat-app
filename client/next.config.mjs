/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
      BASE_URL: process.env.BASE_URL,
      BASE_URL_USERS: process.env.BASE_URL_USERS,
      BASE_URL_AREAS: process.env.BASE_URL_AREAS,
      BASE_URL_MESSAGES: process.env.BASE_URL_MESSAGES,
      BASE_URL_CHATS: process.env.BASE_URL_CHATS,
      BASE_URL_WEBSOCKET: process.env.BASE_URL_WEBSOCKET,
    },
  };
  
  export default nextConfig;
