import ChatLayout from '../components/ChatLayout';
import ChatList from '../components/ChatList';
import ChatWindow from '../components/ChatWindow';

const Home = () => {
  return (
    <ChatLayout>
      <ChatList />
      <ChatWindow />
    </ChatLayout>
  );
};

export default Home;