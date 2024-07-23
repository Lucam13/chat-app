import { GET_ALL_AREAS } from "../actions/area";
import { GET_AREA } from "../actions/area";

const initialState = {
    area: [],
};

export default (state = initialState, action) => {
    switch (action.type) {
    case GET_ALL_AREAS:
        return {
        ...state,
        area: action.payload,
        };
    case GET_AREA:
        return {
        ...state,
        area: action.payload,
        };
    
    default:
        return state;
    }
};




