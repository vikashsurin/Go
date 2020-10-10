import {
  CREATE_PROFILE,
  PROFILE_ERROR,
  CURRENT_PROFILE,
  CLEAR_PROFILE
} from '../actions/types';

const initialState = {
  profile: '',
  profiles: [],
  isLoading: null
};

export default function(state = initialState, action) {
  const { type, payload } = action;
  switch (type) {
    case CREATE_PROFILE:
    case CURRENT_PROFILE:
      return {
        ...state,
        profile: payload,
        isLoading: false
      };
    case PROFILE_ERROR:
      return {
        ...state,
        profile: null,
        isLoading: false
      };

    case CLEAR_PROFILE:
      return {
        ...state,
        profile: null,
        isLoading: false
      };
    default:
      return state;
  }
}
