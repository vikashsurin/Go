import axios from "axios";
import {
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  REGISTER_SUCCESS,
  REGISTER_FAIL,
  USER_LOADED,
  AUTH_ERROR,
  LOGOUT,
  CLEAR_PROFILE,
  REMOVE_ALERT
} from "./types";
import setAuthToken from "../../utils/setAuthToken";
//load user
export const loadUser = () => async dispatch => {
  const token = localStorage.getItem("token");
  if (token) {
    setAuthToken(token);
  }
  try {
    const res = await axios.get("/api/auth");
    dispatch({
      type: USER_LOADED,
      payload: res.data
    });
    // console.log('loaded', res.data);
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: AUTH_ERROR
    });
  }
};

//register a user
export const register = ({ name, email, password }) => async dispatch => {
  const config = {
    headers: {
      "Content-Type": "application/json"
    }
  };

  const body = JSON.stringify({ name, email, password });
  try {
    const res = await axios.post("http://localhost:8080/create/user", body, config);
    dispatch({
      type: REGISTER_SUCCESS,
      payload: res.data
    });
    console.log(res)
  } catch (err) {
    const errors = err.response;
    console.log(errors);
    dispatch({
      type: REGISTER_FAIL,
      payload: errors
    });
    setTimeout(() => {
      dispatch({
        type: REMOVE_ALERT,
        payload: ""
      });
    }, 3000);
  }
};

//login a user
export const login = ({ email, password }) => async dispatch => {
  const config = {
    headers: {
      "Content-Type": "application/json"
    }
  };
  const body = JSON.stringify({ email, password });
  try {
    const res = await axios.post("/login/user", body, config);
    dispatch({
      type: LOGIN_SUCCESS,
      payload: res.data
    });
    console.log(res)
  } catch (err) {
    // const errors = err.response.data;
    console.log(err);
    // setTimeout(() => {
    //   console.log("this is time out");
    // }, 3000);
    dispatch({
      type: LOGIN_FAIL,
      payload: err
    });
    setTimeout(() => {
      dispatch({
        type: REMOVE_ALERT,
        payload: ""
      });
    }, 3000);
  }
};

//logout user
export const logout = history => async dispatch => {
  await window.history.go(2);

  dispatch({ type: LOGOUT });
  dispatch({ type: CLEAR_PROFILE });
};
