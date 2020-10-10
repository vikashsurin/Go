import axios from 'axios';
import {
  CREATE_PROFILE,
  PROFILE_ERROR,
  CURRENT_PROFILE,
  CLEAR_PROFILE,
  ACCOUNT_DELETED
} from './types';
import setAuthToken from '../../utils/setAuthToken';

//create a profile
export const createProfile = (profileData, history) => async dispatch => {
  try {
    const token = localStorage.token;
    setAuthToken(token);

    const config = {
      headers: {
        'Content-Type': 'application/json'
      }
    };
    const res = await axios.put('/api/profiles', profileData, config);

    dispatch({
      type: CREATE_PROFILE,
      payload: res.data
    });
    history.push('/dashboard');
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: PROFILE_ERROR
    });
  }
};
//update a profile
export const updateProfile = (profileData, history) => async dispatch => {
  try {
    setAuthToken(localStorage.token);

    const config = {
      headers: {
        'Content-Type': 'application/json'
      }
    };

    const res = await axios.put('/api/profiles', profileData, config);
    dispatch({
      type: CREATE_PROFILE,
      payload: res.data
    });
    console.log(res.data);
    history.push('/dashboard');
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: PROFILE_ERROR
    });
  }
};

//get current profile
export const getCurrentProfile = () => async dispatch => {
  const token = localStorage.token;
  setAuthToken(token);
  try {
    const res = await axios.get('/api/profiles/me');
    dispatch({
      type: CURRENT_PROFILE,
      payload: res.data
    });
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: PROFILE_ERROR
    });
    dispatch({
      type: CLEAR_PROFILE
    });
  }
};

//delete profile and user
export const deleteAccount = () => async dispatch => {
  const token = localStorage.token;
  try {
    setAuthToken(token);

    await axios.delete('/api/profiles');
    dispatch({
      type: CLEAR_PROFILE
    });
    dispatch({
      type: ACCOUNT_DELETED
    });
    console.log('i was called');
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: PROFILE_ERROR
    });
  }
};
