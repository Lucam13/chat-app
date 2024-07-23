require("dotenv").config()
import axios from 'axios';

const { BASE_URL_USERS } = process.env;

export const GET_USER = 'GET_USER';
export const GET_ALL_USERS = 'GET_ALL_USERS';

export const getUser = (userId) => async (dispatch) => {
  try {
    const { data } = await axios.get(`${BASE_URL_USERS}/${userId}`);
    dispatch({
      type: GET_USER,
      payload: data,
    });
  } catch (error) {
    console.error(error);
  }
};

export const getAllUsers = () => async (dispatch) => {
    try {
        const { data } = await axios.get(BASE_URL_USERS);
        dispatch({
        type: GET_ALL_USERS,
        payload: data,
        });
    } catch (error) {
        console.error(error);
    }
}

