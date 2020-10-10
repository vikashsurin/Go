import {
  GET_POSTS,
  GET_POST,
  ADD_POST,
  EDIT_POST,
  UPDATE_LIKES,
  POST_ERROR,
  ADD_COMMENT,
  REMOVE_COMMENT,
  REMOVE_POST
} from "../actions/types";

const initialState = {
  posts: [],
  post: null,
  isLoading: true
};

export default function(state = initialState, action) {
  const { type, payload } = action;
  switch (type) {
    case GET_POSTS:
      return {
        ...state,
        posts: payload,
        post: null,
        isLoading: false
      };
    case ADD_POST:
      return {
        ...state,
        posts: [payload, ...state.posts],
        isLoading: false
      };
    case UPDATE_LIKES:
      return {
        ...state,
        posts: state.posts.map(post =>
          post._id === payload.id ? { ...post, likes: payload.likes } : post
        ),
        isLoading: false
      };
    case REMOVE_POST:
      return {
        ...state,
        posts: state.posts.filter(post => post._id !== payload),
        isLoading: false
      };
    case ADD_COMMENT:
      return {
        ...state,
        posts: state.posts.map(post =>
          post._id === payload.id
            ? { ...post, comments: payload.comments }
            : post
        )
      };
    case REMOVE_COMMENT:
      return {
        ...state,
        posts: state.posts.map(post =>
          post._id === payload.postId
            ? { ...post, comments: payload.comments }
            : post
        )
      };
    default:
      return state;
  }
}
