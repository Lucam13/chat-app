require("dotenv").config()
import axios from 'axios';

const { baseUrlChats: BASE_URL_CHATS } = process.env;

export const GET_CHATS_BY_AREA_ID = 'GET_CHATS_BY_AREA_ID';
export const GET_ALL_CHATS = 'GET_ALL_CHATS';

export const getChatsByAreaId = (areaId) => async (dispatch) => {
  try {
    const { data } = await axios.get(`${BASE_URL_CHATS}/${areaId}`);
    dispatch({
      type: GET_CHATS_BY_AREA_ID,
      payload: data,
    });
  } catch (error) {
    console.error(error);
  }
};

export const getAllChats = () => async (dispatch) => {
  try {
    const { data } = await axios.get(BASE_URL_CHATS);
    dispatch({
      type: GET_ALL_CHATS,
      payload: data,
    });
  } catch (error) {
    console.error(error);
  }
};
