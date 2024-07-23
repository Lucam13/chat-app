import {
    GET_MESSAGE,
    GET_MESSAGE_LIST,
    GET_USER_MESSAGE,
    CREATE_MESSAGE,
} from "../actions/message";

const initialState = {
    message: {},
    messageList: [],
    userMessage: [],
};

export default (state = initialState, action) => {
    switch (action.type) {
    case GET_MESSAGE:
        return {
            ...state,
            message: action.payload,
        };
    case GET_MESSAGE_LIST:
        return {
            ...state,
            messageList: action.payload,
        };
    case GET_USER_MESSAGE:
        return {
            ...state,
            userMessage: action.payload,
        };
    case CREATE_MESSAGE:
        return {
            ...state,
            messageList: [...state.messageList, action.payload],
        };
        default:
        return state;
    }
};
