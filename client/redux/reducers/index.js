import { combineReducers } from 'redux';
import user from './user';
import chat from './chat';
import message from './message';
import area from './area';

const rootReducer = combineReducers({
    userStore: user,
    chatSore: chat,
    messageStore: message,
    areaStore: area,
    });

export default rootReducer;