require("dotenv").config()

const BASE_URL_AREAS = process.env.BASE_URL_AREAS;

import axios from 'axios';

export const GET_ALL_AREAS = 'GET_ALL_AREAS';
export const GET_AREA = 'GET_AREA';


export const getAllAreas = () =>  async (dispatch) => {
  try {
    const { data } = await axios.get(process.env.NEXT_PUBLIC_BASE_URL_AREAS);
    dispatch({
      type: GET_ALL_AREAS,
      payload: data,
    });
  } catch (error) {
    console.error(error);
  }

};

export const getAreabyId = (areaId) => async (dispatch) => {
  try {
    const { data } = await axios.get(`${BASE_URL_AREAS}/${areaId}`);
    dispatch({
      type: GET_AREA,
      payload: data,
    });
  } catch (error) {
    console.error(error);
  }
}
