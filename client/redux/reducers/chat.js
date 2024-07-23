import { GET_CHAT, GET_ALL_CHAT, CREATE_CHAT } from "../actions/chat";

const initialState = {
    chat: {},
    chatList: [],
};

export default (state = initialState, action) => {
    switch (action.type) {
    case GET_CHAT:
        return {
        ...state,
        chat: action.payload,
        };
    case GET_ALL_CHAT:
        return {
        ...state,
        chatList: action.payload,
        };
    case CREATE_CHAT:
        return {
        ...state,
        chatList: [...state.chatList, action.payload],
        };
    default:
        return state;
    }
};
