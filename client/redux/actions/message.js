require("dotenv").config()
import axios from 'axios';

const { BASE_URL_MESSAGES } = process.env;

export const GET_MESSAGES_BY_CHAT_ID = 'GET_MESSAGES_BY_CHAT_ID';
export const CREATE_MESSAGE = 'CREATE_MESSAGE';

export const getMessagesByChatId = (chat_id) => async (dispatch) => {
    try {
        const { data } = await axios.get(`${BASE_URL_MESSAGES}/${chat_id}`);
        dispatch({
        type: GET_MESSAGES_BY_CHAT_ID,
        payload: data,
        });
    } catch (error) {
        console.error(error);
    }
}

export const createMessage = ({ chat_id, date_created, user_id, status, text}) => async (dispatch) => {
  try {
    const { data } = await axios.post(BASE_URL_MESSAGES, { chat_id, date_created, user_id, status, text});
    dispatch({
      type: CREATE_MESSAGE,
      payload: data,
    });
  } catch (error) {
    console.error(error);
  }
};  
