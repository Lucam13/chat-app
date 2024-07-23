import {GET_USER, GET_ALL_USERS} from '../actions/user';


const initialState = {
    user: {},
    };

export default (state = initialState, action) => {
    switch (action.type) {
    case GET_USER:
        return {
            ...state,
            user: action.payload,
        };
    case GET_ALL_USERS:
        return {
            ...state,
            user: action.payload,
        };    
    default:
        return state;
    }
}
